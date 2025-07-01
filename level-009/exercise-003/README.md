# Exercise 3

Using goroutines, create an incrementing program:

-   Have a variable with the value of the count

Create a bunch of goroutines, where each one must:

-   Read the value of the counter
-   Save this value
-   Yield the thread with `runtime.Gosched()`
-   Increment the saved value
-   Copy the new value to the initial variable

Use WaitGroups so that your program does not finish before your goroutines.
Prove that there is a race condition using the -race flag.
