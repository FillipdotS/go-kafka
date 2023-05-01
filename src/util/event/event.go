package event

import (
	"github.com/google/uuid"
)

type event struct {
	Id    uuid.UUID
	Buyer string
	Item  string
	Price int
}

func NewEvent(buyer string, item string, price int) event {
	return event{
		Id:    uuid.New(),
		Buyer: buyer,
		Item:  item,
		Price: price,
	}
}
