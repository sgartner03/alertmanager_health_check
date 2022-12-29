package logging

import (
	"go.uber.org/zap"
	"encoding/json"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() Logger {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	}`)
  	var cfg zap.Config
  	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
	   panic(err)
	}
	
	var wrapper Logger
	wrapper.logger = logger
	return wrapper
}

func (logger Logger) Info(msg string)  {
	defer logger.logger.Sync()
	logger.logger.Info(msg)
}

func (logger Logger) Warn(msg string)  {
	logger.logger.Warn(msg)
}

func (logger Logger) Error(msg string)  {
	logger.logger.Error(msg)
}