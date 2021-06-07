package data

import (
	"fmt"
)

type Data struct {
	slicee []string
}
type DataBase interface {
	Add(s string)
	Del(s string) error
	Find(s string) (int, error)
	Print()
}

func NewData() *Data {
	return new(Data)
}
func (d *Data) Add(s string) {
	d.slicee = append(d.slicee, s)
}

func (d *Data) Find(s string) (int, error) {
	for i := 0; i < len(d.slicee); i++ {
		if d.slicee[i] == s {
			return i, nil
		}
	}
	return 0, fmt.Errorf("aaaaaaaa")
}

func (d *Data) Del(s string) error {
	index, err := d.Find(s)
	if err != nil {
		return err
	}
	d.slicee[index] = d.slicee[len(d.slicee)-1]
	d.slicee = d.slicee[:len(d.slicee)-1]
	return nil
}

func (d *Data) Print() {
	for i := range d.slicee {
		fmt.Printf(" %v", d.slicee[i])
	}
	fmt.Print("\n")
}
