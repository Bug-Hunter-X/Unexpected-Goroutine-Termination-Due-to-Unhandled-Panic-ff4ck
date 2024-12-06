# Unexpected Goroutine Termination Due to Unhandled Panic in Go
This example demonstrates how a `panic` in a Go program can unexpectedly terminate only the goroutine in which it occurs.  This can lead to subtle errors if not handled properly using `recover()` within a `defer` statement.
The `bug.go` file shows a `panic` that is not handled.  The `bugSolution.go` file demonstrates how to properly use `recover()` to handle the panic.

## How to reproduce the bug:
1. Save the code from `bug.go`.
2. Run the code using `go run bug.go`.
3. Observe that the program does not print "After the panic".

## How to fix the bug:
1. Review the code and identify potential panic points.
2. Use `defer` and `recover()` in functions to gracefully handle panics.
3. Ensure that all error handling mechanisms are designed to prevent unexpected goroutine terminations.
