package concurrency

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

type favContextKey string

func TestContext(t *testing.T) {
	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("start call context cancel,",time.Now())
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine,",time.Now())
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stop getting url:%s %v\n", url, time.Now())
			return
		default:
			fmt.Printf("start getting url:%s %v\n", url, time.Now())
			r, err := http.Get(url)
			if r.StatusCode == http.StatusOK && err == nil {
				body, _ := ioutil.ReadAll(r.Body)
				subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
				wg.Add(1)
				go showResp(subCtx, wg)
			}
			r.Body.Close()
			//启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒
			// doSometing()
			time.Sleep(time.Second * 1)
		}
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop showing resp,",time.Now())
			return
		default:
			//子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印
			fmt.Println("printing ", ctx.Value(favContextKey("resp")), time.Now())
			time.Sleep(time.Second * 1)
		}
	}
}
