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

// go run . --K=1000 --N=50 --M=10
// Done! Total cake 71 out of 1000 in 5.716077123s, delay was 3s

// go run . --K=1000 --N=8 --M=5
// Done! Total cake 28 out of 1000 in 4.993860512s, delay was 3s

// go run . --K=1000 --N=250 --M=125
// Done! Total cake 849 out of 1000 in 7.796713981s, delay was 3s

// go run . --K=1000 --N=125 --M=250
// Done! Total cake 886 out of 1000 in 6.817758505s, delay was 3s

var r = struct {
	m   *sync.Mutex
	src *rand.Rand
}{m: &sync.Mutex{}, src: rand.New(rand.NewSource(time.Now().Unix()))}

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
	r.m.Lock()
	defer r.m.Unlock()
	return time.Microsecond*time.Duration(i) +
		time.Nanosecond*time.Duration(r.src.Int63n(1000000)) +
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
			} else {
				return
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
			} else {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	totalCake := flag.Uint64("K", 0, "total cakes required")
	N := flag.Uint64("N", 0, "total furnaces")
	M := flag.Uint64("M", 0, "total packers")

	flag.Parse()

	// var q uint64 = 1000
	// var a uint64 = 50
	// var z uint64 = 10
	// var totalCake *uint64 = &q
	// var N *uint64 = &a
	// var M *uint64 = &z

	delay := 3 * time.Second
	start := time.Now()

	currentCakeCh := make(chan uint64, runtime.NumCPU())
	bakedCakeCh := make(chan cake, runtime.NumCPU())
	packedCakeCh := make(chan cake, runtime.NumCPU())

	var wgFurnaces sync.WaitGroup
	var wgPackers sync.WaitGroup

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, delay)
	defer cancel()

	go func(ctx context.Context, totalCakesTo uint64, outCh chan<- uint64) {
		for i := uint64(0); i < totalCakesTo; i++ {
			outCh <- i
		}
		close(outCh)
	}(ctx, *totalCake, currentCakeCh)

	go func() {
		for i := uint64(0); i < *N; i++ {
			wgFurnaces.Add(1)
			go bake(ctx, &wgFurnaces, bakedCakeCh, currentCakeCh, i)
		}

		wgFurnaces.Wait()
		close(bakedCakeCh)
	}()

	go func() {
		for i := uint64(0); i < *M; i++ {
			wgPackers.Add(1)
			go pack(ctx, &wgPackers, packedCakeCh, bakedCakeCh, i)
		}

		wgPackers.Wait()
		close(packedCakeCh)
	}()

	go func(cf context.CancelFunc) {
		var canseledLine string
		fmt.Scanln(&canseledLine)
		cf()

		fmt.Println("canseled by this line", canseledLine)
	}(cancel)

	totalPackedBakedCake := uint64(0)
	for c := range packedCakeCh {
		totalPackedBakedCake++
		fmt.Println(totalPackedBakedCake, c)
	}
	fmt.Printf("Done! Total cake %v out of %v in %v, delay was %v\n", totalPackedBakedCake, *totalCake, time.Since(start), delay)
}
