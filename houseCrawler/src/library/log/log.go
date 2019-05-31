package log

type Config struct {
	// stdout
	Stdout bool
	// file
	Dir string
	// buffer size
	FileBufferSize int64
	// MaxLogFile
	MaxLogFile int

}
