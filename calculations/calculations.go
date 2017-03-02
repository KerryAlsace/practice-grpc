// Package calculations contains functions for manipulating integers
package calculations

// AddOne takes and integer and adds 1
func AddOne(n *int32) {
	*n = *n + 1
}

// Square takes and integer and squares it
func Square(n *int32) {
	*n = *n * n
}
