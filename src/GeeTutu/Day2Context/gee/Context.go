package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 是类别别名
type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter // 响应对象
	Req    *http.Request       // 请求对象

	Path   string // 请求路径
	Method string // 请求方法（GET POST PUT DELETE 等）

	StatusCode int // 返回的状态码
}

// newContext 工厂方法，由 http 服务器注入请求信息，并存入 Context 实例中
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 获取请求表单中的参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取请求参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头，如 Content-Type: text/html
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 响应字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 响应 JSON 数据
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 响应二进制数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 响应 HTML
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
