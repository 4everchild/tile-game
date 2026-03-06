package test

import (
	"azul/game"
	"fmt"
	"sort"
)

func TestXS64() {
	const l uint64 = 13
	const n uint64 = 1 << 20
	m := uint64(10000)
	var vals [1 << l]uint64
	fmt.Println("### ", l, " ", n, "\n")
	for j := uint64(1); j < n; j++ {
		if j%m == m-1 {
			fmt.Println("### ", j)
		}
		var s game.XS64 = game.XS64(j)
		for i := 0; i < 1<<l; i++ {
			vals[i] = s.Value()
			//fmt.Println(vals[i], uint64(s))
			s = s.Next()
		}
		sort.Slice(vals[:], func(i, j int) bool {
			return vals[i] < vals[j]
		})

		for i := uint64(0); i < l-1; i++ {
			if vals[i] == vals[i+1] {
				fmt.Println(j, " ", vals[i])
			}
		}
	}

}
