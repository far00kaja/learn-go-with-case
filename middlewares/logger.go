package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}

func SetupLogger() {
	if _, err := os.Stat("/logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModeAppend.Perm())
	}

	f, err := os.OpenFile("logs/auth-service-"+time.Now().UTC().Format("2006-01-02")+".log", os.O_APPEND, 0544)
	if err != nil {
		f, err := os.Create("logs/auth-service-" + time.Now().UTC().Format("2006-01-02") + ".log")
		if err != nil {
			panic(err)
		}
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	} else {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

}
