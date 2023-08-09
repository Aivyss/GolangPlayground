package structures

import "time"

type Item struct {
	Name      string
	Price     int
	Special   bool
	createdAt time.Time
}

func NewItem1(name string, price int, special bool) Item {
	return Item{
		Name:      name,
		Price:     price,
		Special:   special,
		createdAt: time.Now(),
	}
}

func NewItem2(name string, price int, special bool) *Item {
	return &Item{
		Name:      name,
		Price:     price,
		Special:   special,
		createdAt: time.Now(),
	}
}
