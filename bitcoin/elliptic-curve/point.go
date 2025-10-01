package elliptic_curve

import "math/big"

type OP_TYPE int

const (
	ADD OP_TYPE = iota
	SUB
	MUL
	DIV
	EXP
)

type Point struct {
	//coeficients of the curve
	a *big.Int
	b *big.Int
	//x, y should be the point on the curve
	x *big.Int
	y *big.Int
}

func opOnBig(x *big.Int, y *big.Int, optype OP_TYPE) *big.Int {
	var op big.Int
	switch optype {
	case ADD:
		return op.Add(x, y)
	case SUB:
		return op.Sub(x, y)
	case MUL:
		return op.Mul(x, y)
	case DIV:
		return op.Div(x, y)
	case EXP:
		return op.Exp(x, y, nil)
	}

	panic("should nor come to here")
}
