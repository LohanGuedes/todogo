package main

import "log"

// TODO: Alguma coisa que descreve a tarefa
// @ticket https://google.com
// asdfkjaf
func foo(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}

// TODO: A small description
// @ticket https://google.com
func bar(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}

// TODO A small description
func bez(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}

// TODO
func baz(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}
