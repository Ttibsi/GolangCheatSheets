# Golang

Developed by google in 2007, open source in 2009

Go was designed to run on multiple cores and utilise concurrency
It's primarily used for applications that are performant, modern, scaled/distributed

Go is meant to combine the simplicity of Python and the speed/efficiency of C++.
It's used on serverside/backend. 
    - Microservices
    - Web applications
    - database services

Examples: Docker, k8s, Hashicorp Vault

Simple syntax, fast to startup/run, requires fewer resources.
It's a compiled language that is consistent across all OSes

## Setting up your environment
Go can be installed in multiple ways. I've installed it using `brew install go`

This can be tested by just running `go` in your terminal, or even `go version`
to see the installed version.

### Neovim
- Run this command to install `gopls` language server: `go install golang.org/x/tools/gopls@latest`
    - src: https://github.com/golang/tools/tree/master/gopls#installation
- Make sure to add `go` to treesitter and the below snippet to nvim-lsp
    - Run `TSInstall go` to add go to treesitter.
    - Run `LspInfo` in a `.go` file to test the lspconfig is working
- Add the following lines to your environment config (bashrc, zshrc) to create relevant envvars 
    - `export PATH=$PATH:/usr/local/go/bin`
    - `export GOPATH=$HOME/.go`
    - `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`
    (Note that I instead decided to rename `~/go` to `~/.go`)

```lua
-- lsp-config language server setup
require'lspconfig'.gopls.setup({
    on_attach = on_attach,
    capabilities = capabilities
})
```

## Syntax

Run `go mod init <name>` in the terminal at the top level of your repo. This creates the
module with a name that should be the same as your github repo name. 
    - This generates the `go.mod` file

Standard hello world code:
```go
// Create a package to work in - this by default should be 'main'
package main

import "fmt" // standard library

// Tell GO where the entry point to the code is - this is always a main function
func main() {
    fmt.Println("Hello World")
}
```
To execute a project, run `go run <file>`

```go
func main() {
    // define a variable
    var name = "foo"
    const PI = 3.14 // constants

    //concatenation
    // Concatenation has an implicit space entered between values
    fmt.Println("Hello user:", name)

    // formatting data
    fmt.Printf("Welcome user: %v\n", name)
    // There are a number of other placeholders to use instead of %v for different
    // purposes
}
```

### Data types
Go supports a number of different data types. You can use the `var` keyword to
create a variable with an unknown data type

```go
func main() {
    // variable type inferred by assignment
    var name = "foo"

    // variable type assigned by the programmer
    var userName string
    var age int

    // print the type of the variables
    fmt.Printf("userName type is %T\n", userName)

    // syntactic sugar to define a variable
    welcome := "hello hello"
    // This has to use assingment but it can infer the type
}
```

Integers can have different lengths (`int8`, `int16` etc) that represent the
number of bytes or use `uint8, uint16` which are _unsigned_ integers, that are
always positive, and I think this also saves on some memory

```go
func main() {
    var count uint = 180
    var otherCount int8 = 127

    count = count - uint(otherCount) // This converts it to a uint because this
    // is a type mismatch. Or change the instanciation
```

Convert from one format to another using the `strconv` library
```go
import (
    "fmt"
    "strconv"
)

func main() {
    var a string = "1"
    var b int = strconv.FormatUint(uint64(a), 10) // 10 == base 10 
}
```

### User input and Pointers

```go
func main() {
    // get user input
    var userName string
    fmt.Scan(&userName) //This assigns the user input to the variable userName
}
```
Pointers point from a variable to the location of the data in memory.
Running `fmt.Println(&var)` with the ampersand will print the memory address


### Arrays and slices

```go
func main() {
    // Array instanciation
    // You need to specify the length (in this case 10) and the data type
    var my_friends = [10]string{} // assign this an empty list with {} 

    // An array with some starting values
    var my_friends = [10]string{"Tom", "Dick", "Harry"} 

    // array creation without any assignment
    var my_friends = [10]string
    my_friends[0] = "Tom"
    my_friends[1] = "Dick"
    my_friends[2] = "Harry"

    fmt.Println(my_friends)
    my_friends[0] = "Lily"
    fmt.Print(my_friends[0])

    // Array of unknown length
    // These are called slices instead and are more memory efficient
    my_friends = append(my_friends, "Rosanna")
}
```

## Loops

