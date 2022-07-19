```go
package main

func main() {}

func isIntGreaterThan(n, m int) bool {
	return n > m
}

func isFloatGreaterThan(n, m float64) bool {
	return n > m
}

func isStringGreaterThan(n, m string) bool {
	return n > m
}

func greaterThan(n, m Comparable) bool {
	...
}

```
