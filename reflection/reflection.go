package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getVal(x)

	numberOfValues := 0
	getVal := func(i int) reflect.Value {
		return val
	}

	switch val.Kind() {
	case reflect.Struct:
		numberOfValues = val.NumField()
		getVal = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getVal = val.Index
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getVal(i).Interface(), fn)
	}
}

func getVal(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
