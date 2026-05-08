package steamworks

import "unsafe"

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// C char* 转 Go string
func CStrToString(cstr uintptr) string {
	if cstr == 0 {
		return ""
	}
	// 读取 C 字符串直到 \0
	var b []byte
	for i := 0; ; i++ {
		ch := *(*byte)(unsafe.Pointer(cstr + uintptr(i)))
		if ch == 0 {
			break
		}
		b = append(b, ch)
	}
	return string(b)
}
