package composite

type Component interface {
	Add(Component) error
	Remove(Component) error
	Display(int)
}
