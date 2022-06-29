package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
)

var db = make(map[string]string)

func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			fmt.Println()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": rec,
			})
		}
	}(c)
	c.Next()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(globalRecover)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithError(404, errors.New("not found"))
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//        "foo":  "bar",
	//        "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}

	r.Run(":8080")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func inOrderTraversalByIter(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil && root.Left != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			root = node.Right
		}
	}
	return res
}

func longestLine(M [][]byte) int{
    if len(M) ==0 || len(M[0]) == 0 {
        return 0
    }
    res,m,n := 0,len(M),len(M[0])
    row := make([]int,m)
    col := make([]int,n)
    d   := make([]int,m+n)
    ad  := make([]int,m+n)
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if M[i][j] == 'F'{
                row[i]++
                col[j]++
                d[i+j]++
                ad[j-i+m]++
                res = max(res, max(row[i], col[j]))
                res = max(res, max(d[i+j], ad[j-i+m]))
            }else{
                row[i] = 0
                col[j] = 0
                d[i+j] = 0
                ad[j-i+m] = 0
            }
        }
    }
    return res                  
}

type Counter struct {
    rate  int           //计数周期内最多允许的请求数
    begin time.Time     //计数开始时间
    cycle time.Duration //计数周期
    count int           //计数周期内累计收到的请求数
    lock  sync.Mutex
}

func (l *Counter) Allow() bool {
    l.lock.Lock()
    defer l.lock.Unlock()

    if l.count == l.rate-1 {
        now := time.Now()
        if now.Sub(l.begin) >= l.cycle {
            //速度允许范围内， 重置计数器
            l.Reset(now)
            return true
        } else {
            return false
        }
    } else {
        //没有达到速率限制，计数加1
        l.count++
        return true
    }
}

func (l *Counter) Set(r int, cycle time.Duration) {
    l.rate = r
    l.begin = time.Now()
    l.cycle = cycle
    l.count = 0
}

func (l *Counter) Reset(t time.Time) {
    l.begin = t
    l.count = 0
}

func main() {
    var wg sync.WaitGroup
    var lr Counter
    lr.Set(3, time.Second) // 1s内最多请求3次
    for i := 0; i < 10; i++ {
        wg.Add(1)
        log.Println("创建请求:", i)
        go func(i int) {
          if lr.Allow() {
              log.Println("响应请求:", i)
          }
          wg.Done()
        }(i)

        time.Sleep(200 * time.Millisecond)
    }
    wg.Wait()
}

var (
	count = 0 //请求数
	begin = time.Now() //计数开始时间
	cycle = time.Duration(10*time.Second) //计时周期
	rate  = 5 //计数限制，最多允许请求数目
	lock = sync.Mutex{}
)

func limit(ts time.Time) bool{
	lock.Lock()
    defer lock.Unlock()

	if count == rate-1 {
        if ts.Sub(begin) < cycle {
            return false
        } else {
			//计数周期内，重置计数器即可
    		begin = ts
    		cycle = cycle
    		count = 0
            return true
        }
    } else {
        //没有达到速率限制，计数加1
        count++
        return true
    }
}