package sos_test

import (
	"testing"

	"github.com/kwesiRutledge/goControl/sos"
)

/*
TestPolynomial_FindPolynomialInSlice1
Description:

	Tests to see if there is a Polynomial in a slice of objects (objects model the empty interface).
	Example slice contains a Polynomial object.
*/
func TestPolynomial_FindPolynomialInSlice1(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Variable{Name: "Lupe Fiasco"},
		sos.Monomial{},
		sos.Polynomial{},
	}

	// Algorithm
	if sos.FindPolynomialInSlice(crazySlice) != 4 {
		t.Errorf("The monomial was found at index %v; expected 4.", sos.FindPolynomialInSlice(crazySlice))
	}
}

/*
TestPolynomial_FindPolynomialInSlice2
Description:

	Tests to see if there is a Polynomial in a slice of objects (objects model the empty interface).
	Example slice DOES NOT contain a Polynomial object.
*/
func TestPolynomial_FindPolynomialInSlice2(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Variable{Name: "Lupe Fiasco"},
		sos.Monomial{},
	}

	// Algorithm
	if sos.FindPolynomialInSlice(crazySlice) != -1 {
		t.Errorf("The variable was found at index %v; expected -1.", sos.FindPolynomialInSlice(crazySlice))
	}
}

/*
TestPolynomial_FindPolynomialLikeObject1
Description:

	Tests to see if FindPolynomialLikeObject() finds the Variable's index in a slice, when that is the first
	thing we find.
*/
func TestPolynomial_FindPolynomialLikeObject1(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Variable{Name: "Lupe Fiasco"},
		sos.Monomial{},
		sos.Polynomial{},
	}

	// Algorithm
	if sos.FindPolynomialLikeObject(crazySlice) != 2 {
		t.Errorf("The monomial was found at index %v; expected 2.", sos.FindPolynomialLikeObject(crazySlice))
	}
}

/*
TestPolynomial_FindPolynomialLikeObject2
Description:

	Tests to see if FindPolynomialLikeObject() finds the Monomial's index in a slice, when that is the first
	thing we find.
*/
func TestPolynomial_FindPolynomialLikeObject2(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Monomial{},
		sos.Variable{Name: "Lupe Fiasco"},
		sos.Polynomial{},
	}

	// Algorithm
	if sos.FindPolynomialLikeObject(crazySlice) != 2 {
		t.Errorf("The monomial was found at index %v; expected 2.", sos.FindPolynomialLikeObject(crazySlice))
	}
}

/*
TestPolynomial_FindPolynomialLikeObject3
Description:

	Tests to see if FindPolynomialLikeObject() finds the Polynomial's index in a slice, when that is the first
	thing we find.
*/
func TestPolynomial_FindPolynomialLikeObject3(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Polynomial{},
		sos.Monomial{},
		sos.Variable{Name: "Lupe Fiasco"},
	}

	// Algorithm
	if sos.FindPolynomialLikeObject(crazySlice) != 2 {
		t.Errorf("The monomial was found at index %v; expected 2.", sos.FindPolynomialLikeObject(crazySlice))
	}
}
