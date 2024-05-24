# Using `go build`

`GOOS="windows" go build` builds an exe file which can be directly run in a windows environment, we can use darwin for mac and linux for linux environment

# Memory Management in Golang

## The new() Function

Allocates memory for a new zeroed value of a specific type and returns a pointer to it.  
Example,

> x := new(int)  
> *x = 5  
> fmt.Println(*x) // prints 5

## The make() Function

Returns an initialised (non-zeroed) value of a specific type
Example,

> arr := make([]int, 5, 10) // creates a slice of integers of length 5 and capacity 10
