package composite

import (
	"errors"
	"fmt"
	"strings"
)

type Leaf struct {
	name string
}

func newLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

func (this *Leaf) Add(c Component) error {
	return errors.New("Cannot add to a leaf.")
}

func (this *Leaf) Remove(c Component) error {
	return errors.New("Cannot add to a leaf.")
}

func (this *Leaf) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), this.name)
}
