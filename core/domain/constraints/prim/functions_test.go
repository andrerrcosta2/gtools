// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prim

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestToBytes(t *testing.T) {
	valueA := 12345
	valueB := -12345

	iA := make([]byte, 8) // size of int64 in bytes
	binary.BigEndian.PutUint64(iA, uint64(valueA))

	iB := make([]byte, 8) // size of int64 in bytes
	binary.BigEndian.PutUint64(iB, uint64(valueB))

	sA := []byte(fmt.Sprintf("%d", valueA))
	sB := []byte(fmt.Sprintf("%d", valueB))

	fmt.Printf("A: %v\nB: %v\n", iA, iB)
	fmt.Printf("iA == iB? %t\n", bytes.Equal(iA, iB))

	fmt.Printf("sA == sB? %t\n", bytes.Equal(sA, sB))

}
