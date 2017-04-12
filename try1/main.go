package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var sites []Site

func main() {
	rand.Seed(time.Now().Unix())

	// ini the sites
	for i := 0; i < 4; i++ {
		sites = append(sites, Site{"site " + strconv.Itoa(i), GetNumberOfPages(), 0})
	}

	//ini progress
	progress := Progress{&Site{"temp", 0, 0}}

	go progress.StartTrackingProgress()

	for _, site := range sites {
		progress.CurrentSite = &site
		//site.toPrint()
		//fmt.Println(site.name, site.numOfPages)
		site.ReadAllPages()

	}

	fmt.Println("\nfinished running")
}
