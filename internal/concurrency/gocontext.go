package concurrency

import (
	"context"
	"fmt"
	"time"
)

type bid struct {
	AdUrl string
	Price float64
}

var defaultBid bid = bid{AdUrl: "https://adsrus.com/default", Price: 0.02}

func GoContext() {

	bestBid := func(url string) bid {
		time.Sleep(20 * time.Millisecond)
		return bid{AdUrl: "https://adsrus.com/19", Price: 0.05}
	}

	findBid := func(ctx context.Context, url string) bid {
		ch := make(chan bid, 1) // buffered channel to avoid goroutine leak
		go func() {
			ch <- bestBid(url)
		}()

		select {
		case bid := <-ch:
			return bid
		case <-ctx.Done():
			return defaultBid
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	bid := findBid(ctx, "https://http.cat/418")
	fmt.Printf("Found bid %+v\n", bid)
}
