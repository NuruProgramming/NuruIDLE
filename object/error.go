package object

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return e.Message
}
func (e *Error) Type() ObjectType { return ERROR_OBJ }
