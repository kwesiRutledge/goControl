package sos_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kwesiRutledge/goControl/sos"
)

/*
TestMonomial_FindMonomialInSlice1
Description:

	Tests to see if there is a Monomial in a slice of objects (objects model the empty interface).
	Example slice contains a Monomial object.
*/
func TestMonomial_FindMonomialInSlice1(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Variable{Name: "Lupe Fiasco"},
		sos.Monomial{},
	}

	// Algorithm
	if sos.FindMonomialInSlice(crazySlice) != 3 {
		t.Errorf("The monomial was found at index %v; expected 3.", sos.FindMonomialInSlice(crazySlice))
	}
}

/*
TestMonomial_FindMonomialInSlice2
Description:

	Tests to see if there is a Monomial in a slice of objects (objects model the empty interface).
	Example slice DOES NOT contain a Monomial object.
*/
func TestMonomial_FindMonomialInSlice2(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Variable{Name: "Lupe Fiasco"},
	}

	// Algorithm
	if sos.FindMonomialInSlice(crazySlice) != -1 {
		t.Errorf("The variable was found at index %v; expected -1.", sos.FindMonomialInSlice(crazySlice))
	}
}

/*
TestMonomial_Multiply1
Description:

	Tests how well our multiplication of a float with a monomial works.
*/
func TestMonomial_Multiply1(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"y"}
	coeff1 := 1.0
	m1 := sos.Monomial{
		Coefficient: coeff1,
		Variables:   []*sos.Variable{&v1, &v2},
		Exponents:   []int{1, 1},
	}

	coeff2 := 3.14

	// Algorithm
	m2, err := m1.Multiply(coeff2)
	if err != nil {
		t.Errorf("There was an error computing the product of a monomial: %v", err)
	}

	if m2.Coefficient != coeff2 {
		t.Errorf("Expected product's coefficient to change to %v; received %v", coeff2, m2.Coefficient)
	}
}

/*
TestMonomial_Multiply2
Description:

	Tests how well our multiplication of a variable with a monomial works.
*/
func TestMonomial_Multiply2(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"y"}
	coeff1 := 1.0
	m1 := sos.Monomial{
		Coefficient: coeff1,
		Variables:   []*sos.Variable{&v1, &v2},
		Exponents:   []int{1, 1},
	}

	// Algorithm
	m2, err := m1.Multiply(v2)
	if err != nil {
		t.Errorf("There was an error computing the product of a monomial: %v", err)
	}

	if m2.Exponents[1] != 2 {
		t.Errorf("Expected product's exponent value to change to 2; received %v", m2.Exponents[1])
	}
}

/*
TestMonomial_Multiply3
Description:

	Tests how well our multiplication of a variable pointer with a monomial works.
*/
func TestMonomial_Multiply3(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"y"}
	coeff1 := 1.0
	m1 := sos.Monomial{
		Coefficient: coeff1,
		Variables:   []*sos.Variable{&v1, &v2},
		Exponents:   []int{1, 1},
	}

	// Algorithm
	m2, err := m1.Multiply(&v2)
	if err != nil {
		t.Errorf("There was an error computing the product of a monomial: %v", err)
	}

	if m2.Exponents[1] != 2 {
		t.Errorf("Expected product's exponent value to change to 2; received %v", m2.Exponents[1])
	}
}

/*
TestMonomial_Multiply4
Description:

	Tests how well our multiplication of a variable pointer with a monomial works WHEN A NON-nil error is present.
	Should produce error.
*/
func TestMonomial_Multiply4(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"y"}
	coeff1 := 1.0
	m1 := sos.Monomial{
		Coefficient: coeff1,
		Variables:   []*sos.Variable{&v1, &v2},
		Exponents:   []int{1, 1},
	}

	errorText := "Test error"

	// Algorithm
	_, err := m1.Multiply(&v2, fmt.Errorf(errorText))
	if err == nil {
		t.Errorf("There was NO error computing the product of this monomial with a variable. There should have been one with text: %v", errorText)
	}

	if !strings.Contains(err.Error(), errorText) {
		t.Errorf("The error returned by multiply should contain our input text \"%v\", but it does not: %v", errorText, err)
	}
}

/*
TestMonomial_Multiply5
Description:

	Tests how well our multiplication of a Monomial with a monomial works.
*/
func TestMonomial_Multiply5(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"y"}
	coeff1 := 2.0
	m1 := sos.Monomial{
		Coefficient: coeff1,
		Variables:   []*sos.Variable{&v1, &v2},
		Exponents:   []int{1, 1},
	}

	// Algorithm
	m2, err := m1.Multiply(m1)
	if err != nil {
		t.Errorf("There was an error multiplying the two monomials: %v", err)
	}

	if m2.Coefficient != m1.Coefficient*m1.Coefficient {
		t.Errorf("Expected for exponent to be %v^2; received %v", m1.Coefficient, m2.Coefficient)
	}

	for varIndex := 0; varIndex < len(m2.Variables); varIndex++ {
		if m2.Exponents[varIndex] != 2*m1.Exponents[varIndex] {
			t.Errorf("The exponent of variable %v was %v; expected 2*%v", m2.Variables[varIndex], m2.Exponents[varIndex], m1.Exponents[varIndex])
		}
	}

}
