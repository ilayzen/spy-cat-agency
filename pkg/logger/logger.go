package logger

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Level     string
	Formatter Formatter
}

// Formatter defines output formatter
type Formatter string

// Format of outputs
const (
	TextFormatter Formatter = "TXT"
	JSONFormatter Formatter = "JSON"
)

func (c Config) Validate() error {
	if !(c.Formatter == TextFormatter || c.Formatter == JSONFormatter) && c.Formatter != "" {
		return errors.New("incorrect logs formatter type")
	}

	if _, err := logrus.ParseLevel(c.Level); c.Level != "" && err != nil {
		return errors.New("incorrect logging level")
	}

	return nil
}

func NewLogger(config Config) logrus.FieldLogger {
	logger := logrus.New()

	switch config.Formatter {
	case JSONFormatter:
		logger.Formatter = &logrus.JSONFormatter{}
	default:
		logger.Formatter = &logrus.TextFormatter{}
	}

	// error checked on validation step
	lvl, err := logrus.ParseLevel(config.Level)
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(lvl)
	}

	return logger
}
