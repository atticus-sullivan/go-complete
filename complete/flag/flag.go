package flag

import (
	"errors"
	"strings"

	"github.com/atticus-sullivan/go-complete/complete/completeArgs"
	sh"mvdan.cc/sh/v3/syntax"
)

// you should create instances of this with [NewFlag] or [MustNewFlag].
// Alternatively, the preprocessing and validating methods are exported so if
// really needed you can construct Flag instances yourself.
type Flag struct {
	Short    rune
	Long     string
	Help     string
	Args     []completeargs.CompleteTypeNoSub
	Optional bool
}

func (f *Flag) Preprocess() error {
	var err error
	f.Help, err = sh.Quote(f.Help, sh.LangBash)
	return err
}

func (f Flag) Validate() error {
	if strings.ContainsAny(f.Long, " \t\n\r") {
		return errors.New("flag.Long must not contain any of '\\t', '\\s', ' ', '\\r'")
	}
	return nil
}

func NewFlag(fo ...FlagOpt) (Flag, error) {
	f := Flag{}
	for _,i := range fo {
		i(&f)
	}
	if err := f.Validate(); err != nil {
		return f, err
	}
	if err := f.Preprocess(); err != nil {
		return f, err
	}
	return f, nil
}

func MustNewFlag(fo ...FlagOpt) Flag {
	f, err := NewFlag(fo...)
	if err != nil {
		panic(err)
	}
	return f
}

type FlagOpt func(*Flag)

func FlagWithShort(short rune) FlagOpt {
	return func(f *Flag) {
		f.Short = short
	}
}

func FlagWithLong(long string) FlagOpt {
	return func(f *Flag) {
		f.Long = long
	}
}

func FlagWithHelp(help string) FlagOpt {
	return func(f *Flag) {
		f.Help = help
	}
}

func FlagWithArgs(args []completeargs.CompleteTypeNoSub) FlagOpt {
	return func(f *Flag) {
		f.Args = args
	}
}

func FlagWithOptional(optional bool) FlagOpt {
	return func(f *Flag) {
		f.Optional = optional
	}
}
