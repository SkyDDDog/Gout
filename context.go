package gout

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H defined to save params in request
type H map[string]interface{}

// Context defined to save request & response message
type Context struct {
	// origin objects
	Writer  http.ResponseWriter
	Request *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int
}

// new one context instance
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
		index:   -1,
	}
}

// Next turn to next middleware
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// Abort end the handle process
func (c *Context) Abort() {
	c.index = len(c.handlers)
}

// AbortWithMsg end the handle process with Status & Message by json format
func (c *Context) AbortWithMsg(code int, msg string) {
	c.index = len(c.handlers)
	c.JSON(code, msg)
}

// PostForm used for getting post params
func (c *Context) PostForm(key string) string {
	return c.Request.PostFormValue(key)
}

// Query used for getting query params
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// Param used for getting params saved in the context
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// Status used for setting the StatusCode
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader used for setting Header in the response
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String used for return Strings in the response
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON used for return Json in the response
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// Data used for return bytes data in the response
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML used for return html in the response
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
