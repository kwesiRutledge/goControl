package symbolic_test

import (
	"testing"

	"github.com/kwesiRutledge/goControl/symbolic"
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
		symbolic.Variable{Name: "Lupe Fiasco"},
		symbolic.Monomial{},
		symbolic.Polynomial{},
	}

	// Algorithm
	if symbolic.FindPolynomialInSlice(crazySlice) != 4 {
		t.Errorf("The monomial was found at index %v; expected 4.", symbolic.FindPolynomialInSlice(crazySlice))
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
		symbolic.Variable{Name: "Lupe Fiasco"},
		symbolic.Monomial{},
	}

	// Algorithm
	if symbolic.FindPolynomialInSlice(crazySlice) != -1 {
		t.Errorf("The variable was found at index %v; expected -1.", symbolic.FindPolynomialInSlice(crazySlice))
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
		symbolic.Variable{Name: "Lupe Fiasco"},
		symbolic.Monomial{},
		symbolic.Polynomial{},
	}

	// Algorithm
	if symbolic.FindPolynomialLikeObject(crazySlice) != 2 {
		t.Errorf("The monomial was found at index %v; expected 2.", symbolic.FindPolynomialLikeObject(crazySlice))
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
		symbolic.Monomial{},
		symbolic.Variable{Name: "Lupe Fiasco"},
		symbolic.Polynomial{},
	}

	// Algorithm
	if symbolic.FindPolynomialLikeObject(crazySlice) != 2 {
		t.Errorf("The monomial was found at index %v; expected 2.", symbolic.FindPolynomialLikeObject(crazySlice))
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
		symbolic.Polynomial{},
		symbolic.Monomial{},
		symbolic.Variable{Name: "Lupe Fiasco"},
	}

	// Algorithm
	if symbolic.FindPolynomialLikeObject(crazySlice) != 2 {
		t.Errorf("The monomial was found at index %v; expected 2.", symbolic.FindPolynomialLikeObject(crazySlice))
	}
}
