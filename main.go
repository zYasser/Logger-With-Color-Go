package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (Logger *Logger) InfoF(format string, args ...any) {
	color.Set(color.FgGreen)
	fmt.Print("[INFO]")
	color.Unset()

	Logger.infoLog.Printf(format+"\n", args...)

}
func (Logger *Logger) Info(message string) {
	color.Set(color.FgGreen)
	fmt.Print("[INFO]")
	color.Unset()

	Logger.infoLog.Println(message)

}

func (Logger *Logger) ErrorF(format string, args ...any) {
	color.Set(color.FgRed)
	fmt.Print("[ERROR]")
	color.Unset()

	Logger.infoLog.Printf(format+"\n", args...)

}
func (Logger *Logger) Error(message string) {
	color.Set(color.FgRed)
	fmt.Print("[ERROR]")
	color.Unset()

	Logger.infoLog.Println(message)

}
