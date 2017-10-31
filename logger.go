package pgt

// Logger defines the reporting interface for utils error functions
type Logger interface {
	// Fatal should stop the execution
	Fatal(msg string, ctx ...interface{})
}
