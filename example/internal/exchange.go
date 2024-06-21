package internal

import (
	"net/url"
	"strings"

	logger "github.com/sirupsen/logrus"

	"github.com/ouqiang/goproxy/example/internal/jd"
)

func Distribution(body []byte, source string) {
	u, err := url.Parse(source)
	if err != nil {
		logger.Warnf("Failed to parse source URL: %v\n", err)
		return
	}
	logger.Infof("Parsed URL: %s\n", u.Host)
	switch {
	case strings.HasSuffix(u.Host, "jd.com"):
		jd.Item(body, source)
	}
}
