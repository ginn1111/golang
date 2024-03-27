# Code Blocks and Identifier Scopes

## Code blocks

1. **Universe block** contain all project source code.
2. **Package block** containing all source code, excluding the package import declarations in that package.
3. **File block** containing all source code in that file, including the package import declarations in that file.
4. **Local block** encloses a `{}`.

- `if`, `switch`,`for` keyword is followed by two nested local blocks. One is implicit and the other is explicit.

## Scopes

- `Type` define can use itself, counterpart with define a `value`

- Declarations in the package-level block (top local block) can reorder declaration depend on the dependent of these variables

## Shadowing

- The variable of innermost will shadowing the outermost.

```go

var x = 999 //x1

func main() {
  var x = 888 //x2

  for x:= 70; x < 100; x++ { //x3
    x := x - 70 // x4
  //^    ^ the old x (70)
  //^ the new x variable
    if x:= 10; x < 20 { //x5
      x := x - 10 // x6
    }

  }
}

```

> [!NOTE]
> Take care of short variable declarations

```go
n, err := strconv.Atoi(s)
	if err != nil {
		// Some new gophers may think err is an
		// already declared variable in the following
		// short variable declaration. However, both
		// b and err are new declared here actually.
		// The new declared err variable shadows the
		// err variable declared above.
		b, err := strconv.ParseBool(s)
		if err != nil {
			return 0, err
		}

		// If execution goes here, some new gophers
		// might expect a nil error will be returned.
		// But in fact, the outer non-nil error will
		// be returned instead, for the scope of the
		// inner err variable ends at the end of the
		// outer if-clause.
		if b {
			n = 1
		}
	}
	return n, err

```
