> [!CAUTION]
> Note that this is still in an experimental state. Certain inputs might break things.
> If you're still willing to test this, your're more than welcome to help in the
> development and especially to harden this library against bad inputs / bad use.

# go-complete

Implements support for generating shell (bash+zsh) completion scripts. See this
[issue/comment](https://github.com/alexflint/go-arg/issues/186#issuecomment-2211750247)

Goal is to have this as a library that can be used. For now see `main.go` for an
example.

## How to use / test
For the most part, this is thought as a library. Still it comes with a `main.go`
which can serve for testing.

You can quickly install the completion by executing `eval "$(go run main.go)"`
(on the long run you should write the output to a file and place it somewhere in
your `PATH` (bash)/`fpath` (zsh, make sure the filename is right)). Remember
calling `eval` usually is considered evil, so you might want to have a look at
the generated output first.

The `main.go` in the uploaded form installs an example completion for a
function/program `testing`. So you might want to define a dummy function
`testing`: `function testing() {}`.

With this setup, after typing `testing <tab><tab>` you should get the first
completion results.

## ToDos
In general this should be near to be complete. Still we don't sanitize/escape
strings passed as parameters/arguments (see [#1](https://github.com/atticus-sullivan/go-complete/issues/1)).

## Small reference for contributers
### zsh
```
function _command {
    do some inizialization
    generate the _arguments call by calling flags.GenerateZsh() and positionals.GenerateZsh()

    execute collected AddFunc.Cas() (switch over the positionals and call functions)
}
register completion function

execute AddFunc.Fun()
```

Note that `AddFunc`s are collected when calling `GenerateZsh()`.
They are used to be able to generate/inject code to another place in the
template.

### bash
The generated file generally looks like this
```
function _command {
    do some inizialization
    for words
        execute positionals -> positionals.GenerateBash()
        skip flag parameters -> flags.GenerateBash() / collect positionals
    end
    execute collected AddIf -> generate actual completions
    default COMPREPLY
}
register completion function

execute AddFunc
```

Note that `AddFunc` and `AddIf` are collected when calling `GenerateBash()`.
They are used to be able to generate/inject code to another place in the
template.
