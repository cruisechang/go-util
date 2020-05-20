package slice

import (
	"fmt"
	"testing"
)

func TestRandomIntSlice(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}

	r, err := RandomIntSlice(src)

	if err != nil {
		fmt.Printf("%v", err)

	} else {
		fmt.Printf("%v", r)

	}
}
