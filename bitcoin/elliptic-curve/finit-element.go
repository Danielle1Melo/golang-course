package elliptic_curve

import (
	"fmt"
)

type FieldElement struct {
	order uint64 //fiel order
	num   uint64 //value of the given element in the field
}

func NewFieldElement(order uint64, num uint64) *FieldElement {
	//init function for FieldElement

	if num >= order {
		err := fmt.Sprintf("Num not in the range of 0 to %d", order-1)
		panic(err)
	}

	return &FieldElement{
		order: order,
		num:   num,
	}
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement {order: %d, num %d}\n", f.order, f.num)
}

func (f *FieldElement) EquatTo(other *FieldElement) bool {
	return f.order == other.order && f.num == other.num
}

func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	if f.order == other.order {
		panic("add need to do on the field element with the same order")
	}

	return NewFieldElement(f.order, (f.num+other.num)%f.order)
}

// a ,b (a +b) % order = 0, b is called negate of a, b = -a
func (f *FieldElement) Negate() *FieldElement {
	return NewFieldElement(f.order, f.order-f.num)
}
