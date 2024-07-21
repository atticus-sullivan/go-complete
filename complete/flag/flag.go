package flag

import (
	"complete/complete/completeArgs"
)

type Flag struct {
	Short    rune
	Long     string
	Help     string
	Args     []completeargs.CompleteType
	Optional bool
}
