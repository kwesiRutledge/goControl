/*
   polyhedron.go
   Description:
       An implementation "like" MPT3.0's implementation of the Polyhedron class.
*/

package goControl

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type Polyhedron struct {
	A mat.Matrix
	b mat.Vector
}

func (p *Polyhedron) Get_A() mat.Matrix {
	return p.A
}

func (p *Polyhedron) Get_b() mat.Vector {
	return p.b
}

/*
GetPolyhedron()
Is This Function Necessary?
*/
func GetPolyhedron(AIn mat.Matrix, bIn mat.Vector) Polyhedron {
	return Polyhedron{
		A: AIn,
		b: bIn,
	}
}

/*
Dimension
Description:

	Obtains the dimension of the input Polyhedron.
*/
func (polyhedronIn Polyhedron) Dimension() int {

	if polyhedronIn.Check() != nil {
		//If this is not a valid Polyhedron, then return garbage.
		return -1
	}

	_, N := polyhedronIn.A.Dims()

	return N // The Matrix's .M field is the number of rows and the .N file is the number of columns

}

/*
AreAorbUndefined
Description:

	Returns true if either the matrix A or the vector b are undefined for the input Polyhedron.
*/
func (polyhedronIn Polyhedron) AreAorbUndefined() bool {
	// Constants
	M, N := polyhedronIn.A.Dims()

	//Check if A has zero dimensions
	if (M == 0) || (N == 0) {
		return true
	}

	//Check if b has zero dimensions
	if polyhedronIn.b.Len() == 0 {
		return true
	}

	// If you make it this far, then return false!
	return false
}

/*
Check
Description:

	Returns true if the Polytope is correct in all respects,
	returns false otherwise.
*/
func (polyhedronIn Polyhedron) Check() error {
	// Constants
	M, _ := polyhedronIn.A.Dims()

	//Check to see if the Polyhedron has been defined
	if polyhedronIn.AreAorbUndefined() {
		return errors.New("The A or b property of the input Polyhedron is not defined.")
	}

	//Check to see if the dimensions of the A and b matrices are compatible!
	if M != polyhedronIn.b.Len() {
		return errors.New("The dimensions of the A and b matrices are not compatible!")
	}

	//If nothing is wrong, return nil
	return nil
}

/*
	Contains
	Description:
		Returns true if:
		- the provided point is in the the target polyhedron, ...

*/

func (polyhedronIn Polyhedron) Contains(targetObject interface{}) bool {

	// Check the type of the input object
	switch v := targetObject.(type) {
	case mat.Vector:
		print(v)
		return true
	default:
		//
		fmt.Println("The target object is not one of the allowed types.")
		return false
	}

}
