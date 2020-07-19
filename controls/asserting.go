package controls

// PanicMsg is the panic message provided by `WillPanic`
const PanicMsg string = "Will panic"

// NotPanic is a no-op
func NotPanic() {}

// WillPanic panicks with `PanicMsg`
func WillPanic() { panic(PanicMsg) }

// MatchesMsg matches any panic message
func MatchesMsg(interface{}) bool { return true }

// RejectMsg doesn't match any panic message
func RejectMsg(interface{}) bool { return false }
