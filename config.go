// Global configuration.
package main

type Config struct {
	// Screen
	width  int
	height int

	// Visual configuration
	lineLen float64

	// Game rules
	roundDuration int // in seconds
}

var globalConfig = Config{
	width:  800,
	height: 300,

	lineLen: 250,

	roundDuration: 3,
}
