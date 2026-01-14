package bins

import "time"

type BinDl interface {
	SaveBins(fileName string, list BinList) error
	LoadBins(fileName string) (BinList, error)
}

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

func (bl *BinList) Add(bin Bin) {
	bl.Bins = append(bl.Bins, bin)
}
