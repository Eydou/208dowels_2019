//
// EPITECH PROJECT, 2020
// 208dowels_2019
// File description:
// libBigFloat
//

package functions

import "math/big"

//Zero to zero
func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

//Mul Multiply float
func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

//Add float
func Add(a, b *big.Float) *big.Float {
	return Zero().Add(a, b)
}

//Sub sub float
func Sub(a, b *big.Float) *big.Float {
	return Zero().Sub(a, b)
}

//Pow power float
func Pow(a *big.Float, e uint64) *big.Float {
	if e == 0 {
		one := big.NewFloat(float64(1))
		return one
	}
	result := Zero().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}
