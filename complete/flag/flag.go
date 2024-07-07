package flag

import (
	"complete/internal"
)

type Flag struct {
	Short    rune
	Long     string
	Help     string
	Args     []internal.CompleteType
	Optional bool
}
