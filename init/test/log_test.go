package test

import (
	config "Go-Project-Template/init"
	"testing"
)

func TestLog(t *testing.T) {
	config.InitConfig()
	config.ZapInit()
	config.Logger.Info("sakura")
}
