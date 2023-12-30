package main

import (
	config "Go-Project-Template/init"
)

func main() {
	config.InitConfig()
	config.ZapInit()
	//config := zap.NewDevelopmentConfig()
	//build, _ := config.Build()
	//build.Error("Sss")
	//build.
	config.Logger.Debug("debug")
	config.Logger.Info("Info")
	config.Logger.Warn("warn")
	config.Logger.Error("Error")
	//config.Logger.Panic("Panic")
}
