package main

import "log"

// TODO: @https://google.com // Should also work
func foo(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}

// TODO: This should do X
func bar(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}

// TODO: This should do X @https://google.com
func baz(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}
