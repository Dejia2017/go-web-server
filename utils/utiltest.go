package utils

import "fmt"

func UtilTest() {
	LogError("error测试")
	var list string
	SetCache("mykeytest","hello world",10)
	//utils.DelCache("mykeytest")
	GetCache("mykeytest",&list)
	fmt.Println(list)

	str2 := String2md5("hello")
	str1 := RandomString(5)
	fmt.Println(str1)
	fmt.Println(str2)
}
