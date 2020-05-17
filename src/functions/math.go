//
// EPITECH PROJECT, 2020
// 208dowels_2019
// File description:
// error
//

package functions

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type array struct {
	Ox     []int
	xBis   []int
	xTotal []int
	TxBis  []float64
	Tx     []float64
	p      float64
	square float64
}

func (st *array) fillSquare(table []float64, x2 [][]float64) []float64 {
	j := len(st.Ox) - 2
	for i := 0; i < len(x2[j]); i++ {
		if x2[j][i] < st.square && (i < 12) {
			table = append(table, x2[0][i+1])
		}
	}
	return table
}

func (st *array) fitValid(x2 [][]float64) {
	table := make([]float64, 0)
	j := len(st.Ox) - 2

	table = st.fillSquare(table, x2)
	if len(table) == 0 {
		fmt.Printf("Fit validity:\t\tP > 99%%\n")
	} else if len(table) > 0 && (st.square > x2[j][len(table)]) {
		fmt.Printf("Fit validity:\t\tP < 1%%\n")
	} else {
		fmt.Printf("Fit validity:\t\t%.0f%% < P < %.0f%%\n", table[1], table[0])
	}
}

func (st *array) chiSquare() {
	for i := 0; i < len(st.Ox); i++ {
		st.square += math.Pow(float64(st.Ox[i])-st.Tx[i], 2) / st.Tx[i]
	}
	fmt.Printf("Chi-squared:\t\t%.3f\n", st.square)
}

func (st *array) FreedomDegrees() {
	fmt.Printf("Degrees of freedom:\t%d\n", (len(st.Ox) - 2))
}

func (st *array) PrintP() {
	fmt.Printf("Distribution:\t\tB(100, %.4f)\n", st.p)
}

