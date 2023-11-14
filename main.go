package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
С помощью горутин и каналов смоделировать работу пекарни:

Пекарня состоит из 2+1 этапов: торт выпекается, торт запаковывается, все торты смотрятся. Связь между этапами происходит через каналы,
этапы выполняются с помощью горутин. Торт представляет из себя объект с полями типа int: BakedBy, BakeTime, PackedBy, PackTime

Программа должна работать следующим образом:

Запускается N горутин, каждая из которых за время T1 = i +-t1 (i - номер рутины, t1 - выбранный вами параметр) создаёт объект тортика,
заполняет поля BackedBy=i, BakeTime=T1 и отправляет его в канал

Существует пул из M горутин. когда в канал приходят торитики с предыдущего этапа,
свободная рутина из пула начинает его упаковывать,
то есть за время T2 = j+-t2 (j - номер рутины из пула, t2 - выбранный вами параметр, причём t2>=t1)
выставляет тортику параметры PackedBy=j, PackTime = T2 и отправляет в канал

Мы ждём пока придут все тортики, или мы получим сигнал о завершении работы(тогда дожидаемся завершения работы текущих рутин),
после чего просто выводим тортики в порядке, в котором они пришли к нам. При полном отрабатывании программы должно быть K тортиков.

Для данной задачи можно использовать различные наборы параметров, однако рекоммендуется попробовать следующие комбинации:

K = 10_000, N = 1, M = 1

K = 10_000, N = 8, M = 5

K = 10_000, N = 100, M = 50
*/

var r = rand.New(rand.NewSource(time.Now().Unix()))

type cake struct {
	id       uint64
	BakedBy  uint64
	BakeTime time.Duration
	PackedBy uint64
	PackTime time.Duration
}

// not pointer
func (c cake) String() string {
	return fmt.Sprintf("cake: id %v, BakedBy %v,\tBakeTime %v,\tPackedBy %v,\tPackTime %v", c.id, c.BakedBy, c.BakeTime, c.PackedBy, c.PackTime)
}

func getRandDuration(i uint64) time.Duration {
	return time.Microsecond*time.Duration(i) +
		time.Microsecond*time.Duration(int64(float64(r.Intn(1000))*r.Float64())) +
		time.Microsecond
}

func bake(ctx context.Context, wg *sync.WaitGroup, backedCakeOutCh chan<- cake, currentCakeInCh <-chan uint64, i uint64) {
	defer wg.Done()

	for {
		select {
		case id, ok := <-currentCakeInCh:
			if ok {
				T1 := getRandDuration(i)
				time.Sleep(T1)

				backedCakeOutCh <- cake{id: id, BakedBy: i, BakeTime: T1}
			}
		case <-ctx.Done():
			return
		}
	}
}

func pack(ctx context.Context, wg *sync.WaitGroup, packedCakeOutCh chan<- cake, bakedCakeToPackInCh <-chan cake, i uint64) {
	defer wg.Done()

	for {
		select {
		case c, ok := <-bakedCakeToPackInCh:
			if ok {
				T2 := getRandDuration(i)
				time.Sleep(T2)
				c.PackedBy = i
				c.PackTime = T2

				packedCakeOutCh <- c
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	// totalCake := flag.Uint64("K", 0, "total cakes for baking")
	// N := flag.Int("N", 0, "total furnaces")
	// M := flag.Int("M", 0, "total packers")
	// c := cake{1, time.Duration(1 * time.Second), 1, time.Microsecond * 2}
	// fmt.Println(c)
	// return
	start := time.Now()
	var totalCake uint64 = 1000
	var N uint64 = 1
	var M uint64 = 1
	delay := 5 * time.Second

	flag.Parse()

	currentCakeProducerCh := make(chan uint64, runtime.NumCPU())
	bakedCakeProducerCh := make(chan cake, runtime.NumCPU())
	packedCakeProducerCh := make(chan cake, runtime.NumCPU())

	var wgFurnace sync.WaitGroup
	var wgPacker sync.WaitGroup

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, delay)
	time.AfterFunc(5*time.Second, cancel)

	go func(ctx context.Context, totalCakesTo uint64, outCh chan<- uint64) {
		i := uint64(0)
		for ; i < totalCakesTo; i++ {
			outCh <- i
		}
		fmt.Println("all orders passed. Total:", i)
		close(outCh)
	}(ctx, totalCake, currentCakeProducerCh)

	go func() {
		for i := uint64(0); i < N; i++ {
			wgFurnace.Add(1)
			go bake(ctx, &wgFurnace, bakedCakeProducerCh, currentCakeProducerCh, i)
		}

		wgFurnace.Wait()
		close(bakedCakeProducerCh)
	}()

	go func() {
		for i := uint64(0); i < M; i++ {
			wgPacker.Add(1)
			go pack(ctx, &wgPacker, packedCakeProducerCh, bakedCakeProducerCh, i)
		}

		wgPacker.Wait()
		close(packedCakeProducerCh)
	}()

	go func(cf context.CancelFunc) {
		var canseledLine string
		fmt.Scanln(&canseledLine)
		cf()

		fmt.Println("canseled by this line", canseledLine)
	}(cancel)

	totalPackedBakedCake := uint64(0)
	for c := range packedCakeProducerCh {
		totalPackedBakedCake++
		fmt.Println(totalPackedBakedCake, c)
	}
	fmt.Printf("Done! Total cake %v out of %v in %v, delay was %v", totalPackedBakedCake, totalCake, time.Since(start), delay)
}
