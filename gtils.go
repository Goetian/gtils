package gtils

import "golang.org/x/exp/constraints"

func Map[T constraints.Ordered](values []T, mapfunc func(T) T) []T {
	var ret []T
	for _, v := range values {
		nvalue := mapfunc(v)
		ret = append(ret, nvalue)
	}
	return ret
}
