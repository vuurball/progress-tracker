package main

import (
	"fmt"
	"time"
)

// Observer type
type Observer struct {
	changes []string
}

// Update - recieves a new notification and pushes it to the changes slice
func (this *Observer) Update(notif string) {
	this.changes = append(this.changes, notif)
}

// GetChanges - returns all new notifications
func (this *Observer) GetChanges() []string {
	return this.changes
}

// Empty - removes all notifications from the slice
func (this *Observer) Empty() {
	this.changes = this.changes[:0]
}

// TrackProgress - goroutine loop that checks for new notifications, if found push to client and empty the notifs slice
func (this *Observer) TrackProgress() {
	fmt.Println("tracking")
	for {
		changes := this.GetChanges()
		if len(changes) > 0 {
			for _, notif := range changes {
				fmt.Println(notif)
			}
			this.Empty()
		}
		time.Sleep(200 * time.Millisecond) //run every half a sec and check progress

	}
}
