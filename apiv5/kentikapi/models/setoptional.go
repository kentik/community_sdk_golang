package models

import "reflect"

// SetOptional sets output pointer to a given value that can be a constant. This is a shortcut for eg.:
// name := "device name"
// device.DeviceName = &name
// output must be pointer to a pointer
func SetOptional(output interface{}, value interface{}) {
	tOutput := reflect.TypeOf(output)
	if tOutput.Kind() != reflect.Ptr || tOutput.Elem().Kind() != reflect.Ptr {
		panic("setOptional output argument must be pointer to a pointer")
	}

	vInput := reflect.ValueOf(value)
	vOutput := reflect.ValueOf(output)
	outputType := tOutput.Elem().Elem() // ie. *output
	fieldPtr := reflect.New(outputType)
	fieldPtr.Elem().Set(vInput)
	vOutput.Elem().Set(fieldPtr)
}
