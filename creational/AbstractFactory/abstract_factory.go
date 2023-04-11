package bbstract_factory

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("Now I cooking Rise.")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("Now I cooking Tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
}

func NewSimpleLunchFactory() *SimpleLunchFactory {
	return &SimpleLunchFactory{}
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
