package sos

/*
polynomial.go
Description:
	Defines the Polynomial object and attempts to define a new interface for it.
*/

type Polynomial struct {
	Monomials []Monomial
}

type PolynomialInterface interface {
	Find(sliceIn []interface{}) int
}

// Functions
// =========

/*
FindPolynomialInSlice
Description:

	Looks for a Polynomial in a slice of objects that implement the empty interface (should be any object).
	Returns -1 if none are found.
*/
func FindPolynomialInSlice(sliceIn []interface{}) int {
	// Constants

	// Algorithms
	for eltIndex := 0; eltIndex < len(sliceIn); eltIndex++ {
		// check to see if current element is a Variable
		tempElt := sliceIn[eltIndex]

		if _, ok := tempElt.(Polynomial); ok {
			return eltIndex
		}

	}
	return -1
}

/*
FindPolynomialLikeObject
Description:

	This function attempts to find any objects of type:
	- Variable
	- Monomial, or
	- Polynomial
	in an input slice.
	It makes use of the Find methods for each variable's type.
*/
func FindPolynomialLikeObject(sliceIn []interface{}) int {
	// Constants

	// Algorithm
	varIndex := FindVariableInSlice(sliceIn)
	monomIndex := FindMonomialInSlice(sliceIn)
	polynomIndex := FindPolynomialInSlice(sliceIn)

	minIndex := varIndex // Assign minimum index to be varIndex

	if minIndex > monomIndex { // If monomIndex is less than the current minimum, then reduce it.
		minIndex = monomIndex
	}

	if minIndex > polynomIndex { // If polynomIndex is less than the current minimum, then reduce it.
		minIndex = polynomIndex
	}

	return minIndex

}
