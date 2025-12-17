package main

import "time"
import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func NewBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func (bl *BinList) Add(b Bin) {
	bl.Bins = append(bl.Bins, b)
}

func main() {
	list := BinList{}

	bin := NewBin("1", "config.json", false)
	list.Add(bin)

	fmt.Println(list)
}
