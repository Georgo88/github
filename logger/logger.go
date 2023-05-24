package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// if output true save in console, else in file
func Log(level string, output bool, data any) {
	currentTime := time.Now().Local().Format("2017.09.07 17:06:06")
	if output == true {
		fmt.Printf("Log %s: %s %s \n", level, data, currentTime)
	} else {
		os.Mkdir("logs", 0755)
		f, err := os.Create(fmt.Sprintf("%s %s.txt", level, currentTime))
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		Writer := bufio.NewWriter(f)
		_, err = fmt.Fprintf(Writer, "Log %s: %s %s \n", level, data, currentTime)
		if err != nil {
			fmt.Println(err)
		}
		Writer.Flush()
	}

}

func Info(output bool, data interface{}) {
	Log("Info", output, data)
}

func Error(output bool, data interface{}) {
	Log("Error", output, data)
}

func Warning(output bool, data interface{}) {
	Log("Warning", output, data)
}

func Fatal(output bool, data interface{}) {
	Log("Fatal", output, data)
}
