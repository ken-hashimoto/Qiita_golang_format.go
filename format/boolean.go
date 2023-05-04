package format

import "fmt"

func Boolean() {
	var T bool = true
	var F bool = false
	PrintTitle("%t - trueまたはfalseで出力")
	fmt.Printf("%t,%t\n", T, F)
	// ----------------------------------------
	// %t - trueまたはfalseで出力
	// ----------------------------------------
	// true,false
}
