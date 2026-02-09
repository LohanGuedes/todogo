package td

// TODO: Alguma coisa que descreve a tarefa
// @ticket https://google.com
// Some more useful description of what is needed to be done
// Will not be parsed, but can be useful for devs reading the todo.
func foo() {}

// TODO: A small description
// @ticket https://google.com
func bar() {}

// TODO: A small description // want "TODO's must have \"@ticket\" comment-property"
func baz() {}

// TODO A small description // want "TODO's must have a ':' before its description"
func oof() {}

// TODO // want "TODO's must have a ':' before its description"
func foof() {}

// TODO A small description // want "TODO's must have a ':' before its description"
// @ticket https://google.com
func faaf() {}
