package main

import (
	"complete/complete/completer"
	"complete/complete/completeArgs"
	"complete/complete/flag"
	"complete/complete/positional"
	"complete/complete/types"
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
						Cmds: []types.Completer{
							{
								Flags: []flag.Flag{
									{
										Short:    'p',
										Long:     "path",
										Help:     "path to file to process",
										Args:     []completeargs.CompleteType{
											completeargs.CTfile{
												Glob:     "\\*.yaml",
												OnlyDirs: false,
											},
										},
										Optional: true,
									},
									{
										Short:    'c',
										Long:     "color",
										Help:     "colorful output",
										Args:     []completeargs.CompleteType{
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
