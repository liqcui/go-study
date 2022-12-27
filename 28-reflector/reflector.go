package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	v1 := reflect.TypeOf(x)
	fmt.Printf("The type is:%v\n", v1)
	v2 := reflect.ValueOf(x)
	fmt.Printf("The value is:%v\n", v2)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		//v.Int()从反射中获取类型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		//v.Float()从放射中获取类型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		//v.Float()从放射中获取类型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
	}
}

//通过反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func main() {
	str := `{name:"Jack","age":900}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)

	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	reflectValue(a)
	reflectSetValue(&b)
	fmt.Println(b)
	
	reflectSetValue1(&b)
	fmt.Println(b)
}
