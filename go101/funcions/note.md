## Function

### More About Function Declarations and Calls
  - The name function do not be duplicate (except the `init` or blank `_` (the function with blank name do not call anymore)) 
  - The return results of custom function can be discard, but do not true with some built-in function like `recover` or `copy`.
  - If the return results of function can't be discard, this can't  use in `defer` or `goroutine`.
  - Function call can use in the expression like pass to source of function call like this: `fn1(fn2())`, but do not mixed like this: `fn1(fn2(), var1)` or `fn1(fn2(), fn3())`
  - Call `nil` function or `defer` `nil` function will cause panic error, but use `nil` in goroutine will cause fatal error -> do not `recover` -> cause crash program.

