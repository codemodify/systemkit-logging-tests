package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingWithFields "github.com/codemodify/systemkit-logging-extenders-withfields"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-consolewithcolors"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-file"
)

func Test_03(t *testing.T) {
	withFieldsLogger := loggingWithFields.NewWithFieldsLogger(
		logging.NewPipe([]logging.CoreLogger{
			loggingFormat.NewTimeRFC3339NanoFormatter(),
			loggingMixMulti.NewMultiLogger([]logging.CoreLogger{
				loggingPersistConsole.NewConsoleLoggerDefaultColors(),
				loggingPersistFile.NewFileLoggerDefaultName(),
			}),
		}),
	)

	withFieldsLogger.TraceWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Trace line",
	})
	withFieldsLogger.PanicWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Panic line",
	})
	withFieldsLogger.FatalWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Fatal line",
	})
	withFieldsLogger.ErrorWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Error line",
	})
	withFieldsLogger.WarningWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Warning line",
	})
	withFieldsLogger.InfoWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Info line",
	})
	withFieldsLogger.SuccessWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Success line",
	})
	withFieldsLogger.DebugWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Debug line",
	})
}
