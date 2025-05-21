// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package numbers

import "math/big"

// BinCoef calculates C(n, k)
func BinCoef(n, k int) *big.Int {
	if k > n {
		return big.NewInt(0)
	}
	factN := big.NewInt(1)
	factK := big.NewInt(1)
	factNMinusK := big.NewInt(1)

	for i := 1; i <= n; i++ {
		factN.Mul(factN, big.NewInt(int64(i)))
		if i <= k {
			factK.Mul(factK, big.NewInt(int64(i)))
		}
		if i <= (n - k) {
			factNMinusK.Mul(factNMinusK, big.NewInt(int64(i)))
		}
	}

	result := new(big.Int)
	result.Div(factN, factK.Mul(factK, factNMinusK))
	return result
}

func ExcBinCoef(n, k int) *big.Int {
	if k == 0 || k == n {
		return big.NewInt(0)
	}
	return BinCoef(n, k).Set(BinCoef(n-1, k-1))
}

func ExcAprc(subsets, setSize, subsetSize int) int {
	return subsets * subsetSize / setSize
}

func Sum(values ...int) int { //
	var total int
	for _, value := range values {
		total += value
	}
	return total
}
