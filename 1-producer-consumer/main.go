//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

var tweets = make(chan *Tweet)

func producer(stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}

		tweets <- tweet
	}

	// Close channel to make it non-blocking and signal end of producer
	close(tweets)
}

func consumer(tweets chan *Tweet) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	// writing into channel goroutine
	go producer(stream)

	// Consumer
	// not a goroutine since channel is closed and no longer blocking
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
