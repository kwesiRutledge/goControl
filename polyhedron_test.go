/*
   polyhedron_test.go
   Description:
	   Tests for the Polyhedron() type defined in polyhedron.go.
*/

package goControl

import (
	"testing"
	"gosl/la"
	// "errors"
	// "fmt"
)

/*
	TestPolyhedronConstructor1
	Description:
		Attempts to make a simple Polyhedron and checks that the values were properly initialized.
*/
func TestPolyhedronConstructor1(t *testing.T) {
	//Practices using the simple A,b constructor

	A := la.NewMatrixDeep2([][]float64{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2},
	})
	b := la.NewMatrixDeep2([][]float64{
		{1},
		{2},
		{3},
	})
	//fmt.Println(b)

	//t.Errorf("Type of A = %T",A)

	//Use Constructor
	poly1 := Polyhedron{
		A: A,
		b: b,
	}

	//Test object	
	if poly1.A != A {
		t.Errorf("poly1.A = %v; want %v", poly1.A, A)
	}

	if poly1.b != b {
		t.Errorf("poly1.b = %v; want %v", poly1.b, b)
	}
}

/*
	TestPolyhedronConstructor2
	Description:
		Attempts to make an empty Polyhedron and checks that the values were properly initialized.
*/
func TestPolyhedronConstructor2(t *testing.T) {
	//Practices using the simple A,b constructor

	//Use Constructor
	poly1 := Polyhedron{}

	// t.Errorf("poly1.A = %v, poly1.b = %v",poly1.A,poly1.b)

	//Test object	
	if poly1.A != nil {
		t.Errorf("poly1.A = %v; want nil", poly1.A)
	}

	if poly1.b != nil {
		t.Errorf("poly1.b = %v; want nil", poly1.b)
	}
}

/*
	TestPolyhedronDimension1
	Description:
		Tests the Dimension function when applied to a reasonably defined Polyhedron object.
*/
func TestPolyhedronDimension1(t *testing.T) {
	//Create a simple Polyhedron

	A := la.NewMatrixDeep2([][]float64{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2},
	})
	b := la.NewMatrixDeep2([][]float64{
		{1},
		{2},
		{3},
	})

	//Use Constructor
	poly1 := Polyhedron{
		A: A,
		b: b,
	}

	//Testing the Dimension function
	dim := poly1.Dimension()
	if dim != 3 {
		t.Errorf("Dimension of Polytope = %v; want 3", dim)
	}

}

/*
	TestPolyhedronCheck1
	Description:
		Tests the Check() function which determines if a Polyhedron is reasonable enough to perform computations with.
*/
func TestPolyhedronCheck1(t *testing.T) {
	//Construct empty Polyhedron
	poly1 := Polyhedron{}

	//Check it
	err := poly1.Check()
	//desiredError := errors.New("The A or b property of the input Polyhedron is not defined.")
	if err == nil {
		t.Errorf("err = %v; want \"The A or b property of the input Polyhedron is not defined.\"",err)
	}
}