# vim9gorn

vim9gorn is a Go library for generating Vim9Script configurations and plugins programmatically.

## Features

- Define Vim **options** (`set number`, `tabstop=4`)
- Define **variables** with arbitrary scope (`g:`, `b:`, `w:`, `t:`, `s:`, `var`, `const`)
- Configure **colorschemes** and **syntax highlighting**
- Define **keymaps** in normal, visual, etc.
- Create **functions** with arguments, return types, and nested code blocks
- Inject **raw Vim9Script** for advanced customization
- Fluent API for programmatic Vim configuration

## Installation

```bash
go get github.com/greeschenko/vim9gorn
