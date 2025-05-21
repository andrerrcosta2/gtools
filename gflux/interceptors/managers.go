// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interceptors

import (
	"reflect"
	"sync"
)

type ObjectManager struct {
	objects sync.Map // Stores registered objects
	hooks   sync.Map // Stores hooks per method
}

// Hook defines a function that intercepts method calls
// It receives the method name, arguments, and original function
// The function can modify the arguments or the return value
type Hook func(method string, args []reflect.Value, original func([]reflect.Value) []reflect.Value) []reflect.Value

// NewObjectManager creates an instance of ObjectManager
func NewObjectManager() *ObjectManager {
	return &ObjectManager{}
}

// Register stores an object and allows method interception
func (m *ObjectManager) Register(name string, obj any) {
	m.objects.Store(name, obj)
}

// AddHook adds a method-level hook for an object
func (m *ObjectManager) AddHook(objectName, methodName string, hook Hook) {
	key := objectName + "." + methodName
	m.hooks.Store(key, hook)
}

// Call invokes a method on a registered object, applying any hooks
func (m *ObjectManager) Call(objectName, methodName string, args ...any) any {
	obj, ok := m.objects.Load(objectName)
	if !ok {
		panic("object not found")
	}

	objValue := reflect.ValueOf(obj)
	method := objValue.MethodByName(methodName)
	if !method.IsValid() {
		panic("method not found")
	}

	inputArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		inputArgs[i] = reflect.ValueOf(arg)
	}

	key := objectName + "." + methodName
	if hook, exists := m.hooks.Load(key); exists {
		return hook.(Hook)(methodName, inputArgs, method.Call)
	}

	results := method.Call(inputArgs)
	if len(results) > 0 {
		return results[0].Interface()
	}
	return nil
}
