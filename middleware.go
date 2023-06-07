package buggy_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/surfinggo/mc"
	"os"
	"time"
)

func PrintHeaders(c *gin.Context) {
	if os.Getenv("X-Print-Headers") != "" || c.Request.Header.Get("X-Print-Headers") != "" {
		fmt.Println()
		fmt.Println("Headers:")
		for k, v := range c.Request.Header {
			if len(v) > 1 {
				fmt.Printf("  %s:\n", k)
				for _, val := range v {
					fmt.Printf("  - %v\n", val)
				}
			} else {
				fmt.Printf("  %s: %v\n", k, v[0])
			}
		}
		fmt.Println()
	}
}

func FakeStatusCode(c *gin.Context) {
	codeHeader := mc.StringToInt(c.Request.Header.Get("X-Fake-Status-Code")) // higher priority
	codeEnv := mc.StringToInt(os.Getenv("X-Fake-Status-Code"))
	if code := mc.VarOr(codeHeader, codeEnv); code != 0 {
		c.JSON(code, mc.VarOr(c.Request.Header.Get("X-Fake-Response"), "fake response"))
		c.Abort()
	} else {
		c.Next()
	}
}

func FakeResponseTime(c *gin.Context) {
	millisecondsHeader := mc.StringToInt(c.Request.Header.Get("X-Fake-Response-Milliseconds")) // higher priority
	millisecondsEnv := mc.StringToInt(os.Getenv("X-Fake-Response-Milliseconds"))
	if milliseconds := mc.VarOr(millisecondsHeader, millisecondsEnv); milliseconds != 0 {
		time.Sleep(time.Duration(milliseconds) * time.Second)
	}
}

func UseAll(engine *gin.Engine) {
	engine.Use(PrintHeaders)
	engine.Use(FakeResponseTime)

	engine.Use(FakeStatusCode) // might abort, put it last
}
