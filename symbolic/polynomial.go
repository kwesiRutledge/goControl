package symbolic

import "fmt"

/*
polynomial.go
Description:
	Defines the Polynomial object and attempts to define a new interface for it.
*/

type Polynomial struct {
	Monomials []Monomial
}

// Functions
// =========

/*
FindPolynomialInSlice
Description:

	Looks for a Polynomial in a slice of objects that implement the empty interface (should be any object).
	Returns -1 if none are found.
*/
func FindPolynomialInSlice(sliceIn []interface{}) int {
	// Constants

	// Algorithms
	for eltIndex := 0; eltIndex < len(sliceIn); eltIndex++ {
		// check to see if current element is a Variable
		tempElt := sliceIn[eltIndex]

		if _, ok := tempElt.(Polynomial); ok {
			return eltIndex
		}

	}
	return -1
}

/*
FindPolynomialLikeObject
Description:

	This function attempts to find any objects of type:
	- Variable
	- Monomial, or
	- Polynomial
	in an input slice.
	It makes use of the Find methods for each variable's type.
*/
func FindPolynomialLikeObject(sliceIn []interface{}) int {
	// Constants

	// Algorithm
	return FindExpressionObject(sliceIn)

}

/*
Multiply
Description:

	Multiplies the input polynomial by an unknown number of other Expression objects.
*/
func (p Polynomial) Multiply(terms ...interface{}) (Expression, error) {
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
					return &Polynomial{}, termAsError // throw error
				}
			}

			// Second, compute sub-product.
			pNew, err := p.Multiply(terms[0])
			if err != nil {
				return &Polynomial{}, err
			}

			// Compute Product Among the
			return pNew.Multiply(terms[followingIndex:])

		} else {
			if termAsError, ok := terms[1].(error); ok && (termAsError != nil) { // Converted extra term to an error
				return &Polynomial{}, termAsError // throw error
			}
		}
	}

	// Algorithm
	term1 := terms[0] // Collect Term 1
	switch term1.(type) {
	case float64:
		product := p
		for i, monomial := range p.Monomials {
			multMonomial, err := monomial.Multiply(term1.(float64))
			if err != nil {
				return product, err
			}
			product.Monomials[i] = multMonomial.(Monomial)
		}
		return product, nil
	case Variable:
		termAsV, _ := term1.(Variable)
		product := p

		// Try to find termAsV in input monomials variables
		for i, monomial := range p.Monomials {
			multMonomial, err := monomial.Multiply(termAsV)
			if err != nil {
				return product, err
			}
			product.Monomials[i] = multMonomial.(Monomial)
		}
		return product, nil

	//case *Variable:
	//	product := Monomial{Coefficient: m.Coefficient, Variables: m.Variables, Exponents: m.Exponents}
	//
	//	// Try to find termAsV in input monomials variables
	//	tavFoundAt := -1
	//	for varIndex := 0; varIndex < len(product.Variables); varIndex++ {
	//		tempVar := product.Variables[varIndex]
	//		if tempVar.IsEqualTo(term1.(*Variable)) {
	//			tavFoundAt = varIndex
	//		}
	//	}
	//
	//	// If found, then simply update the monomial exponent.
	//	if tavFoundAt != -1 {
	//		product.Exponents[tavFoundAt] += 1
	//	} else {
	//		product.Variables = append(product.Variables, term1.(*Variable))
	//		product.Exponents = append(product.Exponents, 1)
	//	}
	//
	//	return &product, nil
	//
	//case Monomial:
	//	termAsMonom, _ := term1.(Monomial)
	//
	//	// Create Product
	//	product := m.Copy()
	//	fmt.Println(termAsMonom)
	//	productPointer, err := product.Multiply(termAsMonom.Coefficient)
	//	product2, _ := productPointer.(*Monomial)
	//
	//	if err != nil {
	//		return product2, err
	//	}
	//	for varIndex := 0; varIndex < len(termAsMonom.Variables); varIndex++ {
	//		varInTAM := termAsMonom.Variables[varIndex]
	//
	//		indexOfvIT := varInTAM.FoundIn(product2.Variables)
	//		if indexOfvIT != -1 {
	//			product2.Exponents[indexOfvIT] += termAsMonom.Exponents[varIndex]
	//		} else {
	//			product2.Variables = append(product2.Variables, varInTAM)
	//			product2.Exponents = append(product2.Exponents, termAsMonom.Exponents[varIndex])
	//		}
	//
	//		fmt.Println(termAsMonom)
	//
	//	}
	//	return product2, nil
	//
	//case *Monomial:
	//	termAsMonom, _ := term1.(*Monomial)
	//
	//	// Create Product
	//	tempProduct := m.Copy()
	//	productPointer, err := tempProduct.Multiply(termAsMonom.Coefficient)
	//	product, _ := productPointer.(*Monomial)
	//
	//	if err != nil {
	//		return product, err
	//	}
	//	for varIndex := 0; varIndex < len(termAsMonom.Variables); varIndex++ {
	//		varInTAM := termAsMonom.Variables[varIndex]
	//
	//		indexOfvIT := varInTAM.FoundIn(product.Variables)
	//		if indexOfvIT != -1 {
	//			product.Exponents[indexOfvIT] += termAsMonom.Exponents[varIndex]
	//		} else {
	//			product.Variables = append(product.Variables, varInTAM)
	//			product.Exponents = append(product.Exponents, termAsMonom.Exponents[varIndex])
	//		}
	//
	//	}
	//	return product, nil

	default:
		return &Monomial{}, fmt.Errorf("The input type %T was not expected!", term1)
	}
}
