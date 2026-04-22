# vim9gorn

[![CI](https://github.com/greeschenko/vim9gorn/actions/workflows/ci.yml/badge.svg)](https://github.com/greeschenko/vim9gorn/actions/workflows/ci.yml)

vim9gorn is a Go library for generating Vim9Script configurations and plugins programmatically.

## Features

### Core
- **Options** — `set number`, `tabstop=4`, etc.
- **Variables** — `g:`, `b:`, `var`, `const`
- **Functions** — `def` with args, return types, nested blocks
- **Control Flow** — `if/elseif/else`, `for`, `while`, `try/catch`
- **Collections** — `List`, `Dict` literals
- **Keymaps** — `nmap`, `imap`, `vmap`, etc.
- **Highlights** — group links, custom highlights
- **Colorschemes** — background, termguicolors

### Advanced
- **Classes** — `class` with fields, methods, inheritance
- **Lambda** — `(x) => x + 1` expressions
- **Range** — `range(1, 10)`, `range(0, 100, 5)`
- **Error Handling** — `throw`, `assert`, `assert_equal`, custom error types

### Plugin Development
- **Commands** — user-defined commands
- **Autocmd** — autocommand groups
- **Plugin Structure** — scaffold `plugin/`, `autoload/`, `ftplugin/` directories
- **Autoload** — library functions
- **Plugin Manager** — fetch/update external plugins (git clone)

### Utilities
- **Fluent API** — chainable method calls

## Requirements

- Go 1.24+
- Git (for plugin manager)
- Vim 9+ (for validation tests)

## Installation

```bash
go get github.com/greeschenko/vim9gorn
```

## Quick Start

```go
package main

import (
    vi "github.com/greeschenko/vim9gorn"
)

func main() {
    g := vi.New()

    // Options
    g.AddSection(
        vi.NewOptions().
            Set("number", true).
            Set("tabstop", 4))

    // Variables
    g.AddSection(
        vi.NewVariables().
            Legacy(vi.Global, "mapleader", `"\<Space>"`))

    // Colorscheme
    g.AddSection(vi.ColorScheme{
        Background:   "dark",
        TermGuiColors: true,
        SyntaxEnable: true,
        Name:        "cyberpunk99",
    })

    // Keymaps
    g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>q", RHS: "ZZ"})

    // Functions
    greet := vi.NewFunction("Greet").
        SetScope(vi.Global).
        Add(vi.Raw{Code: `echo "Hello from vim9gorn!"`})

    g.AddSection(greet)

    // Write to file
    writer := &vi.DefaultFileWriter{}
    g.Forge(".vimrc", writer)
}
```

## API Reference

### Options

```go
vi.NewOptions().
    Set("number", true).
    Set("tabstop", 4).
    Set("background", "dark")
```

### Variables

```go
// Script-local variable
vi.NewVariables().Var("name", `"value"`)

// With type
vi.NewVariables().VarTyped("count", "number", "0")

// Constant
vi.NewVariables().Const("PI", "3.14")

// Legacy scope (g:, b:, w:, t:)
vi.NewVariables().Legacy(vi.Global, "mapleader", `"<Space>"`)
```

### Functions

```go
vi.NewFunction("Add").
    Arg("a", "number").
    Arg("b", "number").
    Returns("number").
    Add(vi.NewReturn("a + b"))
```

### Classes

```go
c := vi.NewClass("Person").
    AddField("name", "string").
    AddFieldWithDefault("age", "number", "0").
    AddMethod(
        vi.NewFunction("GetAge").
            Returns("number").
            Add(vi.NewReturn("this.age")))

// Usage:
// class Person
//     this.name: string
//     this.age: number = 0
//     def GetAge(): number
//       return this.age
//     enddef
// endclass
```

### Lambda

```go
// Simple lambda
l := vi.NewLambda("x + 1").Arg("x")

// Lambda with multiple args and return type
l := vi.NewLambda("x + y").
    Arg("x: number").
    Arg("y: number").
    SetReturn("number")

// Usage: (x: number, y: number): number => x + y
```

### Range

```go
// Simple range
r := vi.NewRange(1, 10)
// Usage: range(1, 10)

// Range with step
r := vi.NewRangeWithStep(0, 100, 5)
// Usage: range(0, 100, 5)
```

### Control Flow

```go
// If/Else
vi.NewIfElse("cond").
    ThenAdd(vi.Raw{Code: "echo 'yes'"}).
    ElseAdd(vi.Raw{Code: "echo 'no'"})

// For loop
vi.NewForLoop("k", "v", "items(dict)").
    Add(vi.Raw{Code: "echo k"})

// While loop
vi.NewWhileLoop("i < 10").
    Add(vi.Raw{Code: "let i += 1"})

// Try/Catch
vi.NewTryCatch().
    AddTry(vi.Raw{Code: "risky_call()"}).
    SetCatch("e", "Vim:.*").
    AddCatch(vi.Raw{Code: "echo 'Error: ' .. e"})
```

### Error Handling

```go
// Throw
vi.NewThrow(`"something went wrong"`)

// Assert
vi.NewAssert("x > 0").SetMsg("x must be positive")

// Assert equal
vi.NewAssertEqual("actual", "expected")

// Assert exception
vi.NewAssertException("may_fail()").SetError("E1234")

// Custom error type
vi.NewErrorType("MyError", "E1001")
```

### Keymaps

```go
vi.Keymap{Mode: "n", LHS: "<leader>q", RHS: "ZZ"}
vi.Keymap{Mode: "n", LHS: "<C-w>h", RHS: "<C-w>h"}
vi.Keymap{Mode: "v", LHS: "y", RHS: `"y`}
```

### Command

```go
&vi.Command{Name: "PluginsInstall", Cmd: "call.SetupPlugins()"}
```

### Autocmd

```go
ag := vi.NewAutocmdGroup("MySettings").
    Add(*vi.NewAutocmd("FileType", "go").
        SetCmd("setlocal shiftwidth=4"))

g.AddSection(ag)
```

### Plugin Structure

```go
// Create plugin scaffold
manifest := vi.NewManifest().
    Add(vi.Plugin{Name: "myplugin", Type: vi.PluginTypeOpt})

// Create directories
manifest.CreateDirectories(".")
```

### Autoload

```go
af := vi.NewAutoloadFile("mylib").
    AddFunc(*vi.NewAutoloadFunc("mylib", "Hello").
        Add(vi.Raw{Code: "echo 'hello'"}))
```

### Plugin Manager

```go
manager := vi.NewPluginManager().
    Add("tpope/vim-fugitive").
    Add("yegappan/lsp")

// Fetch all plugins
manager.FetchAll()

// Update all plugins
manager.UpdateAll()
```

### Lambda Helpers (Filter/Map)

```go
// Filter a list
pred := vi.NewLambda("x > 0").Arg("x")
result := vi.Filter("[1, -1, 2, -2]", pred)
// Usage: filter([1, -1, 2, -2], (x) => x > 0)

// Map a list
fn := vi.NewLambda("x * 2").Arg("x")
result := vi.Map("[1, 2, 3]", fn)
// Usage: map([1, 2, 3], (x) => x * 2)
```

## Examples

See `examples/` directory:

- `basic/` — Basic usage example
- `dotfiles/` — Full .vimrc generator
- `plugin/` — Plugin template generator

## Testing

```bash
go test ./...
go test -v -run TestGenerateUserVimrc  # Run integration test
go test -v -run TestValidateGeneratedVimrcWithVim  # Validate with Vim
```

> **Note:** Vim must be installed for the validation test (`sudo apt-get install vim`)

## Project Structure

```
.
├── gorn.go           # Core Gorn struct
├── options.go         # Options generation
├── variables.go      # Variables generation
├── functions.go      # Function generation
├── return.go         # Return statement
├── class.go          # Class definition
├── lambda.go         # Lambda expressions
├── ifelse.go        # If/elseif/else
├── loop.go          # For/while loops, range
├── trycatch.go      # Try/catch/finally
├── errors.go        # Error handling
├── collections.go  # List/Dict
├── keymap.go        # Keymaps
├── highlight.go     # Highlights
├── colorscheme.go  # Colorschemes
├── command.go       # User commands
├── autocmd.go       # Autocommands
├── filetype.go      # Filetype detection
├── plugin.go       # Plugin scaffolding
├── autoload.go      # Autoload functions
├── manager.go      # Plugin manager (git fetch)
├── raw.go          # Raw Vim9 code injection
├── examples/
│   ├── basic/      # Basic example
│   ├── dotfiles/   # .vimrc generator
│   └── plugin/     # Plugin template
└── testdata/
    └── output.vimrc  # Integration test output
```

## License

MIT