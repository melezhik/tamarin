package object

type DatabaseConnection struct {
	Conn interface{}
}

func (c *DatabaseConnection) Type() Type {
	return DB_CONNECTION_OBJ
}

func (c *DatabaseConnection) Inspect() string {
	return "<DB_CONNECTION>"
}

func (c *DatabaseConnection) InvokeMethod(method string, args ...Object) Object {
	return nil
}

func (c *DatabaseConnection) ToInterface() interface{} {
	return nil
}
