package evaluator

import (
	"github.com/grahms/samoralang/object"
	"os"
)

func readFileFunc(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	filePathObj, ok := args[0].(*object.String)
	if !ok {
		return newError("argument to `readFile` must be STRING, got %s", args[0].Type())
	}

	filePath := filePathObj.Value

	content, err := os.ReadFile(filePath)
	if err != nil {
		return newError("failed to read file: %s", err.Error())
	}

	return &object.String{Value: string(content)}
}

func writeFileFunc(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}

	filePathObj, ok := args[0].(*object.String)
	if !ok {
		return newError("first argument to `writeFile` must be STRING, got %s", args[0].Type())
	}

	contentObj, ok := args[1].(*object.String)
	if !ok {
		return newError("second argument to `writeFile` must be STRING, got %s", args[1].Type())
	}

	filePath := filePathObj.Value
	content := contentObj.Value

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return newError("failed to write file: %s", err.Error())
	}

	return NULL
}

func removeFileFunc(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	filePathObj, ok := args[0].(*object.String)
	if !ok {
		return newError("argument to `removeFile` must be STRING, got %s", args[0].Type())
	}

	filePath := filePathObj.Value

	err := os.Remove(filePath)
	if err != nil {
		return newError("failed to remove file: %s", err.Error())
	}

	return NULL
}

func readDirFunc(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	dirPathObj, ok := args[0].(*object.String)
	if !ok {
		return newError("argument to `readDir` must be STRING, got %s", args[0].Type())
	}

	dirPath := dirPathObj.Value

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return newError("failed to read directory: %s", err.Error())
	}

	var fileNames []object.Object
	for _, file := range files {
		fileNames = append(fileNames, &object.String{Value: file.Name()})
	}

	return &object.Array{Elements: fileNames}
}
