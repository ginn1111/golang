### Container value literal
1. Slice: []T{}
2. Map: map[K]T{}
3. Array: [...]T{}
> T is a type and K is a comparable type
> The length of array must be a constant type

### Comparable container
 - Map and slice is incomparable value -> do not use as a map key type.
 - Array can compare each other with corresponding order.

 ### Retrieve and modify container
  - The index of array or slice must be:
    - The index must be the `int` base type non-negative.
    ```go
        [...]bool{false, true, 3:true, false}
        //        ^       ^             ^this index is 4
        //        ^       ^ this index is 1
        //        ^ this index is 0
        [...]bool{3: false, 1: true, true}
        //                            ^ this index is 2
    ```
  - If map is nil map -> retrieve will return zero value of element type. If this operator use as assignment, will cause panic.
  - In the other hand, array or slice is a nil value, retrieve will cause a panic.

### Container assignment
  - When assign a container to another container -> this will copy a direct part, so the underlying elements of each can share with other.
  - Despite the addressable one do not equal (in difference place), but the direct part just point one the underlying part.
  ```go
    map1 := map[string]int{"test": 1}
    map2 := map1
    // print the address of start of underlying part -> the same
    fmt.Printf("%p %p \n", map1, map2)
    // print the address of direct part -> difference
    fmt.Printf("%p %p \n", &map1, &map2)
    fmt.Println(&map1 == &map2) // print false

  ```
  - The above example the same as slice type.
  - When assignment the array to another one, the direct part will copy the other place, so the elements do not share with each other.

### Append and delete container
  1. Delete map `var m map[k]T`
  syntax: `delete(m, k)`
  > Do not delete `NaN` key
  2. Append slice
  - If the capacity of slice is enough -> will append to the base slice. The result slice wil share the first elements with the base slice.
  ```go
    slice := make([]int, 3, 5)
    newSlice := append(slice, 1, 2) // share the three first elements with slice 

    newSlice2 := append(slice, 1, 2, 3) // the capacity do not enough -> create new slice 
                                        // do not share any elements with slice
  ```

### Addressability of container element 
 - Elements of addressability of array values also addressable and vice versa.
 - Elements of any slice always addressability. Because elements of slice in underlying part. 
 - Element of map always un-addressable. Because the issue of perform Go runtime when assignment addressable element to another variable.
 ```go
  map_ := map[string]int{"age": 24}
  address := &map_["age"] // will cause compiler error

  // array
  fmt.Println(&[...]int{1,2,3}[0]) // will cause compiler error, because composite literal always un-addressable 

  // slice
  fmt.Println(&[]int{1,2,3}[0]) // ok 

 ```

### Derive slices from arrays and slices
  - syntax: `baseContainer[low : high]`
    or `baseContainer[low : high : max]`
    > 0 <= low <= high(exclude) <= max(exclude)
  - Derive slice do not the new slice (memory viewed), it will shareable with base slice
  ```go
    arr := [...]int{1,2,3}
    deriveSlice := arr[:1] // equipment arr[0:1:cap(arr)]

    fmt.Println(deriveSlice, len(deriveSlice), cap(deriveSlcie)) //result [1] with len 1 and cap 3
  ```
### Convert Slice to array pointer
  - The underlying part of slice will share with array.
  ```go 
    slice := []int{1,2,3}
    arr := (*[3]int)(slice)

    slice[0] = 4
    fmt.Println(arr) // print [4 2 3]
  ```
### Convert Slice to array
  - The underlying part of slice **DO NOT** share with array.
  ```go 
    slice := []int{1,2,3}
    arr := ([3]int)(slice)

    slice[0] = 4
    fmt.Println(arr) // print [1 2 3]
  ```
### Copy Slice elements with the built-in `copy` function
  - syntax: `numOfEls := copy(dest, source)`
  - This function do not create new slice (memory view), so the underlying elements share with each other
  ```go
    slice := []int{1,2,3,4}
    num := copy(slice[:1], slice)
    fmt.Println(num, slice) // print 1, [1 2 3 4]

    num2 := copy(slice[1:], slice)
    fmt.Println(num2, slice) // print 3, [1 1 2 3]

  ```
