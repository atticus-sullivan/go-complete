package positional

import (
	completeargs "github.com/atticus-sullivan/go-complete/complete/completeArgs"
	sh "mvdan.cc/sh/v3/syntax"
)

// you should create instances of this with [NewPositional] or [MustNewPositional].
// Alternatively, the preprocessing and validating methods are exported so if
// really needed you can construct Positional instances yourself.
type Positional struct {
	Idx      uint
	Help     string
	Arg      completeargs.CompleteTypeSub
	Optional bool
}

func (f *Positional) Preprocess() error {
	var err error
	if f.Arg == nil {
		f.Arg = completeargs.CTempty{}
	}
	f.Help, err = sh.Quote(f.Help, sh.LangBash)
	return err
}

func (f Positional) Validate() error {
	return nil
}

func NewPositional(idx uint, fo ...PositionalOpt) (Positional, error) {
	f := Positional{Idx: idx}
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

func MustNewPositional(idx uint, fo ...PositionalOpt) Positional {
	f, err := NewPositional(idx, fo...)
	if err != nil {
		panic(err)
	}
	return f
}

type PositionalOpt func(*Positional)

func PositionalWithHelp(help string) PositionalOpt {
	return func(f *Positional) {
		f.Help = help
	}
}

func PositionalWithArgs(arg completeargs.CompleteTypeSub) PositionalOpt {
	return func(f *Positional) {
		f.Arg = arg
	}
}

func PositionalWithOptional(optional bool) PositionalOpt {
	return func(f *Positional) {
		f.Optional = optional
	}
}
