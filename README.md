# Todogo

Todogo is a linter that integrates with golangci-lint, used to check for "missing" todos in your codebase.

No more reason to forget TODO comments on your application's code.

todogo uses SA to look for comments in your codebase.

For example, the code below might be lost in a huge codebase, and its hard to give visibility to your stakeholders of that is missing, and when untracked tasks like these get lost and usually the force to get back to it is never taken into account considering the urgency of them.
Therefore, we MUST be able to track them and a simple solution is using a TODO-parser script (to be linked here...) to parse the current state todos, creating tickets using the parser.

But having this does not solve for future added tasks within the codebase. Therefore a Linter is the GOTO to block in your CI, the possibility of tasks being forgotten.
```go
package main

// NOT OKAY!!!!!
// TODO
func foo() {
    // [...] 
}


// NOT OKAY!!!!!
// TODO: create a middleware that does X
func bar() {
    // [...] 
}


// NOT OKAY (per config)!!!!!
// TODO: @https://tracking.software/browse/TCK-2388
func bar() {
    // [...] 
}

// OK!!!!
// TODO: @https://your.company.atlassian.net/browse/TCK-2388 Add a new linter that accounts for TODOS
func baz() {
    // [...] 
}


// OK!!!!
// TODO:
// @ticket @https://tracking.software/browse/TCK-2388
// @desc Add a new linter that accounts for TODOS
func baz() {
    // [...] 
}
```
