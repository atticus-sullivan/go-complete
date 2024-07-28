package main

import (
	"github.com/atticus-sullivan/go-complete/complete/completer"
	"github.com/atticus-sullivan/go-complete/complete/completeArgs"
	"github.com/atticus-sullivan/go-complete/complete/flag"
	"github.com/atticus-sullivan/go-complete/complete/positional"
	"github.com/atticus-sullivan/go-complete/complete/types"
	"fmt"
	"strings"
)

func main() {
	c := completer.MainCompleter{
		Completer: types.Completer{
			Flags: []flag.Flag{
				{
					Short:    'h',
					Long:     "help",
					Help:     "Show help information",
					Args:     nil,
					Optional: true,
				},
			},
			Positionals: []positional.Positional{
				{
					Idx:  0,
					Help: "",
					Arg: types.CTsubcommands{
						Idx: 0,
						Cmds: []types.Completer{
							{
								Flags: []flag.Flag{
									{
										Short:    'p',
										Long:     "path",
										Help:     "path to file to process",
										Args:     []completeargs.CompleteTypeNoSub{
											completeargs.CTfile{
												Glob:     "*.yaml",
												OnlyDirs: false,
											},
										},
										Optional: true,
									},
									{
										Short:    'c',
										Long:     "color",
										Help:     "colorful output",
										Args:     []completeargs.CompleteTypeNoSub{
											completeargs.CTalternatives{
												Alts: []string{"hello", "world"},
											},
										},
										Optional: true,
									},
								},
								Positionals: []positional.Positional{},
								Name:        "process",
							},
							{
								Flags:       []flag.Flag{},
								Positionals: []positional.Positional{},
								Name:        "preprocess",
							},
						},
					},
					Optional: false,
				},
			},
			Name: "testing",
		},
		Opts: completer.CompleterOpts{
			OptionStacking:     false,
			OptionStackingArgs: false,
			DashSep:            false,
		},
	}

	builder := strings.Builder{}
	c.GenerateZsh(&builder, "")
	fmt.Println(builder.String())
}
