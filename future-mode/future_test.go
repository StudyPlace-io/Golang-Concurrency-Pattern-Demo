package future_mode

import (
	"fmt"
	"testing"
	"time"
)

/*
 模式：如果串行调用，会耗时很长，这时就可以使用future模式来进行开发。
*/

//  future 模式
/*
* 1. 使用chan作为函数参数
* 2. 启动goroutine调用函数
* 3. 通过chan传入参数
* 4. 做其他可以并行处理的事情
* 5. 通过chan异步获取结果
 */

type query struct {
	// 查询参数 chan
	sql chan string
	// 接收结果参数
	result chan string
}

func newQuery() *query {
	return &query{sql: make(chan string, 2), result: make(chan string, 2)}
}

// execQuery 执行查询db任务
func execQuery(q *query) {
	go func() {
		for {
			queryCmd := <-q.sql
			fmt.Println("query db，cost a lot of time...")
			time.Sleep(time.Second * 3)
			q.result <- "result from " + queryCmd
		}

	}()
}

func TestFutureMode(test *testing.T) {

	q := newQuery()

	go execQuery(q)
	q.sql <- "select * from table"
	time.Sleep(3 * time.Second)

	fmt.Println("do something.....")

	fmt.Println(<-q.result)
	q.sql <- "select * from table aaa "
	time.Sleep(3 * time.Second)
	fmt.Println(<-q.result)
}
