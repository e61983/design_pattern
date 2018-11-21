package composite

import (
	"errors"
	"fmt"
	"strings"
)

type Composite struct {
	children []Component
	name     string
}

func newComposite(name string) *Composite {
	return &Composite{name: name}
}

func (this *Composite) Add(c Component) error {
	this.children = append(this.children, c)
	return nil
}

func (this *Composite) Remove(c Component) error {
	var i int
	for i = range this.children {
		if this.children[i] == c {
			break
		}
	}
	if this.children[i] != c {
		return errors.New("Cannot find component.")
	}
	this.children = append(this.children[:i], this.children[i+1:]...)
	return nil
}

func (this *Composite) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), this.name)
	for _, c := range this.children {
		c.Display(depth + 2)
	}
}
