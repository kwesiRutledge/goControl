package symbolic

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

func GetVariableVector(baseName string, lengthIn int) ([]Variable, error) {
	// Check to see if size is positive
	if lengthIn < 1 {
		return []Variable{}, fmt.Errorf("The size input must be greater than 1; received %v", lengthIn)
	}

	// Create length_in variables and add them to a slice.
	sliceOut := []Variable{}
	for varIndex := 0; varIndex < lengthIn; varIndex++ {
		sliceOut = append(sliceOut,
			Variable{fmt.Sprintf("%v%v", baseName, varIndex)},
		)
	}

	return sliceOut, nil
}

/*
String
Description:

	Returns the current variable's name as a string.
*/
func (v Variable) String() string {
	return v.Name
}

/*
Multiply
Description:

	Multiplies the variable v with some value (either a constant or a )
*/
func (v Variable) Multiply(terms ...interface{}) (Expression, error) {
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
					return &Monomial{}, termAsError // throw error
				}
			}

			// Second, compute sub-product.
			m, err := v.Multiply(terms[0])
			if err != nil {
				return &Monomial{}, err
			}

			// Compute Product Among the
			return m.Multiply(terms[followingIndex:])

		} else {
			if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
				return &Monomial{}, termAsError // throw error
			}
		}
	}

	// Algorithm
	term1 := terms[0]
	switch term1.(type) {
	case float64:
		return &Monomial{
			Coefficient: term1.(float64),
			Variables:   []Variable{v},
			Exponents:   []int{1},
		}, nil
	case Variable:
		termAsV, _ := term1.(Variable)

		if v == termAsV {
			return &Monomial{
				Coefficient: 1.0,
				Variables:   []Variable{v},
				Exponents:   []int{2},
			}, nil
		} else {
			return &Monomial{
				Coefficient: 1.0,
				Variables:   []Variable{v, termAsV},
				Exponents:   []int{1, 1},
			}, nil
		}

	case *Variable:
		termAsVPointer, _ := term1.(*Variable)

		if v == (*termAsVPointer) {
			return &Monomial{
				Coefficient: 1.0,
				Variables:   []Variable{v},
				Exponents:   []int{2},
			}, nil
		} else {
			return &Monomial{
				Coefficient: 1.0,
				Variables:   []Variable{v, *termAsVPointer},
				Exponents:   []int{1, 1},
			}, nil
		}

	case Monomial:
		termAsMonom, _ := term1.(Monomial)
		monomialOut := termAsMonom.Copy()

		vIndex := v.FoundIn(monomialOut.Variables)
		if vIndex != -1 { // v is already in the monomial
			monomialOut.Exponents[vIndex] += 1
		} else {
			monomialOut.Variables = append(monomialOut.Variables, v)
			monomialOut.Exponents = append(monomialOut.Exponents, 1)
		}

		return &monomialOut, nil

	case *Monomial:
		termAsMonom, _ := term1.(*Monomial)
		monomialOut := termAsMonom.Copy()

		vIndex := v.FoundIn(monomialOut.Variables)
		if vIndex != -1 { // v is already in the monomial
			monomialOut.Exponents[vIndex] += 1
		} else {
			monomialOut.Variables = append(monomialOut.Variables, v)
			monomialOut.Exponents = append(monomialOut.Exponents, 1)
		}

		return &monomialOut, nil

	default:
		return &Monomial{}, fmt.Errorf("The input type %T was not expected!", term1)
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
func (v Variable) IsEqualTo(vIn Variable) bool {
	return v.Name == vIn.Name
}

/*
FoundIn
Description:

	Looks for a specific variable in a slice.
	If variable is not found in slice, then return -1.
*/
func (v Variable) FoundIn(sliceIn []Variable) int {
	// Iterate through all elements in list
	for varIndex := 0; varIndex < len(sliceIn); varIndex++ {
		tempVar := sliceIn[varIndex]
		if v == tempVar {
			return varIndex
		}
	}

	// If v doesn't match anything, then return -1
	return -1
}

/*
Sum
Description:

	Computes the sum between a variable and a symbolic expression.
*/
func (v Variable) Sum(terms ...interface{}) (Expression, error) {
	// Constants

	// Error Handling and Recursion
	if len(terms) > 1 {
		// If there are more than one terms given to the multiply function, then
		// - Determine if any of the terms from 1 to end are a variable, monomial or polynomial
		followingIndex := FindExpressionObject(terms[1:])

		if followingIndex != -1 {
			// If we find a variable, monomial or polynomial,
			// then we must use recursion to compute multiplication.

			// First, check for any errors.
			if followingIndex != 0 { // If error exists, then check it and maybe throw something!
				if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
					return &Monomial{}, termAsError // throw error
				}
			}

			// Second, compute sub-product.
			m, err := v.Multiply(terms[0])
			if err != nil {
				return &Monomial{}, err
			}

			// Compute Product Among the
			return m.Multiply(terms[followingIndex:])

		} else {
			if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
				return &Monomial{}, termAsError // throw error
			}
		}
	}

	// Algorithm
	term1 := terms[0]
	switch term1.(type) {
	case float64:
		term1AsFloat, _ := term1.(float64)
		return &Polynomial{
			Monomials: []Monomial{
				{
					Variables:   []Variable{},
					Exponents:   []int{},
					Coefficient: term1AsFloat,
				},
				{
					Variables:   []Variable{v},
					Exponents:   []int{1},
					Coefficient: 1.0,
				},
			},
		}, nil
	case Variable:
		termAsV, _ := term1.(Variable)

		if v == termAsV {
			return Monomial{
				Coefficient: 2.0,
				Variables:   []Variable{v},
				Exponents:   []int{1},
			}, nil
		} else {
			return Polynomial{
				Monomials: []Monomial{
					{
						Variables:   []Variable{v},
						Exponents:   []int{1},
						Coefficient: 1.0,
					},
					{
						Variables:   []Variable{termAsV},
						Exponents:   []int{1},
						Coefficient: 1.0,
					},
				},
			}, nil
		}

	case *Variable:
		return &Monomial{
			Coefficient: 1.0,
			Variables:   []Variable{v, *term1.(*Variable)},
			Exponents:   []int{1, 1},
		}, nil
	case Monomial:
		termAsMonom, _ := term1.(Monomial)
		monomialOut := termAsMonom.Copy()

		vIndex := v.FoundIn(monomialOut.Variables)
		if vIndex != -1 { // v is already in the monomial
			monomialOut.Exponents[vIndex] += 1
		} else {
			monomialOut.Variables = append(monomialOut.Variables, v)
			monomialOut.Exponents = append(monomialOut.Exponents, 1)
		}

		return &monomialOut, nil

	case *Monomial:
		termAsMonom, _ := term1.(*Monomial)
		monomialOut := termAsMonom.Copy()

		vIndex := v.FoundIn(monomialOut.Variables)
		if vIndex != -1 { // v is already in the monomial
			monomialOut.Exponents[vIndex] += 1
		} else {
			monomialOut.Variables = append(monomialOut.Variables, v)
			monomialOut.Exponents = append(monomialOut.Exponents, 1)
		}

		return &monomialOut, nil

	default:
		return &Monomial{}, fmt.Errorf("The input type %T was not expected!", term1)
	}
}
