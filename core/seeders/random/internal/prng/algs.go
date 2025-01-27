// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prng

func Xorshift128Plus(seed uint64) func() uint64 {
	state0 := seed
	state1 := seed ^ 0xdeadbeefcafebabe

	return func() uint64 {
		s0 := state0
		s1 := state1
		state0 = s1
		s1 ^= s1 << 23
		state1 = s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26)
		return state1 + s0
	}
}
