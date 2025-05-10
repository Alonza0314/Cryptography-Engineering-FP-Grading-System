package main

import (
	"ce/backend/api"
	"ce/backend/factory"
	logger "ce/backend/logger"
	"ce/backend/mongo"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	CeCfg  *factory.CeCfg
	Logger *logger.Logger
)

func main() {
	var err error

	ceCfg, err := factory.LoadCeCfg("config/ceCfg.yaml")
	if err != nil {
		panic(fmt.Sprintf("Failed to load ceCfg: %v", err))
	}
	factory.CeConfig = ceCfg

	Logger = logger.NewLogger(ceCfg.Logger.FileDir+"/"+time.Now().Format("2006-01-02")+".log", ceCfg.Logger.DebugMode)
	api.Log = Logger
	Logger.Kuromi()

	db, err := mongo.NewMongo()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}
	api.Db = db
	defer api.Db.Close()

	router := api.NewRouter()
	srv := &http.Server{
		Addr:    ":" + ceCfg.Gin.Port,
		Handler: router,
	}

	go func() {
		Logger.Info("GIN", "Server started on port "+ceCfg.Gin.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Error("GIN", "Failed to start server: "+err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	Logger.Info("GIN", "Received signal: "+sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	Logger.Info("GIN", "Server shutdown gracefully")
}
