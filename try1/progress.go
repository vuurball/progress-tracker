package main

import (
	"fmt"
	"time"
)

type Progress struct {
	CurrentSite *Site
}

func (this *Progress) StartTrackingProgress() {
	fmt.Println("Reading...")
	lastSite := Site{"test", 0, 0}

	for {

		//	fmt.Println("loop", lastSite, this.currentSite.Name)
		//fmt.Println("progress.currentSite loop", this.CurrentSite)
		if lastSite.Name != this.CurrentSite.Name {
			fmt.Println("")
			fmt.Print(this.CurrentSite.Name, " (", this.CurrentSite.NumOfPages, ") ")
		}
		if lastSite.CurrentPage != this.CurrentSite.CurrentPage {
			fmt.Print(".")
		}
		lastSite = *this.CurrentSite
		time.Sleep(500 * time.Millisecond) //run every half a sec and check progress

	}
}
