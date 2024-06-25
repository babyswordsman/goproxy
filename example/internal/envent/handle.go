package envent

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"

	logger "github.com/sirupsen/logrus"

	"github.com/ouqiang/goproxy"
	"github.com/ouqiang/goproxy/example/internal"
)

type EventHandler struct{}

func (e *EventHandler) Connect(ctx *goproxy.Context, rw http.ResponseWriter) {

	logger.Infof("connect:%v", ctx.Req.URL)
	// 保存的数据可以在后面的回调方法中获取
	ctx.Data["req_id"] = "uuid"

	// 禁止访问某个域名
	if strings.Contains(ctx.Req.URL.Host, "example.com") {
		rw.WriteHeader(http.StatusForbidden)
		ctx.Abort()
		return
	}
}

func (e *EventHandler) Auth(ctx *goproxy.Context, rw http.ResponseWriter) {
	// 身份验证
}

func (e *EventHandler) BeforeRequest(ctx *goproxy.Context) {
	// 修改header
	ctx.Req.Header.Add("X-Request-Id", ctx.Data["req_id"].(string))
	// 设置X-Forwarded-For
	if clientIP, _, err := net.SplitHostPort(ctx.Req.RemoteAddr); err == nil {
		if prior, ok := ctx.Req.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		ctx.Req.Header.Set("X-Forwarded-For", clientIP)
	}
	// 读取Body
	body, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		// 错误处理
		return
	}
	// Request.Body只能读取一次, 读取后必须再放回去
	// Response.Body同理
	ctx.Req.Body = io.NopCloser(bytes.NewReader(body))

}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	data, err = io.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}

func (e *EventHandler) BeforeResponse(ctx *goproxy.Context, resp *http.Response, err error) {
	if err != nil {
		return
	}
	// 修改response
	// 读取Body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		// 错误处理
		logger.Errorf("read resp body err:%s", err.Error())
		return
	} else {
		contentType := resp.Header.Get("Content-Type")
		encoding := resp.Header.Get("Content-Encoding")

		//todo: 处理字符集
		buf := bytes.Buffer{}
		if strings.Contains(encoding, "gzip") {
			err = gunzipWrite(&buf, body)
			if err != nil {
				logger.Error("unzip err:", err.Error())
				return
			}
		}
		if buf.Len() == 0 {
			buf.Write(body)
		}

		if strings.Contains(strings.ToLower(contentType), "text/html") ||
			strings.Contains(strings.ToLower(contentType), "application/json") {

			internal.Distribution(buf.Bytes(), ctx.Req.URL.String(), ctx.Req.Method,
				strings.ToLower(contentType))
		}
	}

	// Request.Body只能读取一次, 读取后必须再放回去
	// Response.Body同理
	resp.Body = io.NopCloser(bytes.NewReader(body))
}

// 设置上级代理
func (e *EventHandler) ParentProxy(req *http.Request) (*url.URL, error) {
	return http.ProxyFromEnvironment(req)
}

func (e *EventHandler) Finish(ctx *goproxy.Context) {
	logger.Infof("请求结束 URL:%s\n", ctx.Req.URL)
}

// 记录错误日志
func (e *EventHandler) ErrorLog(err error) {
	logger.Errorf("err:{ %s }", err.Error())
}

// WebSocketSendMessage websocket发送消息
func (e *EventHandler) WebSocketSendMessage(ctx *goproxy.Context, messageType *int, p *[]byte) {}

// WebSockerReceiveMessage websocket接收 消息
func (e *EventHandler) WebSocketReceiveMessage(ctx *goproxy.Context, messageType *int, p *[]byte) {}

// 实现证书缓存接口
type Cache struct {
	m sync.Map
}

func (c *Cache) Set(host string, cert *tls.Certificate) {
	c.m.Store(host, cert)
}
func (c *Cache) Get(host string) *tls.Certificate {
	v, ok := c.m.Load(host)
	if !ok {
		return nil
	}

	return v.(*tls.Certificate)
}
