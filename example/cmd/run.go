package cmd

import (
	"fmt"
	"net/http"
	"time"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ouqiang/goproxy"
	"github.com/ouqiang/goproxy/example/internal/envent"
	"github.com/ouqiang/goproxy/example/internal/jd"
)

var (
	host string
	port int
)

var serverCmd = &cobra.Command{
	Use:   "run",
	Short: "run goproxy",
	Run: func(cmd *cobra.Command, args []string) {
		jd.Init()
		logger.SetReportCaller(true)
		logger.SetLevel(logger.DebugLevel)
		proxy := goproxy.New(goproxy.WithDecryptHTTPS(&envent.Cache{}), goproxy.WithDelegate(&envent.EventHandler{}))
		server := &http.Server{
			Addr:         fmt.Sprintf("%s:%d", host, port),
			Handler:      proxy,
			ReadTimeout:  1 * time.Minute,
			WriteTimeout: 1 * time.Minute,
		}
		logger.Infof("Listen %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(&host, "host", "h", "", "Listen host")
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Listen port")
	serverCmd.Flags().StringVarP(&jd.Filename, "file", "f", "jd.json", "storage file path")

	rootCmd.AddCommand(serverCmd)
}
