package logger

import (
	loggergo "github.com/Alonza0314/logger-go"
)

type Logger struct {
	*loggergo.FileLogger
}

func NewLogger(loggerFilePath string, debugMode bool) *Logger {
	return &Logger{
		FileLogger: loggergo.NewFileLogger(loggerFilePath),
	}
}

func (l *Logger) Kuromi() {
	l.Info("Kuromi", "⣴⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⡷")
	l.Info("Kuromi", "⠈⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣿⣿")
	l.Info("Kuromi", "⠀⢸⣿⣿⣿⣿⣷⣆⣀⣀⣀⣀⣀⣾⣿⣿⣿⣿⡇")
	l.Info("Kuromi", "⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇")
	l.Info("Kuromi", "⠀⠀⠿⢿⣿⣿⣿⣿⡏⡀⠀⡙⣿⣿⣿⣿⣿⠛")
	l.Info("Kuromi", "⠀⠀⣿⣿⣿⡿⠟⠷⣅⣀⠵⠟⢿⣿⣿⣿⡆")
	l.Info("Kuromi", "⠀⠀⠀⣿⣿⠏⢲⣤⠀⠀⠀⠀⢠⣶⠙⣿⣿⠃")
	l.Info("Kuromi", "⠀⠀⠀⠘⢿⡄⠈⠃⠀⢐⢔⠀⠈⠋⢀⡿⠋")
	l.Info("Kuromi", "⠀⠀⠀⢀⢀⣼⣷⣶⣤⣤⣭⣤⣴⣶⣍")
	l.Info("Kuromi", "⠀⠀⠀⠈⠈⣈⢰⠿⠛⠉⠉⢻⢇⠆⣁⠁")
	l.Info("Kuromi", "⠀⠀⠀⠀⠀⠑⢸⠉⠀⠀⠀⠀⠁⡄⢘⣽⣿")
	l.Info("Kuromi", "⠀⠀⠀⠀⠀⠀⡜⠀⠀⢰⡆⠀⠀⠻⠛⠋")
	l.Info("Kuromi", "⠀⠀⠀⠀⠀⠀⠑⠒⠒⠈⠈⠒⠒⠊")
	l.Info("Kuromi", "||||||||||||||||||||||||||||||")
	l.Info("Kuromi", "||||||||||||KUROMI||||||||||||")
	l.Info("Kuromi", "|||||BLESSING|||||PROGRAM|||||")
	l.Info("Kuromi", "||||||||||||||||||||||||||||||")
}