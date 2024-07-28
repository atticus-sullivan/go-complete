package flag

import (
	"github.com/atticus-sullivan/go-complete/complete/completeArgs"
)

type Flag struct {
	Short    rune
	Long     string
	Help     string
	Args     []completeargs.CompleteTypeNoSub
	Optional bool
}
