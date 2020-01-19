// Global configuration.
package main

type Config struct {
	// Screen
	width  int
	height int

	// Visual configuration
	lineLen float64
}

var globalConfig = Config{
	width:  800,
	height: 800,

	lineLen: 1000,
}