func (st *array) DistributionCalc(numbers []int) {
	st.p = 0.0

	for i := 0; i < 9; i++ {
		st.p += float64(i) * float64(numbers[i]) * 1.0
	}
	st.p = st.p / 10000 * 1.0
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func (st *array) SetupCalcPrintTx() {
	k := 0
	sums := float64(0)

	for i := 0; k < len(st.xTotal); i++ {
		sum := float64(0)
		for j := 0; j != st.xBis[i]+1; j++ {
			sums += st.TxBis[k]
			sum += st.TxBis[k]
			k++
		}
		if k < len(st.xTotal) {
			st.Tx = append(st.Tx, sum)
			fmt.Printf(" %.1f\t|", sum)
		} else {
			st.Tx = append(st.Tx, sum+100-sums)
			fmt.Printf(" %.1f\t|", sum+100-sums)
		}
	}
	fmt.Printf(" 100\n")
}

//100. * BinomialCoeff * (math.Pow((1 - st.p), (100. - float64(st.xTotal[i]))))
func (st *array) CalcPrintTx() {
	var coeff = new(big.Int)

	fmt.Printf("  Tx\t|")
	for i := 0; i != len(st.xTotal); i++ {
		x := big.NewInt(int64(st.xTotal[i]))
		n := big.NewInt(int64(100))
		result := coeff.Binomial(n.Int64(), x.Int64())
		p := big.NewFloat(float64(st.p))
		one := big.NewFloat(float64(1))
		first := Pow(p, uint64(st.xTotal[i]))
		second := Pow(Sub(one, p), 100-uint64(st.xTotal[i]))
		third := Mul(first, second)
		zero := new(big.Float).SetInt(result)
		hundred := new(big.Float).SetInt(n)
		fourth := Mul(zero, third)
		final := Mul(hundred, fourth)
		bigstr := final.String()
		tx, _ := strconv.ParseFloat(bigstr, 64)
		st.TxBis = append(st.TxBis, tx)
	}
	st.SetupCalcPrintTx()
}

func (st *array) PrintOx() {
	fmt.Printf("  Ox\t|")
	for i := 0; i != len(st.Ox); i++ {
		fmt.Printf(" %d\t|", st.Ox[i])
	}
	fmt.Printf(" 100\n")
}

func (st *array) PrintX() {
	i, sum := 0, 0

	for k := 0; k != len(st.xBis)-1; k++ {
		sum += st.xBis[k] + 1
	}
	fmt.Printf("   x\t|")
	for j := 0; j != len(st.xBis)-1; j++ {
		if st.xBis[j] > 0 {
			fmt.Printf(" %d-", i)
			fmt.Printf("%d\t|", i+st.xBis[j])
			i += st.xBis[j] + 1
		} else {
			fmt.Printf(" %d\t|", i)
			i++
		}
	}
	fmt.Printf(" %d+\t| Total\n", sum)
}

//ArrayClass x
func (st *array) ArrayClass(numbers []int) {
	ox, number := 0, 0

	for i := len(numbers) - 1; i >= 0; i-- {
		if (i - 1) > -1 {
			if numbers[i] < 10 {
				check := 0
				for j := i; j >= 0; j-- {
					if number < 10 {
						number += numbers[j]
						check++
					} else if number >= 10 {
						break
					}
				}
				ox = check - 1
				i -= check - 1
			} else if numbers[i] >= 10 {
				number = numbers[i]
				nbBis, check := 0, 0
				for j := i - 1; j >= 0; j-- {
					if nbBis < 10 {
						nbBis += numbers[j]
						check++
					} else if nbBis >= 10 {
						break
					}
				}
				enter := false
				if check == 2 && nbBis >= 10 {
					enter = true
				}
				if check != 1 && enter == false {
					for j := check; j >= 1; j-- {
						number += numbers[i-j]
					}
					ox = check
					i -= check
				}
			}
		}
		st.Ox = append(st.Ox, number)
		st.xBis = append(st.xBis, ox)
		ox, number = 0, 0
	}
	st.Ox = reverseInts(st.Ox)
	st.xBis = reverseInts(st.xBis)
}

//MathParse args
func MathParse(numbers []int) int {
	st := array{}
	x2 := [][]float64{
		{99., 90., 80., 70., 60., 50., 40., 30., 20., 10., 5., 2., 1.},
		{0.00, 0.02, 0.06, 0.15, 0.27, 0.45, 0.71, 1.07, 1.64, 2.71, 3.84, 5.41, 6.63},
		{0.02, 0.21, 0.45, 0.71, 1.02, 1.39, 1.83, 2.41, 3.22, 4.61, 5.99, 7.82, 9.21},
		{0.11, 0.58, 1.01, 1.42, 1.87, 2.37, 2.95, 3.66, 4.64, 6.25, 7.81, 9.84, 11.34},
		{0.30, 1.06, 1.65, 2.19, 2.75, 3.36, 4.04, 4.88, 5.99, 7.78, 9.49, 11.67, 13.28},
		{0.55, 1.61, 2.34, 3.00, 3.66, 4.35, 5.13, 6.06, 7.29, 9.24, 11.07, 13.39, 15.09},
		{0.87, 2.20, 3.07, 3.83, 4.57, 5.35, 6.21, 7.23, 8.56, 10.64, 12.59, 15.03, 16.81},
		{1.24, 2.83, 3.82, 4.67, 5.49, 6.35, 7.28, 8.38, 9.80, 12.02, 14.07, 16.62, 18.48},
		{1.65, 3.49, 4.59, 5.53, 6.42, 7.34, 8.35, 9.52, 11.03, 13.36, 15.51, 18.17, 20.09},
		{2.09, 4.17, 5.38, 6.39, 7.36, 8.34, 9.41, 10.66, 12.24, 14.68, 16.92, 19.68, 21.67},
		{2.56, 4.87, 6.18, 7.27, 8.30, 9.34, 10.47, 11.78, 13.44, 15.99, 18.31, 21.16, 23.21}}

	for i := 0; i != len(numbers); i++ {
		st.xTotal = append(st.xTotal, i)
	}
	st.ArrayClass(numbers)
	st.PrintX()
	st.PrintOx()
	st.DistributionCalc(numbers)
	st.CalcPrintTx()
	st.PrintP()
	st.chiSquare()
	st.FreedomDegrees()
	st.fitValid(x2)
	return 0
}
