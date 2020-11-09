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
	A *la.Matrix
	b *la.Matrix
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
func (poly_in Polyhedron) Dimension() int {

	if poly_in.A == nil {
		return -1 //, errors.New("The A property of the input Polyhedron is not defined.")
	}

	return poly_in.A.N // The Matrix's .M field is the number of rows and the .N file is the number of columns

}

/*
	Check
	Description:
		Returns true if the Polytope is correct in all respects,
		returns false otherwise.
*/
func (poly_in Polyhedron) Check() error {

	//Check to see if the Polyhedron has been defined
	if ((poly_in.A == nil) || (poly_in.b == nil)) {
		return errors.New("The A or b property of the input Polyhedron is not defined.")
	}

	//Check to see if the dimensions of the A and b matrices are compatible!
	if ( poly_in.A.M != poly_in.b.M ) {
		return errors.New("The dimensions of the A and b matrices are not compatible!")
	}

	//If nothing is wrong, return nil
	return nil
}
