package server

import (
	"fmt"
	"os"
	"time"
)

type TracePrefix string

const (
	_control TracePrefix = "control"
	_env     TracePrefix = "env"
	_error   TracePrefix = "error"
	_build   TracePrefix = "build"
	_link    TracePrefix = "link"
	_map     TracePrefix = "map"
	_https   TracePrefix = "https"
)

func trace(prefix TracePrefix, format string, args ...interface{}) {
	// external process control should redirect stdout to a log file
	fmt.Fprintf(os.Stdout, "%s: %s: %s\n", time.Now().UTC().Format(time.RFC3339), prefix, fmt.Sprintf(format, args...))
}
