package goroutine_error_handle

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	for _, url := range urls {
		var u = url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			resp, err := http.Get(u)
			if err == nil {
				fmt.Printf("获取%s成功\n", u)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
