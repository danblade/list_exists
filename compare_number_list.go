package compare

import (
	"sort"
)

type Container struct {
	Numbers []int
}

func NewContainer(n []int) *Container {
	sort.Ints(n)
	return &Container{Numbers: n}
}

func (c *Container) Contains(needle int) bool {
	j := sort.Search(len(c.Numbers), func(i int) bool { return c.Numbers[i] >= needle })
	return j < len(c.Numbers) && c.Numbers[j] == needle
	// return !(j == -1 || c.Numbers[j] != needle)

}
