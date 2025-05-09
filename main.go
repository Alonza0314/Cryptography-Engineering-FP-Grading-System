package main

import (
	"ce/backend/factory"
	logger "ce/backend/logger"
	"fmt"
	"time"
)

/*
	⣴⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⡷
	⠈⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣿⣿
	⠀⢸⣿⣿⣿⣿⣷⣆⣀⣀⣀⣀⣀⣾⣿⣿⣿⣿⡇
	⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
	⠀⠀⠿⢿⣿⣿⣿⣿⡏⡀⠀⡙⣿⣿⣿⣿⣿⠛
	⠀⠀⣿⣿⣿⡿⠟⠷⣅⣀⠵⠟⢿⣿⣿⣿⡆
	⠀⠀⠀⣿⣿⠏⢲⣤⠀⠀⠀⠀⢠⣶⠙⣿⣿⠃
	⠀⠀⠀⠘⢿⡄⠈⠃⠀⢐⢔⠀⠈⠋⢀⡿⠋
	⠀⠀⠀⢀⢀⣼⣷⣶⣤⣤⣭⣤⣴⣶⣍
	⠀⠀⠀⠈⠈⣈⢰⠿⠛⠉⠉⢻⢇⠆⣁⠁
	⠀⠀⠀⠀⠀⠑⢸⠉⠀⠀⠀⠀⠁⡄⢘⣽⣿
	⠀⠀⠀⠀⠀⠀⡜⠀⠀⢰⡆⠀⠀⠻⠛⠋
	⠀⠀⠀⠀⠀⠀⠑⠒⠒⠈⠈⠒⠒⠊
	||||||||||||||||||||||||||||||
	||||||||||||KUROMI||||||||||||
	|||||BLESSING|||||PROGRAM|||||
	||||||||||||||||||||||||||||||
*/

var (
	Logger *logger.Logger
)

func main() {
	var err error

	ceCfg, err := factory.LoadCeCfg("config/ceCfg.yaml")
	if err != nil {
		panic(fmt.Sprintf("Failed to load ceCfg: %v", err))
	}

	Logger = logger.NewLogger(ceCfg.Logger.FileDir+"/"+time.Now().Format("2006-01-02")+".log", ceCfg.Logger.DebugMode)
	Logger.Kuromi()
}
