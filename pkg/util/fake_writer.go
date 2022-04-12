package util

import (
	"bytes"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// BodyLogWriter ...
type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

// BuildFakeContextWriter return fake context and fake writer which you can retreive body from it
func BuildFakeContextWriter() (*gin.Context, *BodyLogWriter) {
	// Create test context to hold data return
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Clone buffer body
	blw := &BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	return c, blw
}
