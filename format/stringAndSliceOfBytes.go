package format

import "fmt"

func StringAndSliceOfBytes() {
	var str string = "Hello"
	var sliceOfBytes []byte = []byte("abcd")
	// = []byte{97, 98, 99, 100}

	PrintTitle("%s - []byteとして解釈される前の（加工されていないそのままの）表示。")
	fmt.Printf("string: %s, slice of bytes: %s\n", str, sliceOfBytes)
	// ----------------------------------------
	// %s - 解釈される前の（加工されていないそのままの）表示。
	// ----------------------------------------
	// string: Hello, slice of bytes: abcd

	PrintTitle("%q - Go構文で安全にエスケープされたダブルクォート囲みの文字列。")
	fmt.Printf("string: %q, slice of bytes: %q\n", str, sliceOfBytes)
	// ----------------------------------------
	// %q - Go構文で安全にエスケープされたダブルクォート囲みの文字列。
	// ----------------------------------------
	// string: "Hello", slice of bytes: "abcd"

	PrintTitle("%x - 小文字を使う16進数表記。(1バイトにつき2文字)")
	fmt.Printf("string: %x, slice of bytes: %x\n", str, sliceOfBytes)
	// ----------------------------------------
	// %x - 小文字を使う16進数表記。1バイトにつき2文字。
	// ----------------------------------------
	// string: 48656c6c6f, slice of bytes: 61626364

	PrintTitle("%X - 大文字を使う16進数表記。(1バイトにつき2文字)")
	fmt.Printf("string: %X, slice of bytes: %X\n", str, sliceOfBytes)
	// ----------------------------------------
	// %X - 大文字を使う16進数表記。(1バイトにつき2文字)
	// ----------------------------------------
	// string: 48656C6C6F, slice of bytes: 61626364
	// ----------------------------------------

}
