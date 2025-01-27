// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interfaces

import "testing"

func TestImplementationOf(t *testing.T) {
	var implA1Val = interfaceAImpl1{A: "A", B: 1}
	var implA2Val = interfaceAImpl2{A: "Z", B: 2}
	var implA1Ptr = &interfaceAImpl1{A: "PtrA", B: 1}
	var implA2Ptr = &interfaceAImpl2{A: "PtrZ", B: 2}
	var implB1Val = interfaceBImpl1{}
	var implB1Ptr = &interfaceBImpl1{}

	t.Run("impl1 as value", func(t *testing.T) {
		if ImplementationOf[testInterfaceA](implA1Val) {
			t.Logf("implA1Val implements testInterfaceA")
		} else {
			t.Errorf("implA1Val does not implement testInterfaceA")
		}
	})

	t.Run("impl1 as pointer", func(t *testing.T) {
		if !ImplementationOf[testInterfaceA](implA1Ptr) {
			t.Errorf("implA1Ptr does not implement testInterfaceA")
		} else {
			t.Logf("implA1Ptr implements testInterfaceA")
		}
	})

	t.Run("impl2 as value", func(t *testing.T) {
		if !ImplementationOf[testInterfaceA](implA2Val) {
			t.Errorf("implA2Val does not implement testInterfaceA")
		} else {
			t.Logf("implA2Val implements testInterfaceA")
		}
	})

	t.Run("impl2 as pointer", func(t *testing.T) {
		if !ImplementationOf[testInterfaceA](implA2Ptr) {
			t.Errorf("implA2Ptr does not implement testInterfaceA")
		} else {
			t.Logf("implA2Ptr implements testInterfaceA")
		}
	})

	t.Run("implB1 as value", func(t *testing.T) {
		if ImplementationOf[testInterfaceA](implB1Val) {
			t.Errorf("implB1Val does not implement testInterfaceB")
		} else {
			t.Logf("implB1Val implements testInterfaceB")
		}
	})

	t.Run("implB1 as pointer", func(t *testing.T) {
		if ImplementationOf[testInterfaceA](implB1Ptr) {
			t.Errorf("implB1Ptr does not implement testInterfaceB")
		} else {
			t.Logf("implB1Ptr implements testInterfaceB")
		}
	})
}

func TestImplements_EdgeCases(t *testing.T) {
	var nilInterface testInterfaceA

	t.Run("nil interface", func(t *testing.T) {
		if ImplementationOf[testInterfaceA](nilInterface) {
			t.Errorf("nilInterface implements testInterfaceA")
		} else {
			t.Logf("nilInterface does not implements testInterfaceA")
		}
	})

	var nilInterfacePtr *testInterfaceA

	t.Run("nil interface pointer", func(t *testing.T) {
		if ImplementationOf[testInterfaceA](nilInterfacePtr) {
			t.Errorf("nilInterfacePtr implements testInterfaceA")
		} else {
			t.Logf("nilInterfacePtr does not implements testInterfaceA")
		}
	})
}
