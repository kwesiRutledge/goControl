package sos_test

import (
	"testing"

	"github.com/kwesiRutledge/goControl/sos"
)

func TestVariable_Name1(t *testing.T) {
	// Create a variable
	sos_var1 := sos.Variable{
		Name: "Christopher",
	}

	// Test Name
	if sos_var1.Name != "Christopher" {
		t.Errorf("Expected variable name to be Christopher; received %v", sos_var1.Name)
	}
}

func TestVariable_GetVariableVector1(t *testing.T) {
	// Constants
	baseName1 := "x"
	vectorDim1 := 5

	// Algorithm
	vectorOfVariables, err := sos.GetVariableVector(baseName1, vectorDim1)
	if err != nil {
		t.Errorf("There was an error collecting the variable vector: %v", err)
	}

	if len(vectorOfVariables) != vectorDim1 {
		t.Errorf("The vector of variables was of length %v; expected %v", len(vectorOfVariables), vectorDim1)
	}

}

/*
TestVariable_String1
Description:

	Tests how well the computation of a string for a variable works.
	I think this means that this variable can be used in Stringer interface code.
*/
func TestVariable_String1(t *testing.T) {
	// Constants
	sos_var1 := sos.Variable{
		Name: "James",
	}

	// Test
	if sos_var1.String() != "James" {
		t.Errorf("Expected for variable to have string form \"James\"'; instead received \"%v\"", sos_var1.String())
	}
}

/*
TestVariable_Multiply1
Description:

	Tests how well the multiplication term works with float inputs.
*/
func TestVariable_Multiply1(t *testing.T) {
	// Constants
	coeff_val := 1.0
	sos_var1 := sos.Variable{
		Name: "James",
	}

	// Test
	prod1, err := sos_var1.Multiply(coeff_val)
	if err != nil {
		t.Errorf("There was an error multiplying a variable with a coefficient! %v", err)
	}

	if prod1.Coefficient != coeff_val {
		t.Errorf("The coefficient was %v; Expected %v", prod1.Coefficient, coeff_val)
	}

	if prod1.Variables[0] != &sos_var1 {
		t.Errorf("The address of variable should be %v; Received %v", &sos_var1, prod1.Variables[0])
	}
}

/*
TestVariable_Multiply2
Description:

	Tests how well the multiplication term works with variable inputs.
*/
func TestVariable_Multiply2(t *testing.T) {
	// Constants
	sos_var1 := sos.Variable{
		Name: "James",
	}

	sos_var2 := sos.Variable{
		Name: "Madison",
	}

	// Test
	prod1, err := sos_var1.Multiply(sos_var2)
	if err != nil {
		t.Errorf("There was an error multiplying a variable with a coefficient! %v", err)
	}

	if prod1.Coefficient != 1.0 {
		t.Errorf("The coefficient was %v; Expected 1", prod1.Coefficient)
	}

	if prod1.Variables[0] != &sos_var1 {
		t.Errorf("The address of variable should be %v; Received %v", &sos_var1, prod1.Variables[0])
	}

	if prod1.Variables[1].Name != sos_var2.Name {
		t.Errorf("The address of variable should be %v; Received %v", &sos_var2, prod1.Variables[1])
	}
}

/*
TestVariable_Multiply3
Description:

	Tests how well the multiplication term works with variable POINTER inputs.
*/
func TestVariable_Multiply3(t *testing.T) {
	// Constants
	sos_var1 := sos.Variable{
		Name: "James",
	}

	sos_var2 := sos.Variable{
		Name: "Madison",
	}

	// Test
	prod1, err := sos_var1.Multiply(&sos_var2)
	if err != nil {
		t.Errorf("There was an error multiplying a variable with a coefficient! %v", err)
	}

	if prod1.Coefficient != 1.0 {
		t.Errorf("The coefficient was %v; Expected 1", prod1.Coefficient)
	}

	if prod1.Variables[0] != &sos_var1 {
		t.Errorf("The address of variable should be %v; Received %v", &sos_var1, prod1.Variables[0])
	}

	if prod1.Variables[1] != &sos_var2 {
		t.Errorf("The address of variable should be %v; Received %v", &sos_var2, prod1.Variables[1])
	}
}

/*
TestVariable_FindVariableInSlice1
Description:

	Tests to see if there is a variable in a slice of objects (objects model the empty interface).
	Example slice contains a Variable object.
*/
func TestVariable_FindVariableInSlice1(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Variable{Name: "Lupe Fiasco"},
	}

	// Algorithm
	if sos.FindVariableInSlice(crazySlice) != 2 {
		t.Errorf("The variable was found at index %v; expected 2.", sos.FindVariableInSlice(crazySlice))
	}
}

/*
TestVariable_FindVariableInSlice2
Description:

	Tests to see if there is a variable in a slice of objects (objects model the empty interface).
	Example slice DOES NOT contain a Variable object.
*/
func TestVariable_FindVariableInSlice2(t *testing.T) {
	// Constants

	crazySlice := []interface{}{
		"total",
		2.3,
		sos.Monomial{},
	}

	// Algorithm
	if sos.FindVariableInSlice(crazySlice) != -1 {
		t.Errorf("The variable was found at index %v; expected -1.", sos.FindVariableInSlice(crazySlice))
	}
}

/*
TestVariable_IsEqualTo1
Description:

	Tests to see if the toolbox recognizes that two variables with THE SAME NAME are considered equal.
*/
func TestVariable_IsEqualTo1(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"x"}

	// Algorithm
	if !v1.IsEqualTo(&v2) {
		t.Errorf("The two variables %v and %v were declared to be different!", v1, v2)
	}
}

/*
TestVariable_IsEqualTo2
Description:

	Tests to see if the toolbox recognizes that two variables with DIFFERENT NAMES are considered not equal.
*/
func TestVariable_IsEqualTo2(t *testing.T) {
	// Constants
	v1 := sos.Variable{"x"}
	v2 := sos.Variable{"y"}

	// Algorithm
	if v1.IsEqualTo(&v2) {
		t.Errorf("The two variables %v and %v were declared to be the same!", v1, v2)
	}
}
