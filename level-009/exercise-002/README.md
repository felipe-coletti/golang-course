# Exercise 2

This exercise will reinforce your knowledge of method sets.

Create a type for a struct called `person`.

Create a method `say` for this type that has a receiver pointer `*person`.

Create an interface, `human`, that is implemented by types with the method `say`.

Create a function `say` whose parameter is of type `human` and that calls the
method `say`.

Demonstrate in your code:

-   That you can use a value of type `*person` in the function `say`
-   That you cannot use a value of type `person` in the function `say`
