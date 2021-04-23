package compare

import (
	"fmt"
	"math/rand"
	"time"
)

type Container struct {
	Numbers []int
}

func NewNumbers() Container {
	rand.Seed(time.Now().Unix())
	numbers := rand.Perm(1000)
	var milNum Container
	milNum.Numbers = numbers
	return milNum
}

func main() {
	n := NewNumbers()
	// fmt.Println(n)
	fmt.Println(n.Contains(45), "45")
	fmt.Println(n.Contains(405), "405")
	fmt.Println(n.Contains(1235), "1235")
	fmt.Println(n.Contains(4500), "4500")
	fmt.Println(n.Contains(511), "511")
	fmt.Println(n.Contains(925), "925")
}

func (c *Container) Contains(n int) bool {
	for _, listNumber := range c.Numbers {
		if listNumber == n {
			return true
		}
	}
	return false

}
