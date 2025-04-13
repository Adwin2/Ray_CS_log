package main

import (
	"fmt"
	"reflect"
)

type config struct {
	dir        string
	dbfileName string
}

func main() {
	conf := &config{
		dir:        "test",
		dbfileName: "dbtest",
	}
	val := reflect.ValueOf(conf).Elem().FieldByName("dir") //结构体实例的反射值
	getVal := val.String()
	fmt.Println(len([]byte(getVal)), getVal)

	fmt.Println(val)
	// fmt.Println(getField(val, "dir"))
	// fmt.Println(getField(val, "dbfilename"))
	// fmt.Println(val.FieldByName("dir"))
}

// func getField(val reflect.Value, field string) reflect.Value {
// 	typ := val.Type() // 获取结构体类型信息
// 	log.Print(typ) // OP: main.config

// 	for i := 0;i<typ.NumField();i++ {
// 		structField := typ.Field(i) // 获取字段元信息
// 		log.Print(structField)  // OP: {dir main string config:"dir" 0 [0] false} or
// 		fieldName := structField.Name // 字段名
// 		log.Printf(fieldName)
// 		tag := structField.Tag.Get("dir") // 标签值

// 		// 匹配标签字段值
// 		if tag == field || strings.EqualFold(fieldName, field){
// 			return val.Field(i) // 返回字段值
// 		}
// 	}

// 	return reflect.Value{} // 未找到返回零值
// }
