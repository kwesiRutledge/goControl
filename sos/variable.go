package sos

/*
variable.go
Description:
	Defines some basic properties for a simple variable to use in SOS.
*/

import "fmt"

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
func (v *Variable) Multiply(term interface{}) (Monomial, error) {
	// Constants

	// Algorithm
	switch term.(type) {
	case float64:
		return Monomial{
			Coefficient: term.(float64),
			Variables:   []*Variable{v},
			Exponents:   []int{1},
		}, nil
	case Variable:
		termAsV, _ := term.(Variable)

		return Monomial{
			Coefficient: 1.0,
			Variables:   []*Variable{v, &termAsV},
			Exponents:   []int{1, 1},
		}, nil

	case *Variable:
		return Monomial{
			Coefficient: 1.0,
			Variables:   []*Variable{v, term.(*Variable)},
			Exponents:   []int{1, 1},
		}, nil

	default:
		return Monomial{}, fmt.Errorf("The input type %T was not expected!", term)
	}
}
