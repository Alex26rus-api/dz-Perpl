package main

import (
	bins "dz/app/Bins"
	"fmt"
)

func main() {
	list := bins.BinList{}
	bin := bins.NewBin("1", "config.json", false)
	list.Add(bin)

	fmt.Println(list)
}
