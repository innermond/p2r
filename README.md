# Pange

## What it does?
This code converts a string description of what numbered pages are requested to a collection (slice) of slices.
Every slice of collection represents a range of pages. A range is a slice having two elements, first element is from where a range starts,
second show where it ends.

It does nothing else more but it takes care of overlapping ranges, converts single pages to a range, etc.

If you need to iterate just iterate a slice of slices.

## Why?
This code arised from a basic need when you are dealing with topics specific to printing industry - **to select certain pages** in order to process them.

## Why that name??
I know it sounds good and misterious and raises great expectations regarding code quality and about what really does but...

Being about pages and ranges I came up naturally with this name.

## How to use it
```golang
// by default, ranges are using dash - for their ends and are separated by comma ,
ss, _ := Selection("2,1-3,5-2,6-7,9-10").Split()
// you get a slice of slices []pange.Interval{{1,7}, {9, 10}}
// iterate it to get 1,2,3,4,5,6,7,<not 8>9,10
for _, ee := range ss {
        for i := ee.A; i <= ee.Z; i++ {
                fmt.Print(i)
        }
}

```
