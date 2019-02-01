package slice

import (
	"errors"

	"wangzugames.com/kumay/gameServer2/util/mathUtil"
)

//IsSliceEqual check if 2 slices are equal
func IsSliceEqual(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

//RandomIntSlice returns new slice after random
func RandomIntSlice(src []int) ([]int, error) {

	slen := len(src)
	dst := make([]int, slen)

	if slen == 0 {
		return nil, errors.New("source slice len==0")
	} else if slen == 1 {
		copy(dst, src)
		return dst, nil
	}

	//fill -1
	for i := 0; i < slen; i++ {
		dst[i] = -1
	}

	for i := 0; i < slen; i++ {
		pos := mathUtil.RandomInt(0, len(src)-1)
		dst[i] = src[pos]
		src = append(src[:pos], src[pos+1:]...)
	}
	return dst, nil
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
