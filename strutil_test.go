package strutil

import (
	"fmt"
	"testing"
)

func TestRemDup(t *testing.T) {
	v := []string{"a", "a", "b", "c", "d", "d", "d", "e", "m", "q", "q"}
	fmt.Printf("Input:\t%s\tCapacity: %d\n", v, cap(v))

	s := RemDup(v)
	fmt.Printf("Output:\t%v\t\tCapacity: %d\n", s, cap(s))
}
