// Copyright 2015-2016 gotokatsuya GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package treflect

import (
	"fmt"
	"reflect"
)

// GetMissingMethods checks if a given object implements all methods of a
// given interface. It returns the interface coverage [0..1] as well as an array
// of error messages. If the interface is correctly implemented the coverage is
// 1 and the error message array is empty.
func GetMissingMethods(objType reflect.Type, ifaceType reflect.Type) (float32, []string) {
	missing := []string{}
	if objType.Implements(ifaceType) {
		return 1.0, missing
	}

	methodCount := ifaceType.NumMethod()
	for mIdx := 0; mIdx < methodCount; mIdx++ {
		ifaceMethod := ifaceType.Method(mIdx)
		objMethod, exists := objType.MethodByName(ifaceMethod.Name)
		signatureMismatch := false

		switch {
		case !exists:
			missing = append(missing, fmt.Sprintf("Missing: \"%s\" %v", ifaceMethod.Name, ifaceMethod.Type))
			continue // ### continue, error found ###

		case ifaceMethod.Type.NumOut() != objMethod.Type.NumOut():
			signatureMismatch = true

		case ifaceMethod.Type.NumIn()+1 != objMethod.Type.NumIn():
			signatureMismatch = true

		default:
			for oIdx := 0; !signatureMismatch && oIdx < ifaceMethod.Type.NumOut(); oIdx++ {
				signatureMismatch = ifaceMethod.Type.Out(oIdx) != objMethod.Type.Out(oIdx)
			}
			for iIdx := 0; !signatureMismatch && iIdx < ifaceMethod.Type.NumIn(); iIdx++ {
				signatureMismatch = ifaceMethod.Type.In(iIdx) != objMethod.Type.In(iIdx+1)
			}
		}

		if signatureMismatch {
			missing = append(missing, fmt.Sprintf("Invalid: \"%s\" %v is not %v", ifaceMethod.Name, objMethod.Type, ifaceMethod.Type))
		}
	}

	return float32(methodCount-len(missing)) / float32(methodCount), missing
}

// Int64 converts any signed number type to an int64.
// The second parameter is returned as false if a non-number type was given.
func Int64(v interface{}) (int64, bool) {

	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return int64(v.(int)), true
	case reflect.Int8:
		return int64(v.(int8)), true
	case reflect.Int16:
		return int64(v.(int16)), true
	case reflect.Int32:
		return int64(v.(int32)), true
	case reflect.Int64:
		return v.(int64), true
	case reflect.Float32:
		return int64(v.(float32)), true
	case reflect.Float64:
		return int64(v.(float64)), true
	}

	fmt.Printf("%t\n%#v\n%#v\n", v, v, reflect.TypeOf(v).Kind())

	return 0, false
}

// Uint64 converts any unsigned number type to an uint64.
// The second parameter is returned as false if a non-number type was given.
func Uint64(v interface{}) (uint64, bool) {

	switch reflect.TypeOf(v).Kind() {
	case reflect.Uint:
		return uint64(v.(uint)), true
	case reflect.Uint8:
		return uint64(v.(uint8)), true
	case reflect.Uint16:
		return uint64(v.(uint16)), true
	case reflect.Uint32:
		return uint64(v.(uint32)), true
	case reflect.Uint64:
		return v.(uint64), true
	}

	return 0, false
}

// Float32 converts any number type to an float32.
// The second parameter is returned as false if a non-number type was given.
func Float32(v interface{}) (float32, bool) {

	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return float32(v.(int)), true
	case reflect.Uint:
		return float32(v.(uint)), true
	case reflect.Int8:
		return float32(v.(int8)), true
	case reflect.Uint8:
		return float32(v.(uint8)), true
	case reflect.Int16:
		return float32(v.(int16)), true
	case reflect.Uint16:
		return float32(v.(uint16)), true
	case reflect.Int32:
		return float32(v.(int32)), true
	case reflect.Uint32:
		return float32(v.(uint32)), true
	case reflect.Int64:
		return float32(v.(int64)), true
	case reflect.Uint64:
		return float32(v.(uint64)), true
	case reflect.Float32:
		return v.(float32), true
	case reflect.Float64:
		return float32(v.(float64)), true
	}

	return 0, false
}

// Float64 converts any number type to an float64.
// The second parameter is returned as false if a non-number type was given.
func Float64(v interface{}) (float64, bool) {

	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return float64(v.(int)), true
	case reflect.Uint:
		return float64(v.(uint)), true
	case reflect.Int8:
		return float64(v.(int8)), true
	case reflect.Uint8:
		return float64(v.(uint8)), true
	case reflect.Int16:
		return float64(v.(int16)), true
	case reflect.Uint16:
		return float64(v.(uint16)), true
	case reflect.Int32:
		return float64(v.(int32)), true
	case reflect.Uint32:
		return float64(v.(uint32)), true
	case reflect.Int64:
		return float64(v.(int64)), true
	case reflect.Uint64:
		return float64(v.(uint64)), true
	case reflect.Float32:
		return float64(v.(float32)), true
	case reflect.Float64:
		return v.(float64), true
	}

	return 0, false
}

// RemovePtrFromType will return the type of t and strips away any pointer(s)
// in front of the actual type.
func RemovePtrFromType(t interface{}) reflect.Type {
	var v reflect.Type
	if rt, isType := t.(reflect.Type); isType {
		v = rt
	} else {
		v = reflect.TypeOf(t)
	}
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

// RemovePtrFromValue will return the value of t and strips away any pointer(s)
// in front of the actual type.
func RemovePtrFromValue(t interface{}) reflect.Value {
	var v reflect.Value
	if rv, isValue := t.(reflect.Value); isValue {
		v = rv
	} else {
		v = reflect.ValueOf(t)
	}
	for v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
