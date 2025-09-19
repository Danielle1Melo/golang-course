package main

import (
	ecc "elliptic_curve"
	"fmt"
)

func main() {
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	res := f44.Add(f33)
	fmt.Printf("field element 44 add to field element 33 is %v\n", res)
	// -44negate of 44 is 57 - 44 is 13
	fmt.Printf("negate of field element 44 is %vv\n", res.Negate())
}
