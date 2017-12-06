# sd
Typed Lambda Calculus - programming language

### Example
`(λ a: Bool → succ succ 0) iszero 0`
=> to 2

### How it works
Lambda Calculus consists of anonymous functions defined by `λ`. Each function has one argument and its type identified by `:`. Two types are supported `Nat` and `Bool`. 

Some predefined functions are `succ`, adds one to the argument. `pred` removes 1 from the argument. `iszero` checks if argument evaluates to zero or not.

##### Followed this fantastic tutorial [http://blog.mgechev.com/2017/08/05/typed-lambda-calculus-create-type-checker-transpiler-compiler-javascript/](http://blog.mgechev.com/2017/08/05/typed-lambda-calculus-create-type-checker-transpiler-compiler-javascript/)
