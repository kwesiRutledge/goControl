package sos

/*
variable.go
Description:
	Defines some basic properties for a simple variable to use in SOS.
*/

import (
	"fmt"
)

type Variable struct {
	Name string
}

// Functions

func GetVariableVector(base_name string, length_in int) ([]Variable, error) {
	// Check to see if size is positive
	if length_in < 1 {
		return []Variable{}, fmt.Errorf("The size input must be greater than 1; received %v", length_in)
	}

	// Create length_in variables and add them to a slice.
	sliceOut := []Variable{}
	for varIndex := 0; varIndex < length_in; varIndex++ {
		sliceOut = append(sliceOut,
			Variable{fmt.Sprintf("%v%v", base_name, varIndex)},
		)
	}

	return sliceOut, nil
}

/*
String
Description:

	Returns the current variable's name as a string.
*/
func (v *Variable) String() string {
	return v.Name
}

/*
Multiply
Description:

	Multiplies the variable v with some value (either a constant or a )
*/
func (v *Variable) Multiply(terms ...interface{}) (Monomial, error) {
	// Constants

	// Error Handling and Recursion
	if len(terms) > 1 {
		// If there are more than one terms given to the multiply function, then
		// - Determine if any of the terms from 1 to end are a variable, monomial or polynomial
		followingIndex := FindPolynomialLikeObject(terms[1:])

		if followingIndex != -1 {
			// If we find a variable, monomial or polynomial,
			// then we must use recursion to compute multiplication.

			// First, check for any errors.
			if followingIndex != 0 { // If error exists, then check it and maybe throw something!
				if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
					return Monomial{}, termAsError // throw error
				}
			}

			// Second, compute sub-product.
			m, err := v.Multiply(terms[0])
			if err != nil {
				return Monomial{}, err
			}

			// Compute Product Among the
			return m.Multiply(terms[followingIndex:])

		} else {
			if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
				return Monomial{}, termAsError // throw error
			}
		}
	}

	// Algorithm
	term1 := terms[0]
	switch term1.(type) {
	case float64:
		return Monomial{
			Coefficient: term1.(float64),
			Variables:   []*Variable{v},
			Exponents:   []int{1},
		}, nil
	case Variable:
		termAsV, _ := term1.(Variable)

		return Monomial{
			Coefficient: 1.0,
			Variables:   []*Variable{v, &termAsV},
			Exponents:   []int{1, 1},
		}, nil

	case *Variable:
		return Monomial{
			Coefficient: 1.0,
			Variables:   []*Variable{v, term1.(*Variable)},
			Exponents:   []int{1, 1},
		}, nil

	default:
		return Monomial{}, fmt.Errorf("The input type %T was not expected!", term1)
	}
}

/*
FindVariableInSlice
Description:

	Looks for a variable in a slice of interfaces.
	Returns -1 if none are found.
*/
func FindVariableInSlice(sliceIn []interface{}) int {
	// Constants

	// Algorithms
	for eltIndex := 0; eltIndex < len(sliceIn); eltIndex++ {
		// check to see if current element is a Variable
		tempElt := sliceIn[eltIndex]

		if _, ok := tempElt.(Variable); ok {
			return eltIndex
		}

	}
	return -1
}

/*
IsEqualTo
Description:

	Determines if two variables are equal to one another using their names.
*/
func (v *Variable) IsEqualTo(vIn *Variable) bool {
	return v.Name == vIn.Name
}

/*
FoundIn
Description:

	Looks for a specific variable in a slice.
	If variable is not found in slice, then return -1.
*/
func (v *Variable) FoundIn(sliceIn []*Variable) int {
	// Iterate through all elements in list
	for varIndex := 0; varIndex < len(sliceIn); varIndex++ {
		tempVar := sliceIn[varIndex]
		if v.IsEqualTo(tempVar) {
			return varIndex
		}
	}

	// If v doesn't match anything, then return -1
	return -1
}
