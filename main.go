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
<<<<<<< HEAD
	totalCake := flag.Uint64("K", 0, "total cakes required")
	N := flag.Uint64("N", 0, "total furnaces")
	M := flag.Uint64("M", 0, "total packers")

	flag.Parse()

	delay := 5 * time.Second
	start := time.Now()

	currentCakeCh := make(chan uint64, runtime.NumCPU())
	bakedCakeCh := make(chan cake, runtime.NumCPU())
	packedCakeCh := make(chan cake, runtime.NumCPU())

=======
	// totalCake := flag.Uint64("K", 0, "total cakes required")
	// N := flag.Uint64("N", 0, "total furnaces")
	// M := flag.Uint64("M", 0, "total packers")

	start := time.Now()
	var totalCake uint64 = 1000
	var N uint64 = 1
	var M uint64 = 1

	delay := 5 * time.Second

	flag.Parse()

	currentCakeCh := make(chan uint64, runtime.NumCPU())
	bakedCakeCh := make(chan cake, runtime.NumCPU())
	packedCakeCh := make(chan cake, runtime.NumCPU())

>>>>>>> dfd8ab85aae3bf75030bebcffa21bdc20d0fe3a7
	var wgFurnaces sync.WaitGroup
	var wgPackers sync.WaitGroup

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, delay)
	defer cancel()
	time.AfterFunc(5*time.Second, cancel)

	go func(ctx context.Context, totalCakesTo uint64, outCh chan<- uint64) {
		for i := uint64(0); i < totalCakesTo; i++ {
			outCh <- i
		}
		close(outCh)
<<<<<<< HEAD
	}(ctx, *totalCake, currentCakeCh)

	go func() {
		for i := uint64(0); i < *N; i++ {
=======
	}(ctx, totalCake, currentCakeCh)

	go func() {
		for i := uint64(0); i < N; i++ {
>>>>>>> dfd8ab85aae3bf75030bebcffa21bdc20d0fe3a7
			wgFurnaces.Add(1)
			go bake(ctx, &wgFurnaces, bakedCakeCh, currentCakeCh, i)
		}

		wgFurnaces.Wait()
		close(bakedCakeCh)
	}()

	go func() {
<<<<<<< HEAD
		for i := uint64(0); i < *M; i++ {
=======
		for i := uint64(0); i < M; i++ {
>>>>>>> dfd8ab85aae3bf75030bebcffa21bdc20d0fe3a7
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
	fmt.Printf("Done! Total cake %v out of %v in %v, delay was %v", totalPackedBakedCake, totalCake, time.Since(start), delay)
}
