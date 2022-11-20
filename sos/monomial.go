package sos

/*
monomial.go
Description:
	A list of files that is relevant/helpful for the
*/

/*
Type Definitions
*/

type Monomial struct {
	Variables   []*Variable
	Exponents   []int
	Coefficient float64
}
