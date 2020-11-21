package main

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	str := "Cgg3NDUyNzgwMRKQAUJGbDEvNjZUSzM4KzBHOU51Uk5lRFZ3WDMyMjk1TERqbGZiaDZGaGlwMVJpUXk5RnVBRnpmcjlRcGVQUVhuNnNicVA1bXBhTTdPMW5YVUxySWpjdnhQallNaFRsWXhWZENjVmE5UWZlL3gxQWZmVEN2bGxWQ2Q1OWdCcjdTbFM1WGxoclN5WndCWlRqRGc9PRkAAAAAAGo4QSEAAAAAwEo4QSoTMjAyMS0wMi0wMlQwMDowMDowMDIOY3JlZGl0NzQ1Mjc4MDE6EzIwMjAtMDctMjRUMDk6NDU6MDlCEzIwMjAtMTAtMTVUMDU6NDg6MzZKCUFDVElWQVRFRFEAAAAAAEC/QA=="
	dd, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println("解码后的信息为:", dd)

	data := &QueryCreditInfoRespVO{}
	proto.Unmarshal(dd, data)
	fmt.Println("反序列化之后的信息为：", data)
}
