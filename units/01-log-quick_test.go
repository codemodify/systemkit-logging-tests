package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
)

func Test_01(t *testing.T) {
	logging.Instance().Trace("Trace line")
	logging.Instance().Panic("Panic line")
	logging.Instance().Fatal("Fatal line")
	logging.Instance().Error("Error line")
	logging.Instance().Warning("Warning line")
	logging.Instance().Info("Info line")
	logging.Instance().Success("Success line")
	logging.Instance().Debug("Debug line")
}
