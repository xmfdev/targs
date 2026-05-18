# targs - terminal arguments

A minimal CLI argument parser written in Go.

## Installation

```
go get github.com/xmfdev/targs
```

## Usage

Define your options, then call `HandleArgs` with `os.Args`:

```
package main

import (
	"fmt"
	"os"
	"github.com/xmfdev/targs"
)

func main() {
	var name string

	options := []targs.Option{
        {
			Options: []string{"-v", "--verbose"},
			Handler: func(_ *string) { fmt.Println("verbose mode") },
		},
		{
			Options:     []string{"-n", "--name"},
			HasExtraArg: true,
			Handler:     func(s *string) { name = *s },
		},
	}

	targs.HandleArgs(os.Args, options, func() {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	})

	fmt.Println("hello,", name)
}
```

```
$ go run . -n Alice
Hello, Alice!

$ go run . --verbose
verbose mode enabled

$ go run . --port=8080
port 8080
```

## API

### `Option`

| Field        | Type            | Description                                                    |
|--------------|-----------------|----------------------------------------------------------------|
| `Options`    | `[]string`      | Flag names that trigger this option (e.g. `-v`, `--verbose`)   |
| `Handler`    | `func(*string)` | Called when a flag is matched; receives the extra arg or `nil` |
| `HasExtraArg`| `bool`          | If true, the next argument is consumed and passed to `Handler` |

### `HandleArgs(args, options, invalidArg)`

Iterates over `args` (pass `os.Args`) and matches each element against the provided options.

- Each flag may appear at most once; duplicates trigger `invalidArg`.
- Unknown flags trigger `invalidArg`.
- If `HasExtraArg` is true and no following argument exists, processing stops silently.
- `invalidArg` may be `nil` to ignore errors.

## License

Licensed under the MIT license. See `LICENSE` for more information.
