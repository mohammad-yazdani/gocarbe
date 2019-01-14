package shared

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var logFile *os.File

func SetUpLog() {
	logName := strings.Builder{}
	logName.WriteString("carbe_server_log.")
	logName.WriteString(strconv.Itoa(int(time.Now().UnixNano())))
	logName.WriteString(".log")

	toBeLog, err := os.Create(logName.String())
	if err != nil {
		log.Panic(err)
	}

	logFile = toBeLog

	fmt.Println("SERVER LOG: " + logName.String())
}

func CloseLog() {
	logFile.Close()
}

func LOG(msg string) {

	timedLog := strings.Builder{}
	timedLog.WriteString("[")
	timedLog.WriteString(time.Now().Format(time.UnixDate))
	timedLog.WriteString("] ")

	timedLog.WriteString(msg)
	timedLog.WriteString("\n")

	_, err := logFile.WriteString(timedLog.String())
	if err != nil {
		log.Panic(err)
	}
}
