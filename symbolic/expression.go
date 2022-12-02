package symbolic

/*
expression.go
Description:
	Defines a polynomial expression which is an INTERFACE. Therefore, it can not be instantiated. A types that implements
	this interface can be implemented though.
*/

// Type Definitions
type Expression interface {
	Multiply(terms ...interface{}) (Expression, error)
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
func FindExpressionObject(sliceIn []interface{}) int {
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
