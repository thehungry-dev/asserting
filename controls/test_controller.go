package controls

import "fmt"

// TestController is a telemetric data structure to record testing interaction
type TestController struct {
	helpers     uint
	hasError    bool
	errorFormat string
	errorArgs   []interface{}
}

// NewTestController initializes a telemetric test controller
func NewTestController() *TestController {
	return &TestController{}
}

// Helper records how many times this function has been invoked
func (t *TestController) Helper() {
	t.helpers++
}

// Errorf records the last error message
func (t *TestController) Errorf(format string, args ...interface{}) {
	t.errorFormat = format
	t.errorArgs = args
	t.hasError = true
}

// IsHelper is true when `Helper` has been called at least once
func (t *TestController) IsHelper() bool {
	return t.helpers > 0
}

// HasError is true when `Errorf` has been called at least once
func (t *TestController) HasError() bool {
	return t.hasError
}

// Error returns the last error recorded
func (t *TestController) Error() (string, []interface{}) {
	return t.errorFormat, t.errorArgs
}

// ErrorMessage returns the last error recorded as a string
func (t *TestController) ErrorMessage() string {
	args := t.errorArgs
	return fmt.Sprintf(t.errorFormat, args...)
}
