# terminal - Simple Terminal manipulation methods

`terminal` implements various functions for doing sophisticated terminal
manipulation, including cursor movement, screen erasing, and others.

### Installation

```
go get -u laptudirm.com/x/terminal
```

### Examples

```go
import (
        "os"
        "laptudirm.com/x/terminal"
)

// create an *terminal.Terminal from os.Stdout
term := terminal.NewTerminal(os.Stdout)

// various terminal manipulation functions
term.EraseScreen()
term.HideCursor()
term.MoveCursorHome()

// all the fmt functions are also defined
term.Print("Hello, ")
term.Println("World!")
term.Printf("PI: %d", 3.1415)
```

### Documentation

The documentation and a comprehensive list of all the manipulation functions
can be found at https://laptudirm.com/x/terminal.

### References

- https://en.wikipedia.org/wiki/ANSI_escape_code
- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
