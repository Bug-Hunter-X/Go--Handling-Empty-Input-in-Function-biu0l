# Go: Handling Empty Input in Function

This repository demonstrates a common error in Go: forgetting to handle empty inputs in functions. The `bug.go` file showcases a function that doesn't handle an empty slice, leading to issues.  `bugSolution.go` provides a corrected version with appropriate error handling.

## Bug Description

The `processData` function in `bug.go` fails to check if the input `data` slice is empty.  If it is, it might panic or produce unintended results. This demonstrates the importance of robust input validation.

## Solution

The `bugSolution.go` file provides a solution that explicitly checks for an empty slice using `len(data) == 0`. If the slice is empty, it returns an appropriate error; otherwise, it proceeds with the summation.