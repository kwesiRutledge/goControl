package symbolic_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kwesiRutledge/goControl/symbolic"
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
		symbolic.Variable{Name: "Lupe Fiasco"},
		symbolic.Monomial{},
	}

	// Algorithm
	if symbolic.FindMonomialInSlice(crazySlice) != 3 {
		t.Errorf("The monomial was found at index %v; expected 3.", symbolic.FindMonomialInSlice(crazySlice))
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
		symbolic.Variable{Name: "Lupe Fiasco"},
	}

	// Algorithm
	if symbolic.FindMonomialInSlice(crazySlice) != -1 {
		t.Errorf("The variable was found at index %v; expected -1.", symbolic.FindMonomialInSlice(crazySlice))
	}
}

/*
TestMonomial_Multiply1
Description:

	Tests how well our multiplication of a float with a monomial works.
*/
func TestMonomial_Multiply1(t *testing.T) {
	// Constants
	v1 := symbolic.Variable{"x"}
	v2 := symbolic.Variable{"y"}
	coeff1 := 1.0
	m1 := symbolic.Monomial{
		Coefficient: coeff1,
		Variables:   []symbolic.Variable{v1, v2},
		Exponents:   []int{1, 1},
	}

	coeff2 := 3.14

	// Algorithm
	prodOut, err := m1.Multiply(coeff2)
	if err != nil {
		t.Errorf("There was an error computing the product of a monomial: %v", err)
	}

	m2, ok := prodOut.(symbolic.Monomial)
	if (m2.Coefficient != coeff2) || (!ok) {
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
	v1 := symbolic.Variable{"x"}
	v2 := symbolic.Variable{"y"}
	coeff1 := 1.0
	m1 := symbolic.Monomial{
		Coefficient: coeff1,
		Variables:   []symbolic.Variable{v1, v2},
		Exponents:   []int{1, 1},
	}

	// Algorithm
	prodOut, err := m1.Multiply(v2)
	if err != nil {
		t.Errorf("There was an error computing the product of a monomial: %v", err)
	}

	m2, ok := prodOut.(symbolic.Monomial)
	if !ok {
		t.Errorf("The output of the Multiply was not of type Monomial; received %T", prodOut)
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
	v1 := symbolic.Variable{"x"}
	v2 := symbolic.Variable{"y"}
	coeff1 := 1.0
	m1 := symbolic.Monomial{
		Coefficient: coeff1,
		Variables:   []symbolic.Variable{v1, v2},
		Exponents:   []int{1, 1},
	}

	// Algorithm
	prodOut, err := m1.Multiply(v2)
	if err != nil {
		t.Errorf("There was an error computing the product of a monomial: %v", err)
	}

	m2, ok := prodOut.(symbolic.Monomial)
	if !ok {
		t.Errorf("The output of the Multiply was not of type Monomial; received %T", prodOut)
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
	v1 := symbolic.Variable{"x"}
	v2 := symbolic.Variable{"y"}
	coeff1 := 1.0
	m1 := symbolic.Monomial{
		Coefficient: coeff1,
		Variables:   []symbolic.Variable{v1, v2},
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

	Tests how well our multiplication of a Monomial with a Monomial works.
*/
func TestMonomial_Multiply5(t *testing.T) {
	// Constants
	v1 := symbolic.Variable{"x"}
	v2 := symbolic.Variable{"y"}
	coeff1 := 2.0
	m1 := symbolic.Monomial{
		Coefficient: coeff1,
		Variables:   []symbolic.Variable{v1, v2},
		Exponents:   []int{1, 1},
	}

	// Algorithm
	prodOut, err := m1.Multiply(m1)
	if err != nil {
		t.Errorf("There was an error multiplying the two monomials: %v", err)
	}

	m2, _ := prodOut.(symbolic.Monomial)
	if m2.Coefficient != m1.Coefficient*m1.Coefficient {
		t.Errorf("Expected for exponent to be %v^2; received %v", m1.Coefficient, m2.Coefficient)
	}

	for varIndex := 0; varIndex < len(m2.Variables); varIndex++ {
		if m2.Exponents[varIndex] != 2*m1.Exponents[varIndex] {
			t.Errorf("The exponent of variable %v was %v; expected 2*%v", m2.Variables[varIndex], m2.Exponents[varIndex], m1.Exponents[varIndex])
		}
	}

}

/*
TestMonomial_Multiply6
Description:

	Tests how well our multiplication of a Monomial with a Monomial Pointer (*Monomial) works.
*/
func TestMonomial_Multiply6(t *testing.T) {
	// Constants
	v1 := symbolic.Variable{"x"}
	v2 := symbolic.Variable{"y"}
	coeff1 := 2.0
	m1 := symbolic.Monomial{
		Coefficient: coeff1,
		Variables:   []symbolic.Variable{v1, v2},
		Exponents:   []int{1, 1},
	}
	coeff2 := 3.0
	m2 := symbolic.Monomial{
		Coefficient: coeff2,
		Variables:   []symbolic.Variable{v1, v2},
		Exponents:   []int{2, 3},
	}

	// Algorithm
	prodOut, err := m1.Multiply(m2)
	if err != nil {
		t.Errorf("There was an error multiplying the two monomials: %v", err)
	}

	m3, ok := prodOut.(symbolic.Monomial)
	if !ok {
		t.Errorf("The output of the Multiply was not of type Monomial; received %T", prodOut)
	}

	if m3.Coefficient != m1.Coefficient*m2.Coefficient {
		t.Errorf("Expected for exponent to be %v^2; received %v", m1.Coefficient, m2.Coefficient)
	}

	for varIndex := 0; varIndex < len(m3.Variables); varIndex++ {
		if m3.Exponents[varIndex] != m1.Exponents[varIndex]+m2.Exponents[varIndex] {
			t.Errorf("The exponent of variable %v was %v; expected 2*%v", m2.Variables[varIndex], m2.Exponents[varIndex], m1.Exponents[varIndex])
		}
	}

}

/*
TestMonomial_String1
Description:

	Tests to see if a simple monomial string is what we expect it to be.
*/
func TestMonomial_String1(t *testing.T) {
	// Constants
	x := symbolic.Variable{"x"}
	y := symbolic.Variable{"y"}

	monom1 := symbolic.Monomial{
		Coefficient: 2.0,
		Variables:   []symbolic.Variable{x, y},
		Exponents:   []int{3, 7},
	}

	// Algorithm
	strOut := monom1.String()
	if !strings.Contains(strOut, "2") {
		t.Errorf("The monomial string %v does not contain the coefficient %v!", monom1.String(), monom1.Coefficient)
	}

	for varIndex := 0; varIndex < len(monom1.Variables); varIndex++ {
		if !strings.Contains(strOut, fmt.Sprintf("(%v)^%v", monom1.Variables[varIndex].String(), monom1.Exponents[varIndex])) {
			t.Errorf("The monomial string %v does not contain the variable %v with exponent %v!", monom1.String(), monom1.Variables[varIndex], monom1.Exponents[varIndex])
		}
	}

}

/*
TestMonomial_String2
Description:

	Tests to see if a simple monomial string is what we expect it to be.
	Contains an exponent of 1
*/
func TestMonomial_String2(t *testing.T) {
	// Constants
	x := symbolic.Variable{"x"}
	y := symbolic.Variable{"y"}

	monom1 := symbolic.Monomial{
		Coefficient: 2.0,
		Variables:   []symbolic.Variable{x, y},
		Exponents:   []int{3, 1},
	}

	// Algorithm
	strOut := monom1.String()
	if !strings.Contains(strOut, "2") {
		t.Errorf("The monomial string %v does not contain the coefficient %v!", monom1.String(), monom1.Coefficient)
	}

	if !strings.Contains(strOut, fmt.Sprintf("(%v)^%v", monom1.Variables[0].String(), monom1.Exponents[0])) {
		t.Errorf("The monomial string %v does not contain the variable %v with coefficient %v!", monom1.String(), monom1.Variables[0], monom1.Exponents[0])
	}
	if strings.Contains(strOut, fmt.Sprintf("(%v)", monom1.Variables[1].String())) {
		t.Errorf("The monomial string %v should not contain parentheses around variable %v!", monom1.String(), monom1.Variables[1])
	}

	if !strings.Contains(strOut, fmt.Sprintf("%v", monom1.Variables[1].String())) {
		t.Errorf("The monomial string %v should contain just the name of variable %v!", monom1.String(), monom1.Variables[1])
	}

}
