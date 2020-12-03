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

	A := *la.NewMatrixDeep2([][]float64{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2},
	})
	b := la.NewVectorSlice([]float64{1,2,3,})
	//fmt.Println(b)

	//t.Errorf("Type of A = %T",A)

	//Use Constructor
	poly1 := Polyhedron{
		A: A,
		b: b,
	}

	//Test object	
	//- Compare poly1.A with A
	for rowIndex := 0 ; rowIndex < poly1.A.M ; rowIndex++ {
		for columnIndex := 0 ; columnIndex < poly1.A.N ; columnIndex++ {
			if (&(poly1.A)).Get(rowIndex,columnIndex) != (&A).Get(rowIndex,columnIndex){
				t.Errorf("poly1.A[%v,%v] = %v; want %v", rowIndex, columnIndex , (&(poly1.A)).Get(rowIndex,columnIndex), (&A).Get(rowIndex,columnIndex),)
			}
		}
	}
	// if poly1.A != A {
	// 	t.Errorf("poly1.A = %v; want %v", poly1.A, A)
	// }

	for vectorIndex := 0 ; vectorIndex < len(poly1.b) ; vectorIndex++ {
		if poly1.b[vectorIndex] != b[vectorIndex] {
			t.Errorf("poly1.A[%v] = %v; want %v", vectorIndex, poly1.b[vectorIndex], b[vectorIndex])
		}
	}

	// if poly1.b != b {
	// 	t.Errorf("poly1.b = %v; want %v", poly1.b, b)
	// }
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
	//t.Errorf("poly1.A = %v, poly1.b = %v",poly1.A,poly1.b)

	//When initialized empty we expect for an empty matrix A to be created
	//	A = { 0 0 [] } i.e. A.M = 0 (0 rows), A.N = 0 (0 columns), A.Data = [] (no entries)
	//	b = [] i.e. b is a vector with no entries

	//Test object	
	if (poly1.A.M != 0) || (poly1.A.N != 0) || (len(poly1.A.Data) != 0) {
		t.Errorf("(poly1.A.M , poly1.A.N , len(poly1.A.Data) ) = (%v,%v,%v); want (0 0,0)", poly1.A.M,poly1.A.N,len(poly1.A.Data))
	}

	if (len(poly1.b) != 0) {
		t.Errorf("len(poly1.b) = %v; want 0", len(poly1.b) )
	}
}

/*
	TestPolyhedronDimension1
	Description:
		Tests the Dimension function when applied to a reasonably defined Polyhedron object.
*/
func TestPolyhedronDimension1(t *testing.T) {
	//Create a simple Polyhedron

	A := *la.NewMatrixDeep2([][]float64{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2},
	})
	b := la.NewVectorSlice([]float64{1,2,3,})

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
	//t.Errorf("poly1.A = %v, poly1.b = %v",poly1.A,poly1.b)

	//desiredError := errors.New("The A or b property of the input Polyhedron is not defined.")
	if err == nil {
		t.Errorf("err = %v; want \"The A or b property of the input Polyhedron is not defined.\"",err)
	}
}

/*
	TestPolyhedronCheck2
	Description:
		Tests the Check() function which determines if a Polyhedron is reasonable enough to perform computations with.
		This example should be valid.
*/
func TestPolyhedronCheck2(t *testing.T) {
	//Create a simple Polyhedron

	A := *la.NewMatrixDeep2([][]float64{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2},
	})

	A2 := *la.NewMatrixDeep2([][]float64{
		{3, 0, 1},
		{0, 1, 0},
		{0, 0, -1},
	}) 

	AProduct := *la.NewMatrix(A.M,A.N)
	la.MatMatMul(&AProduct,1.0,&A,&A2)

	b := la.NewVectorSlice([]float64{1,2,3,})

	//Use Constructor
	poly1 := Polyhedron{
		A: AProduct,
		b: b,
	}

	//Check it
	err := poly1.Check()
	//desiredError := errors.New("The A or b property of the input Polyhedron is not defined.")
	if err != nil {
		t.Errorf("err = %v; want nil",err)
	}

}