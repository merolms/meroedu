package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogLevel(t *testing.T) {
	tests := map[string]logrus.Level{
		"":      logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}
	config := &Config{}
	for level, expected := range tests {
		config.Level = level
		err := Setup(config)
		if err != nil {
			t.Fatalf("error setting logging level %v", err)
		}
		if Logger.Level != expected {
			t.Fatalf("invalid loggin level. expected %v got %v", expected, Logger.Level)
		}
	}
	config.Level = "err"
	err := Setup(config)
	assert.Error(t, err)
}
func TestLogger(t *testing.T) {
	Debug("Debug")
	Debugf("Debug %v", "formatted")
	Info("Info")
	Infof("Info %v", "formatted")
	Error("Error")
	Errorf("Error %v", "formatted")
	Warn("Warning")
	Warnf("Warning %v", "formatted")
	WithFields(logrus.Fields{})
}
