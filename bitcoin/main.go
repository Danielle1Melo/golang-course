package main

import (
	ecc "elliptic_curve"
	"fmt"
	"math/rand"
)

func SolveField19MultiplieSet() {
	//randomly select a num from 1 to 18
	min := 1
	max := 18
	k := rand.Intn(max-min) + min
	fmt.Println("ramdomly select k is %d\n", k)
	element := ecc.NewFieldElement(19, uint64(k))
	for i := 0; i < 19; i++ {
		fmt.Printf("element %d multiplie with %d is %v\n", k, i, element.ScalarMul(uint64(i)))
	}
}

func main() {
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	res := f44.Add(f33)
	fmt.Printf("field element 44 add to field element 33 is %v\n", res)
	// -44negate of 44 is 57 - 44 is 13
	fmt.Printf("negate of field element 44 is %v\n", res.Negate())
	fmt.Printf("Field element 44 - 33 is %v\n", f44.Substract(f33))
	fmt.Printf("Field element 33 - 44 is %v\n", f33.Substract(f44))

	//check (46 + 44) % order(57) == 33
	fmt.Printf("check 46 + 44 over module of 57 is %d\n", (46+44)%57)
	f46 := ecc.NewFieldElement(57, 46)
	fmt.Printf("field element 46 + 44 is %v\n", f46.Add(f44))

	fmt.Printf("multiplie 46 with itself is: %v\n", f46.Mutiplie(f46))
	fmt.Printf("element 46 with the power of 2 is %v\n", f46.Power(2))

	SolveField19MultiplieSet()
}
