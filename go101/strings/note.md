### Some simple facts about string
  - string elements do not addressable taken -> do not modify.
  - The internal struct of string like that:
  ```go
    type string struct {
      elements *byte
      len int
    }
  ```
  - So when copy a string to another one, the both will share underlying elements with each other.
  - To `slice` string, we use `[:] operator` like slice or array container.
  - `string` is comparable, compare one by one, if the one string has the length greater than one, this string will larger than one.
  - `string[idx]` do not addressable -> do not modify.
  - `string` and `indexes` are constant, but `string[idx]` or `string[start:end]` do not a constant
  ```go
    const s = "go101.org"

    var a byte = 1 << len(s) / 128 // 4 -> evaluate in compile time, because this is constant expression
    var b byte = 1 << len(s[:]) / 128 // 0 -> evaluate in runtime
  ```

### String related conversion
- `string` can convert explicit to `[]byte` or `[]rune`. 
- When convert to `[]byte` each char is present one byte, if UTF-8 do not invalid, will remain.
- Otherwise, bad UTF-8 will present `0xFFFD` when convert from `string` to `[]rune`.

### Compiler optimization for convert between string and byte slice
- When convert `string` to `[]byte`, this will express a **deep** copy from underlying element to another one -> ineffective in memory.
- Compiler will perform a optimization (do not copy) in some operations.
  - In `for-range` loop -> convert to `[]byte` do not perform a copy.
  ```go
    for i, v := range ([]byte)(str) {
      // do something
    }
  ```
  - In map key assignment and map key indexing.
  ```go
    slice := []byte{"g", "o"}
    map_ := map[string]string
    map_[string(slice)] = "value" // this will perform a copy

    fmt.Println(map_(string(slice))) // do not perform a copy
  ```
  - In comparison, Go compiler do not perform a copy.
  - When concatenation, Go compiler will perform a copy, but when concatenation with blank string `" "` this will do not perform
  ```go
    s1 := []byte{1023: 'x'}
    s2 := []byte{1023: 'y'}

    if string(s1) != string(s2) { // do not perform a copy
      s := (" ") + string(s1) + string(s2) // do not perform as well
    } else {
      s := string(s1) + string(s2) // perform a copy
    }
  ```
### Range for string
 - `for-range` iterate with each `run` UTF-8 code point.
 - `len(s)` if s is slice -> O(1), s is rune -> O(n)
 - compare two string: if the pointers of element are equal -> O(1), otherwise O(n)
### Using string as byte slice (append, copy)