package tests

import (
	"testing"
	"time"

	loggingAdvanced "github.com/codemodify/systemkit-logging-advanced"
)

// Special Cases Logger - `NewGroupAndSort()` is super useful in
// concurrent apps for debugging exact sequences of logs occurred with-in the tree
// of parent-child threads. The `logLevel` will get increased for the child and also affects
// the tabulation when reading the logs. Looks like a "call-stack" type of thing clearly showing what
// what log line was called from what parent.
//
//		This logger stores the log entries in-memory.
//		Groups the log entries by the LOG-TAG
//		Sorts the log entries in each group by the time
//		When you call `Flush()` it will save to the persisters
//		a nice-call-stack-alike log lines

func Test_05(t *testing.T) {
	var flushable = loggingAdvanced.NewGroupAndSort(
		loggingAdvanced.NewSimpleFormatter(),
	)
	var groupAndSort = loggingAdvanced.NewDefaultLoggerImplementation(flushable)

	//
	// ... smart concurrent stuff that gets logged, grouped and sorted by time and call-stack
	//
	go func(logTag string, logLevel int) {
		groupAndSort.InfoWithTagAndLevel(logTag, logLevel, "Info line")

		go func(logTag string, logLevel int) {
			groupAndSort.InfoWithTagAndLevel(logTag, logLevel, "Info line")

			go func(logTag string, logLevel int) {
				groupAndSort.InfoWithTagAndLevel(logTag, logLevel, "Info line")
			}(logTag, logLevel+1)
		}(logTag, logLevel+1)

		groupAndSort.InfoWithTagAndLevel(logTag, logLevel, "Info line")
	}("LOG-TAG-1", 0)

	go func(logTag string, logLevel int) {
		groupAndSort.WarningWithTagAndLevel(logTag, logLevel, "Warning line")

		go func(logTag string, logLevel int) {
			groupAndSort.WarningWithTagAndLevel(logTag, logLevel, "Warning line")

			go func(logTag string, logLevel int) {
				groupAndSort.WarningWithTagAndLevel(logTag, logLevel, "Warning line")
			}(logTag, logLevel+1)
		}(logTag, logLevel+1)

		groupAndSort.WarningWithTagAndLevel(logTag, logLevel, "Warning line")
	}("LOG-TAG-2", 0)

	go func(logTag string, logLevel int) {
		groupAndSort.ErrorWithTagAndLevel(logTag, logLevel, "Error line")

		go func(logTag string, logLevel int) {
			groupAndSort.ErrorWithTagAndLevel(logTag, logLevel, "Error line")

			go func(logTag string, logLevel int) {
				groupAndSort.ErrorWithTagAndLevel(logTag, logLevel, "Error line")
			}(logTag, logLevel+1)
		}(logTag, logLevel+1)

		groupAndSort.ErrorWithTagAndLevel(logTag, logLevel, "Error line")
	}("LOG-TAG-3", 0)

	//
	// ... more smart concurrent stuff
	//

	// Dump log entries from memory to the configured pipe-line
	time.Sleep(5 * time.Second)
	flushable.Flush()
}
