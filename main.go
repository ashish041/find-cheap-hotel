package main

import (
	"github.com/ashish041/find-cheap-hotel/api/hotel"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/ashish041/find-cheap-hotel/util/agent"
	"github.com/ashish041/find-cheap-hotel/util/logger"
	"github.com/ashish041/find-cheap-hotel/util/peon"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var (
	optPort       int
	optDebugPort  int
	optGinMode    string
	optLogToFile  bool
	optPathPrefix string
)

const (
	dstLogFile = "/tmp/web.log"

	HTTP_UNAUTHORIZE_ACCESS  = 401
	HTTP_FORBIDDEN_ACCESS    = 403
	HTTP_FOUND               = 302
	HTTP_STATUS_SUCCESS      = 200
	HTTP_SERVICE_UNAVAILABLE = 503
	HTTP_NOTFOUND            = 404
)

func init() {
	flag.IntVar(&optPort, "port", 9000, "Http running on the port")
	flag.IntVar(&optDebugPort, "debugPort", 9090, "Port to listen to for debugging, set to 0 to disable")
	flag.StringVar(&optGinMode, "ginMode", "release", "Gin webframework running on release mode")
	flag.BoolVar(&optLogToFile, "logToFile", true, "Log write to file")
}

func setupLogging() io.WriteCloser {
	w := logger.MustWriteTo(dstLogFile)
	log.SetOutput(w)
	log.SetPrefix(fmt.Sprintf("%d ", os.Getpid()))
	return w
}

func noCache(c *gin.Context) {
	// As documented in: http://stackoverflow.com/questions/49547
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")                                         // Proxies.
}

func Flags() int {
	return log.Ldate | log.LUTC | log.Lmicroseconds | log.Lshortfile
}

func main() {
	peon.QuitIfRoot()
	/*main.go run from command line: go run main.go -port=9000 -ginMode=debug -productionMode=false*/
	flag.Parse()
	log.SetFlags(Flags())
	var logfile io.WriteCloser

	logfile = os.Stderr
	if optLogToFile {
		logfile = setupLogging()
	}
	defer agent.Listen().Close()
	if optDebugPort > 0 {
		go func() {
			log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", optDebugPort), nil))
		}()
	}
	gin.SetMode(optGinMode)
	route := gin.Default()

	// This needs to be able to called without requiring access
	route.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true,
	}))
	route.GET("/api/hotel", func(ctx *gin.Context) {
		ret, err := hotel.GetHotelLists()
		if err != nil {
			ctx.JSON(HTTP_NOTFOUND, err.Error())
		} else {
			ctx.JSON(HTTP_STATUS_SUCCESS, ret)
		}
	})
	route.GET("/knockknock", func(ctx *gin.Context) {
		ctx.String(HTTP_STATUS_SUCCESS, "Service is running...\n\r")
	})
	endless.DefaultMaxHeaderBytes = 1 << 20 //1 MB
	//It will avoid hanging connections that the client has no intention of closing
	if err := endless.ListenAndServe(fmt.Sprintf(":%d", optPort), route); err != nil {
		log.Printf("Http not running: %v\n", err)
		if optLogToFile {
			logfile.Close()
		}
		os.Exit(1)
	} else {
		log.Println("shutting down")
	}
	if optLogToFile {
		logfile.Close()
	}
	os.Exit(0)
}
