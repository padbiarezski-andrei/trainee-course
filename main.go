package main

import (
	"context"
	"flag"
	"fmt"
	"math"
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
	BakedBy  uint64
	BakeTime time.Duration
	PackedBy uint64
	PackTime time.Duration
}

// not pointer
func (c cake) String() string {
	return fmt.Sprintf("cake: BakedBy %v,\tBakeTime %v,\tPackedBy %v,\tPackTime %v", c.BakedBy, c.BakeTime, c.PackedBy, c.PackTime)
}

func bake(ctx context.Context, wg *sync.WaitGroup, backedCakeOutCh chan<- cake, currentCakeInCh <-chan uint64, i uint64) {
	var T1 time.Duration
	if r.Intn(2) != 0 {
		T1 = time.Microsecond*time.Duration(i) + time.Microsecond*time.Duration(int64(float64(r.Int63n(int64(i%math.MaxInt64)+1))*r.Float64())) + time.Microsecond
	} else {
		T1 = time.Microsecond*time.Duration(i) - time.Microsecond*time.Duration(int64(float64(r.Int63n(int64(i%math.MaxInt64)+1))*r.Float64())) + time.Microsecond
	}
	time.Sleep(T1)

	for range currentCakeInCh {
		backedCakeOutCh <- cake{BakedBy: i, BakeTime: T1}
	}

	wg.Done()
}

func pack(ctx context.Context, wg *sync.WaitGroup, packedCakeOutCh chan<- cake, cakeToPackInCh <-chan cake, i uint64) {
	var T2 time.Duration
	if r.Intn(2) != 0 {
		T2 = time.Microsecond*time.Duration(i) + time.Microsecond*time.Duration(int64(float64(r.Int63n(int64(i%math.MaxInt64)+1))*r.Float64())) + time.Microsecond
	} else {
		T2 = time.Microsecond*time.Duration(i) - time.Microsecond*time.Duration(int64(float64(r.Int63n(int64(i%math.MaxInt64)+1))*r.Float64())) + time.Microsecond
	}
	time.Sleep(T2)

	for c := range cakeToPackInCh {
		c.PackedBy = i
		c.PackTime = T2

		packedCakeOutCh <- c
	}

	wg.Done()
}

func main() {
	// totalCake := flag.Uint64("K", 0, "total cakes for baking")
	// N := flag.Int("N", 0, "total furnaces")
	// M := flag.Int("M", 0, "total packers")
	// c := cake{1, time.Duration(1 * time.Second), 1, time.Microsecond * 2}
	// fmt.Println(c)
	// return

	var totalCake uint64 = 1000
	var N uint64 = 50
	var M uint64 = 20

	flag.Parse()

	currentCakeProducerCh := make(chan uint64, runtime.NumCPU())
	cakeProducerCh := make(chan cake, runtime.NumCPU())
	packedCakeProducerCh := make(chan cake, runtime.NumCPU())

	var wgFurnace sync.WaitGroup
	var wgPacker sync.WaitGroup

	ctx := context.Background()

	go func(totalCakesTo uint64, outCh chan<- uint64) {
		for i := uint64(0); i < totalCakesTo; i++ {
			outCh <- i
		}
		close(currentCakeProducerCh)
	}(totalCake, currentCakeProducerCh)

	go func() {
		for i := uint64(0); i < N; i++ {
			wgFurnace.Add(1)
			// i / -2 * [0,1) + i
			// i = 10
			// -5 * [0,1) + 10
			// -2.5 + 10 = 7.5

			go bake(ctx, &wgFurnace, cakeProducerCh, currentCakeProducerCh, i)
		}

		wgFurnace.Wait()
		close(cakeProducerCh)
	}()

	go func() {
		for i := uint64(0); i < M; i++ {
			wgPacker.Add(1)

			go pack(ctx, &wgPacker, packedCakeProducerCh, cakeProducerCh, i)
		}

		wgPacker.Wait()
		close(packedCakeProducerCh)
	}()

	totalPackedBakedCage := uint64(0)
	for c := range packedCakeProducerCh {
		totalPackedBakedCage++
		fmt.Println(totalPackedBakedCage, c)
	}

	select {}
}
