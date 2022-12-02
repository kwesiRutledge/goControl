package symbolic

import "fmt"

/*
monomial.go
Description:
	A list of files that is relevant/helpful for the
*/

/*
Type Definitions
*/

type Monomial struct {
	Variables   []Variable
	Exponents   []int
	Coefficient float64
}

//==========
// Functions
//==========

/*
Multiply
Description:
*/
func (m Monomial) Multiply(terms ...interface{}) (Expression, error) {
	// Constants

	// Error Handling and Recursion
	if len(terms) > 1 {
		// If there are more than one terms given to the multiply function, then
		// - Determine if any of the terms from 1 to end are a variable, monomial or polynomial
		followingIndex := FindPolynomialLikeObject(terms[1:])

		if followingIndex != -1 {
			// If we find a variable, monomial or polynomial,
			// then we must use recursion to compute multiplication.

			// First, check for any errors.
			if followingIndex != 0 { // If error exists, then check it and maybe throw something!
				if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
					return &Monomial{}, termAsError // throw error
				}
			}

			// Second, compute sub-product.
			m, err := m.Multiply(terms[0])
			if err != nil {
				return &Monomial{}, err
			}

			// Compute Product Among the
			return m.Multiply(terms[followingIndex:])

		} else {
			if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
				return &Monomial{}, termAsError // throw error
			}
		}
	}

	// Algorithm
	term1 := terms[0] // Collect Term 1
	switch term1.(type) {
	case float64:
		term1AsFloat, _ := term1.(float64)
		product := m.Copy()
		product.Coefficient *= term1AsFloat

		return product, nil
	case Variable:
		termAsV, _ := term1.(Variable)
		// Copy monomial into product
		product := m.Copy()

		// Try to find termAsV in input monomials variables
		tavFoundAt := -1
		for varIndex := 0; varIndex < len(product.Variables); varIndex++ {
			tempVar := product.Variables[varIndex]
			if tempVar == termAsV {
				tavFoundAt = varIndex
			}
		}

		// If found, then simply update the monomial exponent.
		if tavFoundAt != -1 {
			product.Exponents[tavFoundAt] += 1
		} else {
			product.Variables = append(product.Variables, termAsV)
			product.Exponents = append(product.Exponents, 1)
		}

		return product, nil

	case *Variable:
		product := m.Copy()
		term1AsVPointer, _ := term1.(*Variable)

		// Try to find termAsV in input monomials variables
		tavFoundAt := -1
		for varIndex := 0; varIndex < len(product.Variables); varIndex++ {
			tempVar := product.Variables[varIndex]
			if tempVar.IsEqualTo(*term1AsVPointer) {
				tavFoundAt = varIndex
			}
		}

		// If found, then simply update the monomial exponent.
		if tavFoundAt != -1 {
			product.Exponents[tavFoundAt] += 1
		} else {
			product.Variables = append(product.Variables, term1.(Variable))
			product.Exponents = append(product.Exponents, 1)
		}

		return product, nil

	case Monomial:
		termAsMonom, _ := term1.(Monomial)

		// Create Product
		productOut, err := m.Multiply(termAsMonom.Coefficient)
		if err != nil {
			return productOut, err
		}
		product, _ := productOut.(Monomial) // Convert to Monomial
		println(fmt.Sprintf("m = %v", m))
		println(fmt.Sprintf("product after first coefficient multiplication: %v", product))

		for varIndex := 0; varIndex < len(termAsMonom.Variables); varIndex++ {
			varInTAM := termAsMonom.Variables[varIndex]

			indexOfvIT := varInTAM.FoundIn(product.Variables)
			if indexOfvIT != -1 {
				product.Exponents[indexOfvIT] += termAsMonom.Exponents[varIndex]
			} else {
				product.Variables = append(product.Variables, varInTAM)
				product.Exponents = append(product.Exponents, termAsMonom.Exponents[varIndex])
			}

			println(fmt.Sprintf("m = %v", m))
			println(fmt.Sprintf("product after %vth loop coefficient multiplication: %v", varIndex, product))

		}
		return product, nil

	case *Monomial:
		termAsMonom, _ := term1.(*Monomial)

		// Create Product
		tempProduct := m.Copy()
		productPointer, err := tempProduct.Multiply(termAsMonom.Coefficient)
		product, _ := productPointer.(*Monomial)

		if err != nil {
			return product, err
		}
		for varIndex := 0; varIndex < len(termAsMonom.Variables); varIndex++ {
			varInTAM := termAsMonom.Variables[varIndex]

			indexOfvIT := varInTAM.FoundIn(product.Variables)
			if indexOfvIT != -1 {
				product.Exponents[indexOfvIT] += termAsMonom.Exponents[varIndex]
			} else {
				product.Variables = append(product.Variables, varInTAM)
				product.Exponents = append(product.Exponents, termAsMonom.Exponents[varIndex])
			}

		}
		return product, nil

	default:
		return &Monomial{}, fmt.Errorf("The input type %T was not expected!", term1)
	}
}

/*
FindMonomialInSlice
Description:

	Looks for a Monomial in a slice of objects that implement the empty interface (should be any object).
	Returns -1 if none are found.
*/
func FindMonomialInSlice(sliceIn []interface{}) int {
	// Constants

	// Algorithms
	for eltIndex := 0; eltIndex < len(sliceIn); eltIndex++ {
		// check to see if current element is a Variable
		tempElt := sliceIn[eltIndex]

		if _, ok := tempElt.(Monomial); ok {
			return eltIndex
		}

	}
	return -1
}

/*
Copy
Description:

	Doing this to avoid strange pointer and reference issues when creating new monomials.
*/
func (m Monomial) Copy() Monomial {
	// Constants

	// Algorithm
	monomOut := Monomial{}
	monomOut.Coefficient = m.Coefficient

	// Copy slice of Variable pointers
	monomOut.Variables = make([]Variable, len(m.Variables))
	copy(monomOut.Variables, m.Variables)

	// Copy slice of exponents
	monomOut.Exponents = make([]int, len(m.Exponents))
	copy(monomOut.Exponents, m.Exponents)

	return monomOut
}

/*
String
Description:

	Creates a string using the names of the variables in the monomial.
*/
func (m Monomial) String() string {
	// Constants
	numVars := len(m.Variables)

	// Create String
	stringOut := ""
	stringOut = fmt.Sprintf("%v", m.Coefficient)
	for varIndex := 0; varIndex < numVars; varIndex++ {
		switch {
		case m.Exponents[varIndex] == 1:
			stringOut = fmt.Sprintf("%v %v", stringOut, m.Variables[varIndex].String())
		case m.Exponents[varIndex] == 0:
		default:
			stringOut = fmt.Sprintf("%v (%v)^%v", stringOut, m.Variables[varIndex].String(), m.Exponents[varIndex])
		}
	}

	return stringOut

}
