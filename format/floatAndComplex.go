package format

import "fmt"

func FloatAndComplex() {
	var f float64 = 0.5
	var c complex128 = complex(3, 4) // 3+4i
	var bigF float64 = 1234567.89
	var bigC complex128 = complex(12345.678, 9876543.21) // 12345.678 + 9876543.21i
	PrintTitle("%b - 符号付き指数表記。指数は2のべき乗で、strconv.FormatFloatの'b'フォーマットと同様に表される。例：-123456p-78")
	fmt.Printf("float64:      %b, complex:      %b\n", f, c)
	fmt.Printf("float64(Big): %b, complex(Big): %b\n", bigF, bigC)
	// ----------------------------------------
	// %b - 符号付き指数表記。指数は2のべき乗で、strconv.FormatFloatの'b'フォーマットと同様に表され る。例：-123456p-78
	// ----------------------------------------
	// float64:      4503599627370496p-53, complex:      (6755399441055744p-51+4503599627370496p-50i)
	// float64(Big): 5302428712241725p-32, complex(Big): (6787108256889176p-39+5302428760560108p-29i)

	PrintTitle("%e - 指数表記。例：-1.234456e+78")
	fmt.Printf("float64:      %e, complex:      %e\n", f, c)
	fmt.Printf("float64(Big): %e, complex(Big): %e\n", bigF, bigC)
	// ----------------------------------------
	// %e - 指数表記。例：-1.234456e+78
	// ----------------------------------------
	// float64:      5.000000e-01, complex:      (3.000000e+00+4.000000e+00i)
	// float64(Big): 1.234568e+06, complex(Big): (1.234568e+04+9.876543e+06i)

	PrintTitle("%E - eを大文字にした指数表記。例：-1.234456E+78")
	fmt.Printf("float64:      %E, complex:      %E\n", f, c)
	fmt.Printf("float64(Big): %E, complex(Big): %E\n", bigF, bigC)
	// ----------------------------------------
	// %E - eを大文字にした指数表記。例：-1.234456E+78
	// ----------------------------------------
	// float64:      5.000000E-01, complex:      (3.000000E+00+4.000000E+00i)
	// float64(Big): 1.234568E+06, complex(Big): (1.234568E+04+9.876543E+06i)

	PrintTitle("%f - 小数点表記、指数は使わない。例：123.456")
	fmt.Printf("float64:      %f, complex:      %f\n", f, c)
	fmt.Printf("float64(Big): %f, complex(Big): %f\n", bigF, bigC)
	// ----------------------------------------
	// %f - 小数点表記、指数は使わない。例：123.456
	// ----------------------------------------
	// float64:      0.500000, complex:      (3.000000+4.000000i)
	// float64(Big): 1234567.890000, complex(Big): (12345.678000+9876543.210000i)

	PrintTitle("%F - %fの別名")
	fmt.Printf("float64:      %F, complex:      %F\n", f, c)
	fmt.Printf("float64(Big): %F, complex(Big): %F\n", bigF, bigC)
	// ----------------------------------------
	// %F - %fの別名
	// ----------------------------------------
	// float64:      0.500000, complex:      (3.000000+4.000000i)
	// float64(Big): 1234567.890000, complex(Big): (12345.678000+9876543.210000i)

	PrintTitle("%g - 指数が大きい場合は%e、そうでない場合は%f")
	fmt.Printf("float64:      %g, complex:      %g\n", f, c)
	fmt.Printf("float64(Big): %g, complex(Big): %g\n", bigF, bigC)
	// ----------------------------------------
	// %g - 指数が大きい場合は%e、そうでない場合は%f
	// ----------------------------------------
	// float64:      0.5, complex:      (3+4i)
	// float64(Big): 1.23456789e+06, complex(Big): (12345.678+9.87654321e+06i)

	PrintTitle("%G - 指数が大きい場合は%E、そうでない場合は%F")
	fmt.Printf("float64:      %G, complex:      %G\n", f, c)
	fmt.Printf("float64(Big): %G, complex(Big): %G\n", bigF, bigC)
	// ----------------------------------------
	// %G - 指数が大きい場合は%E、そうでない場合は%F
	// ----------------------------------------
	// float64:      0.5, complex:      (3+4i)
	// float64(Big): 1.23456789E+06, complex(Big): (12345.678+9.87654321E+06i)

	PrintTitle("%x - 16進数表記（指数は2のべき乗の10進数で表される）。例：-0x1.23abcp+20")
	fmt.Printf("float64:      %x, complex:      %x\n", f, c)
	fmt.Printf("float64(Big): %x, complex(Big): %x\n", bigF, bigC)
	// ----------------------------------------
	// %x - 16進数表記（指数は2のべき乗の10進数で表される）。例：-0x1.23abcp+20
	// ----------------------------------------
	// float64:      0x1p-01, complex:      (0x1.8p+01+0x1p+02i)
	// float64(Big): 0x1.2d687e3d70a3dp+20, complex(Big): (0x1.81cd6c8b43958p+13+0x1.2d687e6b851ecp+23i)

	PrintTitle("%X - 大文字の16進数表記。例：-0X1.23ABCP+20")
	fmt.Printf("float64:      %X, complex:      %X\n", f, c)
	fmt.Printf("float64(Big): %X, complex(Big): %X\n", bigF, bigC)
	// ----------------------------------------
	// %X - 大文字の16進数表記。例：-0X1.23ABCP+20
	// ----------------------------------------
	// float64:      0X1P-01, complex:      (0X1.8P+01+0X1P+02i)
	// float64(Big): 0X1.2D687E3D70A3DP+20, complex(Big): (0X1.81CD6C8B43958P+13+0X1.2D687E6B851ECP+23i)
}
