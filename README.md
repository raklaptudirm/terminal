# ansi - Easy ANSI Escape Sequences

`ansi` implements functions to use various useful ANSI Escape Sequences for
doing sophisticated terminal manipulation, including cursor movement and erasing.

### Installation

```
go get -u laptudirm.com/x/ansi
```

### Examples

```go
import (
        "os"
        "laptudirm.com/x/ansi"
)

// create an *ansi.Terminal from os.Stdout
terminal := ansi.NewTerminal(os.Stdout)

// various terminal manipulation functions
terminal.EraseScreen()
terminal.HideCursor()
terminal.MoveCursorHome()

// all the fmt functions are also defined
terminal.Print("Hello, ")
terminal.Println("World!")
terminal.Printf("PI: %d", 3.1415)
```

### Documentation

The documentation and a comprehensive list of all the manipulation functions
can be found at https://laptudirm.com/x/ansi.

### References

- https://en.wikipedia.org/wiki/ANSI_escape_code
- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
