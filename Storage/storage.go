package storage

import (
	bins "dz/app/Bins"
	files "dz/app/Files"
	"encoding/json"
	"errors"
)

func SaveBins(fileName string, b bins.BinList) error {
	data, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		return err
	}
	err = files.WriteFile(fileName, data)
	if err != nil {
		return err
	}
	return nil
}

func LoadBins(fileName string) (bins.BinList, error) {
	path := files.IsJson(fileName)
	if path == false {
		err := errors.New("file is not a json")
		return bins.BinList{}, err
	}
	data, err := files.ReadFile(fileName)
	if err != nil {
		return bins.BinList{}, err
	}
	var bin bins.BinList
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return bins.BinList{}, err
	}
	return bin, nil
}
