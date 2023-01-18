package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHaha(t *testing.T) {
	var a int64 = 1
	var b = 1.0
	reflectValue(a)
	reflectValue(b)
}

func reflectValue(obj any) {
	var v = reflect.ValueOf(obj)
	var k = v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
