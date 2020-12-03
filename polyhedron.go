/*
   polyhedron.go
   Description:
       An implementation "like" MPT3.0's implementation of the Polyhedron class.
*/

package goControl

import (
	"fmt"
	la "gosl/la"
	"errors"
)

type Polyhedron struct {
	A la.Matrix
	b la.Vector
}

/*
	GetPolyhedron()
	Is This Function Necessary?
*/
func GetPolyhedron() {
	fmt.Println("Is this working?")
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

	return polyhedronIn.A.N // The Matrix's .M field is the number of rows and the .N file is the number of columns

}

/*
	AreAorbUndefined
	Description:
		Returns true if either the matrix A or the vector b are undefined for the input Polyhedron.
*/
func (polyhedronIn Polyhedron) AreAorbUndefined() bool {

	//Check if A has zero dimensions
	if (polyhedronIn.A.M == 0) || (polyhedronIn.A.N == 0) || (len(polyhedronIn.A.Data) == 0) {
		return true
	}

	//Check if b has zero dimensions
	if (len(polyhedronIn.b) == 0) {
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

	//Check to see if the Polyhedron has been defined
	if polyhedronIn.AreAorbUndefined() {
		return errors.New("The A or b property of the input Polyhedron is not defined.")
	}

	//Check to see if the dimensions of the A and b matrices are compatible!
	if ( polyhedronIn.A.M != len(polyhedronIn.b) ) {
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

func (polyhedronIn Polyhedron) Contains( targetObject interface{} ) bool {

	// Check the type of the input object
	switch v := targetObject.(type) {
	case la.Vector:
		print(v)
		return true
	default:
		//
		fmt.Println("The target object is not one of the allowed types.")
		return false
	}

}