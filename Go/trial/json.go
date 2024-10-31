package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//JSON是一种用于发送和接受结构化信息的标准协议
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	//`Tag` 对应效果 Year变为released, color输出变为小写
	var movies = []Movie{
		{Title: "Castle", Year: 1967, Color: false, Actors: []string{"Humphrey", "Bergman"}}, {Title: "col", Year: 1966, Color: true, Actors: []string{"paul Newman"}}}
	//将形如movies的结构体slice转为JSON的过程叫编组(marshaling)。通过调用json.Marshal()函数完成 \.MarshalIndent() 加缩进
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//编码对应解码 unmarshaling json.Unmarshal() ,可以选择性解码 如下只解码输出对应title的内容
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed :%s", err)
	}
	fmt.Println(titles)

	/*
		许多web服务都提供JSON接口，通过HTTP接口发送JSON格式请求并返回JSON格式的信息。
	*/
}
