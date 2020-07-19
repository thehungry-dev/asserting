# Asserting

Output verification library for asserting in `go test`

## Usage

```go
package my_test

import (
	"testing"

	. "github.com/Fire-Dragon-DoL/lab/asserting"
)

func TestAssert(t *testing.T) {
	Assert(t, 1 == 1) // Test is successful
}

func TestAssertPanic(t *testing.T) {
	Assert.Panic(t, func() {
    panic("A panic")
  }) // Test is successful
}

func TestAssertPanicMsg(t *testing.T) {
  panicFn := func() { panic("A panic") }
	Assert.Panic(t, panicFn, func(msg interface{}) bool {
    prose, ok := msg.(string)
    return ok && prose == "A panic"
  }) // Test is successful
}

func TestAssertDifferentPanicMsg(t *testing.T) {
  panicFn := func() { panic("A panic") }
	Assert.Panic(t, panicFn, func(msg interface{}) bool {
    prose, ok := msg.(string)
    return ok && prose == "Other panic"
  }) // Test fails
}
```
