package format

import (
	"errors"
	"fmt"
)

func Error() {
	Error := errors.New("an error occurred")
	WrapError := fmt.Errorf("wrap!!: %w", Error)
	PrintTitle("%w - エラーのラップ")
	fmt.Println(WrapError)
	// ----------------------------------------
	// %w - エラーのラップ
	// ----------------------------------------
	// wrap!!: an error occurred
}
