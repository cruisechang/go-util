package slice

import (
	"fmt"
	"testing"

	"wangzugames.com/kumay/gameServer2/util/sliceUtil"
)

func TestRandomIntSlice(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}

	r, err := sliceUtil.RandomIntSlice(src)

	if err != nil {
		fmt.Printf("%v", err)

	} else {
		fmt.Printf("%v", r)

	}
}
