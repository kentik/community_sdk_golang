package utils

import "reflect"

// ConvertFunc signature like: func ToUpper(source string) (string, error)
type ConvertFunc interface{}

// ConvertList transforms input list of items into output list using convertFunc.
// "input" must be array or slice
// "output" must be a pointer to slice; a new slice is allocated and returned under that pointer
func ConvertList(input interface{}, convertFunc ConvertFunc, output interface{}) error {
	tInput := reflect.TypeOf(input)
	if tInput.Kind() != reflect.Array && tInput.Kind() != reflect.Slice {
		panic("ConvertList input argument must be array or slice")
	}
	tOutput := reflect.TypeOf(output)
	if tOutput.Kind() != reflect.Ptr || tOutput.Elem().Kind() != reflect.Slice {
		panic("ConvertList output argument must be pointer to slice")
	}
	// todo: validate convertFunc

	vInput := reflect.ValueOf(input)
	vOutput := reflect.ValueOf(output).Elem()
	vConvertFunc := reflect.ValueOf(convertFunc)

	itemCount := vInput.Len()
	itemType := vConvertFunc.Type().Out(0)
	vOutput.Set(reflect.MakeSlice(reflect.SliceOf(itemType), itemCount, itemCount))

	for i := 0; i < itemCount; i++ {
		args := []reflect.Value{vInput.Index(i)}
		results := vConvertFunc.Call(args)

		if err, ok := results[1].Interface().(error); ok {
			return err
		}

		vOutput.Index(i).Set(results[0])
	}
	return nil
}

// ConvertOrNone transforms input into output using convertFunc, unless input is nil ->
// "input" must be a pointer eg *int
// "output" must be a pointer to pointer eg **int
// conversion result is stored under output, or nil is set if input is nil
func ConvertOrNone(input interface{}, convertFunc ConvertFunc, output interface{}) error {
	tInput := reflect.TypeOf(input)
	if tInput.Kind() != reflect.Ptr {
		panic("ConvertOrNone input argument must be a pointer")
	}
	tOutput := reflect.TypeOf(output)
	if tOutput.Kind() != reflect.Ptr || tOutput.Elem().Kind() != reflect.Ptr {
		panic("ConvertOrNone output argument must be pointer to pointer")
	}
	// todo: validate convertFunc

	vInput := reflect.ValueOf(input)
	vOutput := reflect.ValueOf(output).Elem()
	vConvertFunc := reflect.ValueOf(convertFunc)
	outputType := vConvertFunc.Type().Out(0)

	if vInput.IsNil() {
		nilResult := reflect.Zero(reflect.PtrTo(outputType))
		vOutput.Set(nilResult)
		return nil
	}

	args := []reflect.Value{vInput.Elem()}
	results := vConvertFunc.Call(args)

	if err, ok := results[1].Interface().(error); ok {
		return err
	}

	// allocate output pointer
	vOutput.Set(reflect.New(outputType))

	// set value under that pointer's memory
	vOutput.Elem().Set(results[0])
	return nil
}
