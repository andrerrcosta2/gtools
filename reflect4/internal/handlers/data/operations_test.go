// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"reflect"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"github.com/andrerrcosta2/gtools/reflect4/internal/equals"
	"github.com/andrerrcosta2/gtools/reflect4/internal/sprint"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
)

// TestDeepCopyArray_Default should have equal copy based on the constraints:
//   - functions aren't copied
//   - unsafe pointers
func TestDeepCopyArray_Default(t *testing.T) {
	s := DefaultCopyStrat()
	gtests.Arrays.Fuzz().Values().All().Each(func(arr any) {
		a := reflect.ValueOf(arr)
		b, err := deepCopyArray(a, s)
		assertlite.NoError(t, err)
		if !equals.Deep(a, b, read.SkipFunctions) {
			diff, _ := differ.Between(a, b, read.SkipFunctions)
			t.Errorf("test failed for '%T':\n%s", arr, diff)
			sa, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
			assertlite.NoError(t, err)
			sb, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
			t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
		}
	})
}

func TestDeepCopyMap_Default(t *testing.T) {
	s := DefaultCopyStrat()
	gtests.Maps.Fuzz().Values().All().Each(func(arr any) {
		a := reflect.ValueOf(arr)
		b, err := deepCopyMap(a, s)
		assertlite.NoError(t, err)
		if !equals.Deep(a, b, read.SkipFunctions) {
			diff, _ := differ.Between(a, b, read.SkipFunctions)
			t.Errorf("test failed for '%T':\n%s", arr, diff)
			sa, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
			assertlite.NoError(t, err)
			sb, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
			t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
		}
	})
}

func TestDeepCopySlice_Default(t *testing.T) {
	s := DefaultCopyStrat()

	gtests.Slices.Fuzz().Values().All().Each(func(slice any) {
		a := reflect.ValueOf(slice)
		b, err := deepCopySlice(a, s)
		assertlite.NoError(t, err)
		if !equals.Deep(a, b, read.SkipFunctions) {
			diff, _ := differ.Between(a, b, read.SkipFunctions)
			t.Errorf("test failed for '%T':\n%s", slice, diff)
			sa, err := sprint.Slice(a, read.SkipFunctions, read.SkipUnsafePtr)
			assertlite.NoError(t, err)
			sb, err := sprint.Slice(a, read.SkipFunctions, read.SkipUnsafePtr)
			t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
		}
	})
}

func TestDeepCopyStruct_Default(t *testing.T) {
	s := DefaultCopyStrat()
	gtests.Structs.Fuzz().Categories().Values().All().Each(func(stc any) {
		a := reflect.ValueOf(stc)
		b, err := deepCopyStruct(a, s)
		assertlite.NoError(t, err)
		if !equals.Deep(a, b, read.SkipFunctions) {
			diff, _ := differ.Between(a, b, read.SkipFunctions)
			t.Errorf("test failed for '%T':\n%s", stc, diff)
			sa, err := sprint.Struct(a, read.SkipFunctions, read.SkipUnsafePtr)
			assertlite.NoError(t, err)
			sb, err := sprint.Struct(a, read.SkipFunctions, read.SkipUnsafePtr)
			t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
		}
	})
}

func TestDeepCopyChan_Default(t *testing.T) {
	s := DefaultCopyStrat()
	gtests.Chans.Fuzz().Values().All().Each(func(stc any) {
		a := reflect.ValueOf(stc)
		b, err := defaultDeepCopyChan(a, s)
		assertlite.NoError(t, err)
		if !equals.Deep(a, b, read.SkipFunctions) {
			diff, _ := differ.Between(a, b, read.SkipFunctions)
			t.Errorf("test failed for '%T':\n%s", stc, diff)
			sa, err := sprint.Chan(a)
			assertlite.NoError(t, err)
			sb, err := sprint.Chan(a)
			t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
		}
	})
}

func TestDeepCopyFunc_Default(t *testing.T) {
	s := DefaultCopyStrat()
	gtests.Funcs.Consumers().Values().All().Each(func(stc any) {
		a := reflect.ValueOf(stc)
		b, err := defaultDeepCopyFunc(a, s)
		assertlite.NoError(t, err)
		if !equals.Deep(a, b, read.SkipFunctions) {
			diff, _ := differ.Between(a, b, read.SkipFunctions)
			t.Errorf("test failed for '%T':\n%s", stc, diff)
			sa, err := sprint.Func(a)
			assertlite.NoError(t, err)
			sb, err := sprint.Func(a)
			t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
		}
	})
}

