## 1. Goroutines

- There are always exist one Goroutine call `main goroutine`.
- The main goroutine do not know when the another goroutines will finish, so if it exits, the ones will exist corresponding.
- We use the `sync` package to make the goroutine await the anothers goroutines be finished:

  - `Add(n)` to add number of running Goroutines, n can a negative number.
  - `Done()` to notice the main goroutine this is finished
  - `Wait()` make a goroutine is blocking util the others goroutine are finished.

- This technique be call `concurrentcy synchonization`.

- Goroutine has two state: `running` and `blocking`. A blocking goroutine just be unblocking by the another goroutine.
- If all goroutines are blocking state, program will into the deadlock -> Go runtime try to crash the program.

- Each logical CPU can only execute one goroutine. Go runtime must frequently switch execution contexts between goroutines to let each running goroutine have a chance to execute.

## 2. Deferred function call

## 3. Panic/recover
