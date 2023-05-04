package format

import "fmt"

type Person struct {
	name string
	age  int
}

func General() {
	var (
		i      int       = 0
		f      float64   = 0.5
		flag   bool      = false
		str    string    = "Hello"
		arr    [3]string = [3]string{"a", "b", "c"}
		person Person    = Person{
			name: "ken",
			age:  23,
		}
		m map[string]int = map[string]int{"one": 1, "two": 2}
	)
	var variables []interface{} = []interface{}{
		i, f, flag, str, arr, person, m,
	}

	PrintTitle("%v - デフォルトフォーマット")
	for _, v := range variables {
		fmt.Printf("%v\n", v)
	}
	//----------------------------------------
	// %%v - デフォルトフォーマット
	// ----------------------------------------
	// 0
	// 0.5
	// false
	// Hello
	// [a b c]
	// {ken 23}
	// map[one:1 two:2]

	PrintTitle("%+v - デフォルトフォーマット（構造体にはフィールド名つき） ")
	for _, v := range variables {
		fmt.Printf("%+v\n", v)
	}
	//----------------------------------------
	// %%+v - デフォルトフォーマット(フィールド名つき)
	// ----------------------------------------
	// 0
	// 0.5
	// false
	// Hello
	// [a b c]
	// {name:ken age:23}
	// map[one:1 two:2]

	PrintTitle("%#v - Goの構文に合わせたフォーマット")
	for _, v := range variables {
		fmt.Printf("%#v\n", v)
	}
	//----------------------------------------
	// %#v - Goの構文に合わせたフォーマット
	// ----------------------------------------
	// 0
	// 0.5
	// false
	// "Hello"
	// [3]string{"a", "b", "c"}
	// format.Person{name:"ken", age:23}
	// map[string]int{"one":1, "two":2}

	PrintTitle("%T - 型(type)をGoの構文で出力")
	for _, v := range variables {
		fmt.Printf("%T\n", v)
	}
	// ----------------------------------------
	// %%T - typeを出力
	// ----------------------------------------
	// int
	// float64
	// bool
	// string
	// [3]string
	// format.Person
	// map[string]int

	PrintTitle("%% - 「%」を出力")
	fmt.Printf("%% ←これはパーセント\n")
	// ----------------------------------------
	// %% - 「%」を出力
	// ----------------------------------------
	// % ←これはパーセント
}
