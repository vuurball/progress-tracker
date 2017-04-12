package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//var wg sync.WaitGroup
var sites []Site
var notifsChannel = make(chan string)

func main() {
	rand.Seed(time.Now().Unix())

	// ini the sites
	for i := 0; i < 4; i++ {
		sites = append(sites, Site{"site " + strconv.Itoa(i), GetNumberOfPages(), 0})
	}

	//ini progress
	go TrackProgress(notifsChannel)

	for _, site := range sites {
		site.ReadAllPages(notifsChannel)
	}

	time.Sleep(300 * time.Millisecond) //wait for the last notification to reach the goroutine
	fmt.Println("\nfinished running")
}

// TrackProgress - goroutine loop that listens to a channel for new notifications
func TrackProgress(notifsChannel chan string) {
	for {
		select {
		case notif := <-notifsChannel:
			fmt.Printf("\r %v", notif)
		}
	}
}
