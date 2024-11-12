# Functions

**_Functions_** are central in Go. Let's learn about functions with a few different examples.

Here's a function that takes two `integers` and returns their sum as an `int`
 ```
 func plus(a int, b int) int {

 }
 ```

 Go requires explicit returns i.e., it won't automatically return the value of the last expression.
 ```
func plus(a int, b int) int {
    return a + b
}
 ```

 If there are multiple consecutive parameters of the same type, then the user may omit the type name for the like-typed parameters up to the final parameter that declares the type.
```
func plusPlus(a, b, c int) int {
    return a + b + c
}
```
Call a function as expected, with name(args).
```
func main() {
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
```

## Multiple Return Values

Go has built-in support for _multiple return values_. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

The (int, int) in this function signature shows that the function returns 2 ints.
```
func vals() (int, int) {
    return 3, 7
}
```

In this example, 2 different return values are used from the call with multiple assignment. If user wants a subset of the returned values, use the blank identifier _.
```
func main() {
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
    _, c := vals()
    fmt.Println(c)
}
```

## Variadic Functions


<ins>_`Variadic functions`_</ins> can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

Here’s a function that will take an arbitrary number of integers as arguments.

```
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
}
```

Within the function, the type of `nums` is equivalent to []int. We can call `len(nums)`, iterate over it with range, etc.
```
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
     for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

Variadic functions can be called in the usual way with individual arguments. If the user already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
```
func main() {
    sum(1, 2)
    sum(1, 2, 3)
    nums := []int{1, 2, 3, 4}
    sum(nums...)
```

## Closures

Go supports <ins>_`anonymous functions`_</ins>, which can form <ins>_`closures`_</ins>. Anonymous functions are useful when user wants to define a function inline without having to name it.

In the below example, the function `intSeq` returns another function, which user defines anonymously in the body of `intSeq`. The returned function _`closes over`_  the variable i to form a closure.
```
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

After `intSeq` is called, it assigns the result (a function) to `nextInt`.
This function value captures it own `i` values, which will be updated each time we call `nextInt`.
```
func main() {
    nextInt := intSeq()
}
```
Observe the effect of the closure by calling `nextInt` a few times. To confirm that the state is unique to that particular function, create and test a new one.
```
func main() {
    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    // Check if state is unique to particular function
    newInts := intSeq()
    fmt.Println(newInts())
}
```

## Recursion


Go supports <ins>_recursive functions_</ins>. Here’s a classic example.

In below example the `fact` function calls itself until it reaches the base case of `fact(0)`.
```
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
```

Anonymous functions can also be recursive, but this requires explicitly declaring a variable with `var` to store the function before it’s defined. Since `fib` was previously declared in main, Go knows which function to call with `fib` here.
```
func main() {
    fmt.Println(fact(7))
    var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}
```
