package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {
	root := newComposite("root")

	root.Add(newLeaf("Leaf A"))
	root.Add(newLeaf("Leaf B"))

	c1 := newComposite("Composite X")
	c1.Add(newLeaf("Leaf XA"))
	c1.Add(newLeaf("Leaf XB"))

	c2 := newComposite("Composite Y")
	c2.Add(newLeaf("Leaf YA"))
	c2.Add(newLeaf("Leaf YB"))

	root.Add(c1)
	root.Add(c2)
	root.Add(newLeaf("Leaf C"))

	root.Display(1)
}
