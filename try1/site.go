package main

import (
	"math/rand"
	"time"
)

type Site struct {
	Name        string
	NumOfPages  int
	CurrentPage int
}

// read all pages in a site
func (this *Site) ReadAllPages() {
	for i := 0; i < this.NumOfPages; i++ {
		time.Sleep(1000 * time.Millisecond) //wait 1 sec
		this.CurrentPage++
		//fmt.Println(this.Name, "reading page", this.CurrentPage, "out of", this.NumOfPages)
	}
}

// get random number to simulate num of pages in a site
func GetNumberOfPages() int {
	return rand.Intn(5) + 1
}
