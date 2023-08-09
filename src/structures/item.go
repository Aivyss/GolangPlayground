package structures

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	Name      string
	Price     int
	Special   bool
	createdAt time.Time
	id        uuid.UUID
}

func (i Item) ChangeId1(id uuid.UUID) {
	i.id = id
}

func (i *Item) ChangeId2(id uuid.UUID) {
	i.id = id
}

func (i Item) GetId() string {
	return i.id.String()
}

func NewItem1(name string, price int, special bool) Item {
	id, _ := uuid.NewUUID()

	return Item{
		Name:      name,
		Price:     price,
		Special:   special,
		createdAt: time.Now(),
		id:        id,
	}
}

func NewItem2(name string, price int, special bool) *Item {
	id, _ := uuid.NewUUID()

	return &Item{
		Name:      name,
		Price:     price,
		Special:   special,
		createdAt: time.Now(),
		id:        id,
	}
}
