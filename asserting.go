// Package asserting provides macros to verify output, well-suited for use in `go test`
package asserting

import "github.com/thehungry-dev/asserting/rejection"

// TestController provides functionality to interrupt and log during a test execution. It's usually `testing.T`
type TestController interface {
	Helper()
	Errorf(format string, args ...interface{})
}

// Assert tests if the result is true
var Assert Assertion = Assertion(assertf)

// Assertion provides a function to assert results and recover from functions triggering panic
type Assertion func(TestController, bool, ...interface{})

// PanicMsg asserts that the provided function triggers panic with the provided message
func (assert Assertion) PanicMsg(t TestController, do func(), assertMsg func(interface{}) bool) {
	t.Helper()

	panicked := true

	defer func() {
		t.Helper()

		err := recover()

		assert(t, panicked, rejection.PanicExpected)

		if panicked {
			result := assertMsg(err)
			assert(t, result, rejection.InvalidPanicMsg)
		}
	}()
	do()

	panicked = false
}

// Panic asserts that the provided function triggers panic
func (assert Assertion) Panic(t TestController, do func()) {
	t.Helper()

	assert.PanicMsg(t, do, any)
}

func any(_ interface{}) bool { return true }

func assertf(t TestController, result bool, msgArgs ...interface{}) {
	t.Helper()

	msg := rejection.FalseAssertion
	var args []interface{}

	switch len(msgArgs) {
	case 0:
	case 1:
		msg = msgArgs[0].(string)
	default:
		msg = msgArgs[0].(string)
		args = msgArgs[1:]
	}

	if !result {
		t.Errorf(msg, args...)
	}
}
