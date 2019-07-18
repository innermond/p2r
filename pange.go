package pange

import (
	"errors"
	"strconv"
	"strings"
)

type Interval struct {
	A, Z int
}

type Selection string

func (sel Selection) Split(seps ...string) (ends []Interval, err error) {

	var (
		ll           []string
		a, z         int
		intervalSep  = "-"
		selectionSep = ","
	)

	switch len(seps) {
	case 1:
		selectionSep = seps[0]
	case 2:
		selectionSep = seps[0]
		intervalSep = seps[1]
	}

	gg := strings.Split(strings.TrimSpace(string(sel)), selectionSep)

	for _, g := range gg {
		if strings.Contains(g, intervalSep) {
			ll = strings.Split(g, intervalSep)
		} else {
			g = strings.TrimSpace(g)
			ll = []string{g, g}
		}
		ll = trimspace(ll)
		if ll[0] == "" {
			a = 1
		} else {
			a, err = strconv.Atoi(ll[0])
			if err != nil {
				return
			}
		}
		if ll[1] == "" {
			err = errors.New("Interval needs upper end")
			return
		} else {
			z, err = strconv.Atoi(ll[1])
			if err != nil {
				return
			}
		}
		if a > z {
			a, z = z, a
		}
		ends = append(ends, Interval{a, z})
	}
	ends = fuze(ends)
	return
}

func (ss Selection) Full() (full []int, err error) {
	ii, err := ss.Split()
	if err != nil {
		return
	}

	// all intervals as a slice of ints
	for _, i := range ii {
		for x := i.A; x <= i.Z; x++ {
			full = append(full, x)
		}
	}

	return
}