func TestDeepCopyPointer_Default(t *testing.T) {
	s := DefaultCopyStrat()
	t.Run("array references", func(t *testing.T) {
		gtests.Arrays.Fuzz().Refs().All().Each(func(arr any) {
			a := reflect.ValueOf(arr)
			b, err := defaultDeepCopyPointer(a, s)
			assertlite.NoError(t, err)
			if !equals.Deep(a, b, read.SkipFunctions) {
				diff, _ := differ.Between(a, b, read.SkipFunctions)
				t.Errorf("test failed for '%T':\n%s", arr, diff)
				sa, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
				assertlite.NoError(t, err)
				sb, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
				t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
			}
		})
	})

	t.Run("map references", func(t *testing.T) {
		gtests.Maps.Fuzz().Refs().All().Each(func(arr any) {
			a := reflect.ValueOf(arr)
			b, err := defaultDeepCopyPointer(a, s)
			assertlite.NoError(t, err)
			if !equals.Deep(a, b, read.SkipFunctions) {
				diff, _ := differ.Between(a, b, read.SkipFunctions)
				t.Errorf("test failed for '%T':\n%s", arr, diff)
				sa, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
				assertlite.NoError(t, err)
				sb, err := sprint.Array(a, read.SkipFunctions, read.SkipUnsafePtr)
				t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
			}
		})
	})

	t.Run("slice references", func(t *testing.T) {
		gtests.Slices.Fuzz().Refs().All().Each(func(slice any) {
			a := reflect.ValueOf(slice)
			b, err := defaultDeepCopyPointer(a, s)
			assertlite.NoError(t, err)
			if !equals.Deep(a, b, read.SkipFunctions) {
				diff, _ := differ.Between(a, b, read.SkipFunctions)
				t.Errorf("test failed for '%T':\n%s", slice, diff)
				sa, err := sprint.Slice(a, read.SkipFunctions, read.SkipUnsafePtr)
				assertlite.NoError(t, err)
				sb, err := sprint.Slice(a, read.SkipFunctions, read.SkipUnsafePtr)
				t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
			}
		})
	})

	t.Run("struct references", func(t *testing.T) {
		gtests.Structs.Fuzz().Categories().Refs().All().Each(func(stc any) {
			a := reflect.ValueOf(stc)
			b, err := defaultDeepCopyPointer(a, s)
			assertlite.NoError(t, err)
			if !equals.Deep(a, b, read.SkipFunctions) {
				diff, _ := differ.Between(a, b, read.SkipFunctions)
				t.Errorf("test failed for '%T':\n%s", stc, diff)
				sa, err := sprint.Struct(a, read.SkipFunctions, read.SkipUnsafePtr)
				assertlite.NoError(t, err)
				sb, err := sprint.Struct(a, read.SkipFunctions, read.SkipUnsafePtr)
				t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
			}
		})
	})

	t.Run("chan references", func(t *testing.T) {
		gtests.Chans.Fuzz().Refs().All().Each(func(stc any) {
			a := reflect.ValueOf(stc)
			b, err := defaultDeepCopyPointer(a, s)
			assertlite.NoError(t, err)
			if !equals.Deep(a, b, read.SkipFunctions) {
				diff, _ := differ.Between(a, b, read.SkipFunctions)
				t.Errorf("test failed for '%T':\n%s", stc, diff)
				sa, err := sprint.Chan(a)
				assertlite.NoError(t, err)
				sb, err := sprint.Chan(a)
				t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
			}
		})
	})

	t.Run("func references", func(t *testing.T) {
		gtests.Funcs.Consumers().Refs().All().Each(func(stc any) {
			a := reflect.ValueOf(stc)
			b, err := defaultDeepCopyPointer(a, s)
			assertlite.NoError(t, err)
			if !equals.Deep(a, b, read.SkipFunctions) {
				diff, _ := differ.Between(a, b, read.SkipFunctions)
				t.Errorf("test failed for '%T':\n%s", stc, diff)
				sa, err := sprint.Func(a)
				assertlite.NoError(t, err)
				sb, err := sprint.Func(a)
				t.Logf("[A]: \n\n%s\n[B]: \n\n%s\n", sa, sb)
			}
		})
	})
}
