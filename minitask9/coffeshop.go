package minitask9

import (
	"fmt"
	"sync"
	"time"
)

func CoffeeShop() {
	var wg sync.WaitGroup
	var chn chan string
	chn = make(chan string, 3)
	// for _, coffee := range coffees {
	// 	wg.Add(1)
	// 	go Barista(coffee, &wg)
	// }
	baristaCount := 3
	for range baristaCount {
		wg.Add(1)
		go Barista(chn, &wg)
	}

	for i := range 12 {
		order := fmt.Sprintf("Kopi %d", i)
		chn <- order
		fmt.Printf("Order masuk untuk menu %s\n", order)
		time.Sleep(1 * time.Millisecond)
	}
	close(chn)
	wg.Wait()
	fmt.Println("Toko Tutup")
}

func Barista(coffeeChn chan string, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Barista Pulang")
		wg.Done()
	}()
	for coffee := range coffeeChn {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("Start making coffee: %s\n", coffee)
		time.Sleep(1 * time.Second)
		fmt.Printf("Finished making coffee: %s\n", coffee)
	}
}
