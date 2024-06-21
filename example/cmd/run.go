package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ouqiang/goproxy"
	"github.com/ouqiang/goproxy/example/internal/envent"
)

var (
	h = flag.String("h", "", "监听地址")
	p = flag.Int("p", 8080, "监听端口")
)

var serverCmd = &cobra.Command{
	Use:   "run",
	Short: "run goproxy",
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetReportCaller(true)
		logger.SetLevel(logger.DebugLevel)
		proxy := goproxy.New(goproxy.WithDecryptHTTPS(&envent.Cache{}), goproxy.WithDelegate(&envent.EventHandler{}))
		server := &http.Server{
			Addr:         fmt.Sprintf("%s:%d", *h, *p),
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
	rootCmd.AddCommand(serverCmd)
	flag.Parse()
}
