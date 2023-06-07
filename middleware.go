package buggy_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/surfinggo/mc"
	"time"
)

func PrintHeaders(c *gin.Context) {
	if c.Request.Header.Get("X-Print-Headers") != "" {
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
	if code := mc.StringToInt(c.Request.Header.Get("X-Fake-Status-Code")); code != 0 {
		c.JSON(code, mc.VarOr(c.Request.Header.Get("X-Fake-Response"), "fake response"))
		c.Abort()
	} else {
		c.Next()
	}
}

func FakeResponseTime(c *gin.Context) {
	if seconds := mc.StringToInt(c.Request.Header.Get("X-Fake-Response-Seconds")); seconds != 0 {
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

func UseAll(engine *gin.Engine) {
	engine.Use(PrintHeaders)
	engine.Use(FakeResponseTime)

	engine.Use(FakeStatusCode) // might abort, put it last
}
