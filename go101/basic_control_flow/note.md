## 1. Control flow

- if-else

```go
if InitSimpleStatement; Condition {
	// do something
} else {
	// do something
}
```

- if block and else block is nested in implicit code block
  this explain why the two `n` variable do not conflict!

```go
if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "is an even number.")
	} else {
		fmt.Println(n, "is an odd number.")
	}

	n := rand.Int() % 2 // this n is not the above n.
	if n % 2 == 0 {
		fmt.Println("An even number.")
	}

	if ; n % 2 != 0 {
		fmt.Println("An odd number.")
	}
```

- and there

```go
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("Now is AM time.")
	} else if h > 19 {
		fmt.Println("Now is evening time.")
	} else {
		fmt.Println("Now is afternoon time.")
		h := h // the right one is declared above
		// The just new declared "h" variable
		// shadows the above same-name one.
		_ = h
	}

	// h is not visible here.
```

### > Can break control flow

- for

```go
for InitSimpleStatement; Condition; PostSimpleStatement {
	// do something
}
```

- for-range
- switch-case

```go
switch InitSimpleStatement; CompareOperand0 {
case CompareOperandList1:
	// do something
case CompareOperandList2:
	// do something
...
case CompareOperandListN:
	// do something
default:
	// do something
}
```

- `CompareOperand0` will make a implicit convert to the default type

- Do not use `break` to break `switch-case` control flow

```go
switch n := rand.Intn(100); n%9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")

		// Different from many other languages,
		// in Go, the execution will automatically
		// jumps out of the switch-case block at
		// the end of each branch block.
		// No "break" statement is needed here.
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		// Here, this "break" statement is nonsense.
		break
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// case 6, 7, 8:
		// The above case line might fail to compile,
		// for 6 is duplicate with the 6 in the last
		// case. The behavior is compiler dependent.
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
```

- In the order hand, use `fallthrough` to slip to the next code block, and `fallthrough` must be the last statement (except the last case code block)

```go
switch n := rand.Intn(100) % 5; n {
case 0, 1, 2, 3, 4:
	fmt.Println("n =", n)
	// The if-block, not the fallthrough statement,
	// is the final statement in this branch.
	if true {
		fallthrough // error: not the final statement
	}
case 5, 6, 7, 8:
	n := 99
	fallthrough // error: not the final statement
	_ = n
default:
	fmt.Println(n)
	fallthrough // error: show up in the final branch
}
```

- goto
- Note that, if a label is declared within the scope of a variable, then the uses of the label can't appear before the declaration of the variable.

```go
package main

import "fmt"

func main() {
	i := 0
Next:
	if i >= 5 {
		// error: jumps over declaration of k
		goto Exit
	}

	k := i + i
	fmt.Println(k)
	i++
	goto Next

// This label is declared in the scope of k,
// but its use is outside of the scope of k.
Exit:
}
```

- 2 ways to fix

```go
// 1
func main() {
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}
	// Create an explicit code block to
	// shrink the scope of k.
	{
		k := i + i
		fmt.Println(k)
	}
	i++
	goto Next
Exit:
}
// 2

func main() {
	var k int // move the declaration of k here.
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}

	k = i + i
	fmt.Println(k)
	i++
	goto Next
Exit:
}

```

- `break` and `continue` with labels

```go
Outer:
	for n++; ; n++{
		for i := 2; ; i++ {
			switch {
			case i * i > n:
				break Outer
			case n % i == 0:
				continue Outer
			}
		}
	}
	return n
}
```

- select-case
- type-switch
