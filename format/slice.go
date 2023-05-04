package format

import "fmt"

func Slice() {
	var slice []string = []string{"Hello", "World"}
	PrintTitle("%p - スライスの先頭要素のアドレスを16進数表記で表示。先頭に0xが付く。")
	fmt.Printf("%p\n", slice)
	fmt.Printf("変数sliceの先頭要素のアドレスは%#p\n", &slice[0])
	// ----------------------------------------
	// %p - スライスの先頭要素のアドレスを16進数表記で表示。先頭に0xが付く。
	// ----------------------------------------
	// 0xc00011a000
	// 変数sliceの先頭要素のアドレスはc00011a000
}
