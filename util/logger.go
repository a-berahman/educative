package util

import (
	"fmt"

	"go.uber.org/zap"
)

var lg *zap.SugaredLogger

//Initialize prepares basic items for Log process
func Initialize() {
	logger, _ := zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			fmt.Println("Failed to flush logger: ", err)
		}
	}()

	lg = logger.Sugar()
}

//Logger returns an instance of current logger package
func Logger() *zap.SugaredLogger {
	return lg
}
