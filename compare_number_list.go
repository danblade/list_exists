package compare

type Container struct {
	Numbers map[int]bool
}

func NewContainer(numbers []int) *Container {
	m := make(map[int]bool, len(numbers))
	for _, n := range numbers {
		m[n] = true
	}
	return &Container{Numbers: m}
}

func (c *Container) Contains(needle int) bool {
	_, present := c.Numbers[needle]
	return present
}
