package k6

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

type User struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func TestHello(t *testing.T) {

	server := gin.Default()
	server.POST("/hello", func(ctx *gin.Context) {
		var u User

		if err := ctx.Bind(&u); err != nil {
			return
		}

		num := rand.Int31n(1000) + 1

		time.Sleep(time.Millisecond * time.Duration(num))

		if num%100 < 10 {
			ctx.String(http.StatusInternalServerError, "失败")
		} else {
			ctx.String(http.StatusOK, "成功")
		}
	})

	err := server.Run(":8888")
	if err != nil {
		panic(err)
	}
}
