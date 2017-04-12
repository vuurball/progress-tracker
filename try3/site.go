package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Site struct {
	Name        string
	NumOfPages  int
	CurrentPage int
}

func (this *Site) notifyObservers(notifsChannel chan<- string) {
	percent := int(float64(this.CurrentPage) / float64(this.NumOfPages) * 100)
	notif := this.Name + " [" + strconv.Itoa(percent) + "%]"
	notifsChannel <- notif

}

// ReadAllPages read all pages in a site
func (this *Site) ReadAllPages(notifsChannel chan<- string) {
	for i := 0; i < this.NumOfPages; i++ {
		time.Sleep(500 * time.Millisecond) //wait half sec
		this.CurrentPage++
		this.notifyObservers(notifsChannel)

	}
}

// GetNumberOfPages returns a random number to simulate num of pages in a site
func GetNumberOfPages() int {
	return rand.Intn(5) + 1
}
