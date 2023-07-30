package evaluator

import (
	"bufio"
	"fmt"
	"github.com/grahms/samoralang/object"
	"os"
	"strconv"
	"strings"
)

var builtins = map[string]*object.Builtin{
	"len": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}
		switch arg := args[0].(type) {
		case *object.String:
			return &object.Integer{Value: int64(len(arg.Value))}
		case *object.Array:
			return &object.Integer{Value: int64(len(arg.Elements))}
		default:
			return newError("argument to `len` not supported, got %s", args[0].Type())
		}

	}},
	"first": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `first` must be ARRAY, got %s",
				args[0].Type())
		}
		arr := args[0].(*object.Array)
		if len(arr.Elements) > 0 {
			return arr.Elements[0]
		}
		return NULL
	},
	},
	"last": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `last` must be ARRAY, got %s",
				args[0].Type())
		}
		arr := args[0].(*object.Array)
		length := len(arr.Elements)
		if length > 0 {
			return arr.Elements[length-1]
		}
		return NULL
	},
	},
	"rest": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `rest` must be ARRAY, got %s",
				args[0].Type())
		}
		arr := args[0].(*object.Array)
		length := len(arr.Elements)
		if length > 0 {
			newElements := make([]object.Object, length-1, length-1)
			copy(newElements, arr.Elements[1:length])
			return &object.Array{Elements: newElements}
		}
		return NULL
	},
	},
	"push": {Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2",
				len(args))
		}
		if args[0].Type() != object.ArrayObj {
			return newError("argument to `push` must be ARRAY, got %s",
				args[0].Type())
		}
		arr := args[0].(*object.Array)
		length := len(arr.Elements)
		newElements := make([]object.Object, length+1, length+1)
		copy(newElements, arr.Elements)
		newElements[length] = args[1]
		return &object.Array{Elements: newElements}
	},
	},
	"print":   {Fn: printFunc},
	"println": {Fn: printlnFunc},
	"input":   {Fn: inputFunc},
	"int":     {Fn: intFunc},
	"str":     {Fn: strFunc},

	"readFile":   {Fn: readFileFunc},
	"writeFile":  {Fn: writeFileFunc},
	"removeFile": {Fn: removeFileFunc},
	"readDir":    {Fn: readDirFunc},
}

func inputFunc(args ...object.Object) object.Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. got=%d, want=0",
			len(args))
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return newError("failed to read input: %s", err.Error())
	}
	return &object.String{Value: strings.TrimSpace(input)}
}

func printFunc(args ...object.Object) object.Object {
	for _, arg := range args {
		switch arg := arg.(type) {
		case *object.Integer:
			fmt.Printf("%d", arg.Value)
		case *object.Float:
			fmt.Printf("%.2f", arg.Value) // print float with 2 decimal places
		case *object.String:
			fmt.Printf("%s", arg.Value)
		default:
			fmt.Printf("%s", arg.Inspect())
		}
	}
	return NULL
}

func printlnFunc(args ...object.Object) object.Object {
	for _, arg := range args {
		str := arg.Inspect()
		str = strings.ReplaceAll(str, "\\n", "\n")
		str = strings.ReplaceAll(str, "\\t", "\t")
		println(str)
	}
	return NULL
}

func intFunc(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	strObj, ok := args[0].(*object.String)
	if !ok {
		return newError("argument to `int` must be STRING, got %s", args[0].Type())
	}
	str := strObj.Value
	num, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return newError("failed to convert string to int: %s", err.Error())
	}
	return &object.Integer{Value: num}
}

func strFunc(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	intObj, ok := args[0].(*object.Integer)
	if !ok {
		return newError("argument to `str` must be INTEGER, got %s", args[0].Type())
	}
	str := strconv.FormatInt(intObj.Value, 10)
	return &object.String{Value: str}
}
