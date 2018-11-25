package main

//可以给类型起别名，增加可读性。
type (
	byte int8  //byte是uint8的别名。这里是说明这个是可以做到覆盖默认的别名
	rune int32
	ByteSize int64
	中文 string  //这里是说明可以做到。不建议这样使用
)

func main() {

}