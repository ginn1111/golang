### Ex1
```go
  m := map[string]int{"Go": 0}
	s := []int{1, 1, 1}; olds := s
	n := 2
	p := &n
	s, m["Go"], *p, s[n] = []int{2, 2, 2}, s[1], m["Go"], 5
  fmt.Println(s, m, p) // [2 2 2], {"Go": 1}, 0
  fmt.Println(olds) // [1 1 5]
/*
  evaluation phase
    P0 = &s; P1 = &m; P2 = p = &n; P3 = &s; K0="Go"; K1=2
    R0=[]int{2,2,2}(literal slice); R1=1; R2=0; R3=5

  assign phase
    *P0=R0
    (*P1)[K0]=R1
    *P2=R2
    (*P3)[K1]=R3
*/
```
### Ex2
```go
  x := []int{2, 3, 5, 7, 11}
	t := x[0] // 2
	var i int = 2
	for i, x[i] = range x { // 2, 3, 5, 7, 11
  /*
    iter 0: i, x[2] = 0 , 2
    now i = 0
    2 3 2 7 11

    iter 1: i, x[0] = 1, 3
    now i = 1
    3 3 2 7 11

    iter 1: i, x[1] = 2, 2
    now i = 1
    3 2 2 7 11
    ...
    3 2 7 7 11
    3 2 7 11 11
  
  */
	}

	fmt.Println(x, i)
	var _ = t
	// x[i] = t
	// fmt.Println(x, i) // [3 5 7 11 2]
```