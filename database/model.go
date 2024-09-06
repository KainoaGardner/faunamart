package database

import "time"

type Ticket struct {
	ID        int
	Ticket    int
	Name      string
	CreatedAt time.Time
}
