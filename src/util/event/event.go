package event

import (
	"math/rand"

	"github.com/google/uuid"
)

type Event struct {
	Id    uuid.UUID
	Buyer string
	Item  string
	Price int
}

func New(buyer string, item string, price int) Event {
	return Event{
		Id:    uuid.New(),
		Buyer: buyer,
		Item:  item,
		Price: price,
	}
}

func NewRandom() Event {
	names := [...]string{"Fillip", "John", "Arthur", "Bob", "Gary", "Michael", "Mikhail", "Dmitry", "Sam", "Sally", "Jane", "Andrea", "Walter"}
	items := [...]string{"Book", "Pen", "Pencil", "Notebook", "Laptop", "Phone", "Tablet", "Monitor", "Keyboard", "Mouse", "Headphones", "Speakers", "Microphone"}

	return New(
		names[rand.Intn(len(names))],
		items[rand.Intn(len(items))],
		rand.Intn(1000),
	)
}