package pange

import (
	"sort"
	"strings"
)

func norm(ee []int) []int {
	uniq := make(map[int]bool)
	for _, e := range ee {
		uniq[e] = true
	}

	ee = []int{}
	for i, _ := range uniq {
		ee = append(ee, i)
	}
	sort.Ints(ee)
	return ee
}

func fuze(arr []Interval) (ar []Interval) {
	if len(arr) < 1 {
		return
	}
	// sort Intervals ascendind by their starts
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].A < arr[j].A
	})
	// push first Interval
	ar = append(ar, arr[0])
	for i := 1; i < len(arr); i++ {
		// last Interval
		top := ar[len(ar)-1]
		// add 1 to top.z because adjiacent integers must be jused
		// ex '1-3,4-5' must results as {1,5} not {1-3}, {4,5}
		// not overlapping so put it
		if top.Z+1 < arr[i].A {
			ar = append(ar, arr[i])
			// contained so update its end and replace last kept Interval
		} else if top.Z+1 <= arr[i].Z {
			top.Z = arr[i].Z
			ar = ar[:len(ar)-1]
			ar = append(ar, top)
		}
	}
	return
}

func trimspace(ss []string) []string {
	for i, s := range ss {
		ss[i] = strings.TrimSpace(s)
	}
	return ss
}
