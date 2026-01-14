package main

import (
	bins "dz/app/Bins"
	storage "dz/app/Storage"
	"fmt"
)

func main() {
	list := bins.BinList{}
	bin := bins.NewBin("1", "config.json", false)
	list.Add(bin)

	err := storage.SaveBins("bins.json", list)
	if err != nil {
		return
	}
	list, err = storage.LoadBins("bins.json")
	if err != nil {
		return
	}

	fmt.Println(list)
}