```go
import (
    "fmt"
    "strings"
)

func main() {
    -- for loops are the only loops in go
    for idx, elem := range friends { // using range is like python's enumerate()
        // strings.Fields splits on whitespace
        var first_name = strings.Fields(elem) //strings needs importing
        fmt.Println(idx, first_name)
    }
    
    -- simple loop adding numbers
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

You can also use the `continue` and `break` statements inside `if` statements
the same as you would in python. This allows for logic breaking in loops

## Conditionals

```go
func main() {
    var alive bool = true
    if alive {
        fmt.Println("Alive")
    } else if !alive {
        fmt.Println("He dead, jim")
    } else {
        fmt.Println("Houston, we have a problem")
    }
}
```

Instead of IF statements, you can use a switch statement

```go
func main() {
    city := "London"
    switch city {
        case "New York":
            fmt.Println("New York")
        case " Vancouver":
            fmt.Println("Vancouver")
        case "Hong Kong":
            fmt.Println("Hong Kong")
        default:
            fmt.Println("Other")
    }
}
```
You can also use statements that an IF would check for in a `for` statement, and
when they become false, the loop will break. This is what's happening at lines
173-176

## Input validation
This is making sure that the entries from a `fmt.Scan` are correct/valid data
for the continuation of the code

`isValidEmail := len(newName) >= 2 && strings.Contains(newName, "@")`

- `&&` - AND
- `||` - OR
- `!` - NOT

You can also check if a string contains certain text using `strings.Contains`
or check the length with `len()`. 

This should usually be combined with an IF statement and some output to inform
the user.

## Functions

Each go code starts with the `main` function. This is the entry point.
You can create a new function under the previous one, as long as it follows the
standard formatting

```go
func main() {
    var name string = "Lily"
    fmt.Println(welcome_user(name))
}
func welcome_user(user string) string {
    ret := "Welcome:" + user
    return ret
}
```

You also need to specify the return type:
```go
func function_name(param_name param_type) return_type {
}
```

## Package level variables
These are global variables that are accessible in multiple Functions
They use the standard variable instanciation syntax, but are written at the top
of your code file, under your import statements, outside of any syntax block

However, they have to use the `var` keyword and can't use the `:=` method of
creating variables. 

Accessing them across the functions requires nothing special.

## Working with multiple files
Go programs are organised into packages. A package is a collection of `.go` files

Each file needs to have the same `package <name>` line at the top of the file
to access each other. This replaces python's import statements. You will need
to make sure that each file has all the libraries imported that they need.

However you'll need to make sure that when you do `go run`, you name every file
that you require.
    - ex: `go run main.go helper.go util.go`
    - As an alternative, you an just do `go run .` to just run all files in the 
    current folder

### Multiple Packages
You can also organise your various files into multiple packages in the same
repo. This is mostly used for when you want different 'folders' of code 
within the same repo. Different packages have different namespaces, so you'll 
need to import functions from other packages within the `import` statement at 
the top of the repo

```go
import (
    "fmt"
    "module_name/helper" // module_name is specified in go.mod 
)
```

Functions for use in other packages need to be exported. At the function
definition, you just need to make sure the first letter of the function name is 
capitalised. You can then call it with `package.Function` 

This all also applies to variables too

### Scope rules
- local - variable declared within a function/ block of code, and can only be used
within that scope
- package level - these are defined outside a function and can be called
throughout a file
- global level - these are capitalised and accessible throughout multiple
packages

## Maps
These are hash maps, or called dictionaries in python. 

```go
func main() {
    var userData = make(map[string]int) 
    userData["grade_bio"] = 1
    userData["grade_english"] = 2
    userData["grade_geo"] = 3

    fmt.Print(userData)
    var list_of_maps = make([]map[string]int, 0) // the number is the initial length
}
```

## Structs
From what I can tell, this is basically a Class but with only attributes, no 
methods. These are defined outside of a function, at the top level, with this 
syntax:

```go
type personData struct {
    firstName string
    lastName string
    age int
    is_allowed bool
}
```
You can access the individual values the same way you do elsewhere, doing 
something like `personData.firstName`

## concurrency

You need to use `sprintf` to save a formatted string to a variable
` var baz = fmt.Sprintf("%v ages", 5)`

```go
func main() {
    -- This blocks the execution of a thread for the specified duration
    time.Sleep(10 * time.Second)
}
```

You can easily make functions concurrent within Go. This is multithreaded
execution. To make a function call act in it's own thread, add the `go` keyword
befor you call it. IE something like `go addNums(1,2)`

If we need to make the main code wait for a thread to finish, we can use the
`waitgroup` library to create some waitgroups for Go to recognise. Before we 
create a thread with a `go` line, we use `wg.Add(1)` to add the thread to a
stack. We can also use `wg.Done()` to remove from the stack, and add `wg.Wait()`
at the very end of the code to tell it that you can't go past that point for 
the threads

Threads are also called GoRoutines
GoRoutines use "Channels" to share data between threads.