### Container elements iterations
 - syntax: 
  ```go
    for key, element = range aContainer {
      // use key and element ...
    }
  ```
  - `key` and `element` are called `iteration variables`.
  - Can over the **nil** slice or map

  - `element` variable will copy the value of container iteration, so if container is array, modify array elements do not reflect to the `element variable`.
    - In the other hand, if the container is slice or map, modify its elements will reflect to the `element variable` .
  ```go
  type Person struct {
		name string
		age  int
	}
	persons := [2]Person {{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)

		// This modification has no effects on
		// the iteration, for the ranged array
		// is a copy of the persons array.
		persons[1].name = "Jack"

		// This modification has not effects on
		// the persons array, for p is just a
		// copy of a copy of one persons element.
		p.age = 31
	}
	fmt.Println("persons:", &persons)
  /*
  * 0 {Alice 28}
  * 1 {Bob 25}
  * persons: &[{Alice 28} {Jack 25}]
  */
  ```
  ```go
  // A slice.
	persons := []Person {{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)

		// Now this modification has effects
		// on the iteration.
		persons[1].name = "Jack"

		// This modification still has not
		// any real effects.
		p.age = 31
	}
	fmt.Println("persons:", &persons)
  /*
   * 0 {Alice 28}
   * 1 {Jack 25}
   * persons: &[{Alice 28} {Jack 25}]
  */
  ```
 - When use `for-range` with array, should use `array pointer` or `slice ([::])` to make effective to the memory, because the `for-range` will perform the copy operator.

### The `memclr` optimization
- The one-iteration-variable `for-range` :
    ```go
    // ex var a []T with t0 is zero value of type T
      for i := range a {
        a[i] = t0
      }
    ```
    will be Go compiler perform the optimization
  > Using it when we want to flush the memory

### The `clear` built-in function to clear map and reset slice
  - This func will clear the key `NaN` in the map
  ```go
    s := []int{1,2,3}
    clear(s)
    fmt.Println(s) // print [0 0 0]

    a := [4]int{5, 6, 7, 8}
    clear(a[1:3])
    fmt.Println(a) // [5 0 0 8] 

  ```
### When the functions `cap` and `len` will call in array or array pointer, this will evaluated in the compile time -> its type is `constant int type`
```go

	arr := [...]int{1,2,3}
	arrPointer := &[...]int{1,2,3}
	const lenArrPointer = len(arrPointer) // hover it
	const lenArr = len(arr) // hover it
```
### More slice manipulations
#### Clone slice
```go
  s := []int{1,2,3}

  sClone1 := append(s[:0:0], s...)
  sClone2 := append([]int(nil), s...) // result is nil slice if s is blank none-nil slice
  sClone3 := make([]int, len(s))
  copy(sClone3, s) // -> the best way for now
  sClone4 := append(make([]int, 0, len(s)), s...) // slower than sClone3

```

#### Delete segment slice
```go
  s := []int{1,2,3}

  // perverse order
  /*
    ex: slice is: 1, 2, 3(from), 4, 5, 6(to), 7, 8, 9
    from: 2, count: 3, to: from + count(5)
    result:  1,2,6,7,8,9
  */
  //
  from := 2
  count := 3
  to := from + count

  s := append(s[:from], s[count:]...)

  // copy
  num := copy(s[:from], s[:count])

  // slice
  s := s[:from+num]
  // or one line
  s := s[:from + copy(s[:from], s[:count])]

  // Don't preserve order
  /*
    ex: slice is: 1, 2, 3(from), 4, 5(to), 6, 7, 8, 9
    from: 2, count: 2, to: from + count(4)
    result:  1,2,7,8,9,6(do not original order)
  */
  if n := to - from; len(s) - to < n {
    copy(s[from:to], s[to:])
  } else {
    copy(s[from:to], s[len(s)-n:]) 
  }

  s := s[:len(s)-(to-from)]

  // just use when element type of array is pointer
  // flush memory
  for i := range s[len(s)+to-from] {
    s[i] = 0 // zero value
  }
```
#### Delete one element
```go
  s := append(s[:from], s[from+1:])
  s := append(s[:from+copy(s[from:], s[from+1:])])

  copy(s[from:from+1], s[len(s)-1:]) // or s[from] = s[len(s)-1]
  s := [:len(s)-1]

  // just use when element type of array is pointer
  s[len(s)-1] = 0 //zero value
```

#### Insert elements
```go
  elements := []int{3,3,3}
  s = append(s[:i], append(elements, s[i:]...)...)

  // the just small effective
  if cap(s) >= len(elements) + len(s) {
    s = s[:len(s) + len(elements)]
    copy(s[i+len(elements):], s[i:])
    copy(s[i:], elements)
  } else {
    s := make([]int, 0, len(elements) + len(s))
    append(s, s[:i]...)
    append(s, elements...)
    append(s, s[i:]...)
  }

```
### Use maps to simulate set
  - We use the map `map[K]struct{}` to save the value of set is `K`
  - The size of `struct{}` is zero, so this doesn't occupy memory space. 





