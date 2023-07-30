package object

type BuiltInFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltInFunction
}

func (b *Builtin) Type() ObjectType { return BuiltinObj }
func (b *Builtin) Inspect() string  { return "builtin function" }
