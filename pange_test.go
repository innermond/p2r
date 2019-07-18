package pange

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tt := []struct {
		p    string
		want []Interval
	}{
		{"1,2,3",
			[]Interval{{1, 3}},
		},
		{"1,2,3",
			[]Interval{{1, 3}},
		},
		{"2-4,1-3, 5-6, 8-8, 11, 12,9-10,10,9,9,8-9",
			[]Interval{{1, 6}, {8, 12}},
		},
		{"1-3,8",
			[]Interval{{1, 3}, {8, 8}},
		},
		{"1-3,2-5, 7-8, 9, 10",
			[]Interval{{1, 5}, {7, 10}},
		},
	}
	for i, tc := range tt {
		sel := Selection(tc.p)
		got, err := sel.Split()
		if err != nil {
			t.Fatal(err)
		}
		//t.Logf("%d) got %v", i, got)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%d) got %v want %v", i, got, tc.want)
		}
	}
}

func TestFull(t *testing.T) {
	tt := []struct {
		p    string
		want []int
	}{
		{"1,2,3",
			[]int{1, 2, 3},
		},
		{"2-4,1-3, 5-6, 8-8, 11, 12,9-10,10,9,9,8-9",
			[]int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12},
		},
		{"1-3,8",
			[]int{1, 2, 3, 8},
		},
		{"1-3,2-5, 7-8, 9, 10",
			[]int{1, 2, 3, 4, 5, 7, 8, 9, 10},
		},
	}
	for i, tc := range tt {
		sel := Selection(tc.p)
		got, err := sel.Full()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%d) got %v want %v", i, got, tc.want)
		}
	}
}

func TestNorm(t *testing.T) {
	tt := []struct {
		in   []int
		want []int
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 10, 3, 3, 2, 1},
			[]int{1, 2, 3, 10},
		},
		{
			[]int{5, 4, 4, 5, 5, 7, 1},
			[]int{1, 4, 5, 7},
		},
		{
			[]int{},
			[]int{},
		},
	}
	for i, tc := range tt {
		got := norm(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%d) got %v want %v", i, got, tc.want)
		}
	}
}

func TestPagesSelection(t *testing.T) {
	tt := []struct {
		p    string
		want []int
	}{
		{"1,2,3",
			[]int{1, 2, 3}},
		{"1-3,8",
			[]int{1, 2, 3, 8},
		},
		{"3-1,1-3,2,1,4",
			[]int{1, 2, 3, 4},
		},
		{"-2",
			[]int{1, 2},
		},
		{"-2,1,2-3",
			[]int{1, 2, 3},
		},
		{"2,5-6",
			[]int{2, 5, 6},
		},
	}
	for i, tc := range tt {
		ss, err := Selection(tc.p).Split()
		if err != nil {
			t.Errorf("%d) expected error", i)
		}
		got := []int{}
		for _, ee := range ss {
			for i := ee.A; i <= ee.Z; i++ {
				got = append(got, i)
			}
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%d) got %v want %v", i, got, tc.want)
		}
	}
}

func TestSeparators(t *testing.T) {
	expr := "1:3#2:4#6"
	want := []Interval{
		{1, 4},
		{6, 6},
	}

	got, err := Selection(expr).Split("#", ":")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestErrorsSelection(t *testing.T) {
	tt := []struct {
		p    string
		want []int
	}{
		{"  ",
			[]int{-1},
		},
		{"1-",
			[]int{-1}},
		/*{"1-3,8",
			[]int{1, 2, 3, 8},
		},
		{"3-1,1-3,2,1,4",
			[]int{1, 2, 3, 4},
		},
		{"-2",
			[]int{1, 2},
		},
		{"-2,1,2-3",
			[]int{1, 2, 3},
		},
		{"2,5-6",
			[]int{2, 5, 6},
		},*/
	}
	for i, tc := range tt {
		_, err := Selection(tc.p).Split()
		if err == nil {
			t.Errorf("%d) expected error", i)
		}
	}
}

func ExampleOneToTen() {
	ss, _ := Selection("1-10").Split()
	for _, ee := range ss {
		for i := ee.A; i <= ee.Z; i++ {
			fmt.Print(i)
		}
	}
	// Output: 12345678910
}
