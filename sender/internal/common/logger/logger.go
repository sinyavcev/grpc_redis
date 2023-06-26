package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(logLevel string) (*Logger, error) {
	log := logrus.New()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return &Logger{log}, fmt.Errorf("logrus.ParseLevel: %w", err)
	}
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return &Logger{log}, nil
}
