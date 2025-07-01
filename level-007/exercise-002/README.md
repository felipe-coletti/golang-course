# Exercise 2

Create a struct `person`.

Create a function called `setName` that has `*person` as a parameter. This function should change a value stored at the address `*person`.

Hint: the "correct" way to dereference a value in a struct would be `(*value).field`.

But there is an exception in the documentation. Link: https://golang.org/ref/spec#Selectors.

"As an exception, if the type of `x` is a named pointer type and `(*x).f` is a valid selector expression denoting a field (but not a method), `x.f is shorthand for (*x).f`."

In other words, we can use both the shortcut `p1.name` and the traditional `(*p1).name`.

Create a value of type `person`.

Use the `setName` function and demonstrate the result.
