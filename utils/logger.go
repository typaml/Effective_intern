package utils

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

var (
	infoColor  = color.New(color.FgGreen).SprintFunc()
	errorColor = color.New(color.FgRed).SprintFunc()
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)
}

func Info(v ...interface{}) {
	log.Output(2, fmt.Sprintf("%s %s", infoColor("[INFO]"), fmt.Sprintln(v...)))
}

func Error(v ...interface{}) {
	log.Output(2, fmt.Sprintf("%s %s", errorColor("[ERROR]"), fmt.Sprintln(v...)))
}
