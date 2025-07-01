// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package assertlite

import "fmt"

// mockTesting is a simple mock for testing purposes
type mockTesting struct {
	errors   []string
	failures int
	fatal    []string
}

func (m *mockTesting) Helper() {}

func (m *mockTesting) Errorf(format string, args ...interface{}) {
	m.errors = append(m.errors, fmt.Sprintf(format, args...))
}

func (m *mockTesting) Error(args ...interface{}) {
	m.errors = append(m.errors, fmt.Sprint(args...))
}

func (m *mockTesting) FailNow() {
	m.failures++
}

func (m *mockTesting) Fatal(args ...any) {
	m.failures++
	m.fatal = append(m.fatal, fmt.Sprint(args...))
}

func (m *mockTesting) Fatalf(format string, args ...any) {
	m.failures++
	m.fatal = append(m.fatal, fmt.Sprintf(format, args...))
}

func (m *mockTesting) Failed() bool {
	return m.failures > 0
}

func (m *mockTesting) HasErrors() bool {
	return len(m.errors) > 0
}

func (m *mockTesting) Clear() {
	m.errors = make([]string, 0)
	m.fatal = make([]string, 0)
	m.failures = 0
}
