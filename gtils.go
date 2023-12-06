package gtils

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func Map[T constraints.Ordered](values []T, mapfunc func(T) T) []T {
	var ret []T
	for _, v := range values {
		nvalue := mapfunc(v)
		ret = append(ret, nvalue)
	}
	return ret
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	} else {
		return true
	}
}
