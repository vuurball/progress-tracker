package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//var wg sync.WaitGroup
var sites []Site

func main() {
	rand.Seed(time.Now().Unix())
	observer := Observer{}

	// ini the sites
	for i := 0; i < 4; i++ {
		sites = append(sites, Site{"site " + strconv.Itoa(i), GetNumberOfPages(), 0, &observer})
	}

	//ini progress

	//progress := Progress{&Site{"temp", 0, 0}}

	//pprogress := *progress

	//go progress.StartTrackingProgress()
	go observer.TrackProgress()

	for _, site := range sites {
		//progress.CurrentSite = &site
		//site.toPrint()
		//fmt.Println(site.name, site.numOfPages)
		site.ReadAllPages()
		fmt.Println(len(observer.GetChanges()))
	}

	fmt.Println("\nfinished running")
}
