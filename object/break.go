package object

// BreakValue is an implementation detail used to handle break statements
type BreakValue struct{}

func (rv *BreakValue) Type() Type {
	return BREAK_VALUE_OBJ
}

func (rv *BreakValue) Inspect() string {
	return "BREAK"
}

func (rv *BreakValue) InvokeMethod(method string, args ...Object) Object {
	return nil
}

func (rv *BreakValue) ToInterface() interface{} {
	return "<BREAK_VALUE>"
}
