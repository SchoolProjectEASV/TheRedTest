package warehouse

import "time"

type ThreeDRoom struct {
	Height float64
	Width  float64
	Length float64
}

type Item struct {
	ItemId     int
	ItemName   string
	ItemHeight float64
	ItemWidth  float64
	ItemLength float64
	StartDate  time.Time
	EndDate    time.Time
	IsActive   bool
}

type Warehouse struct {
	Id          int
	MaxCapacity ThreeDRoom
	Items       []Item
}
