package cronjob

import (
	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCronExpr(t *testing.T) {
	expr := cron.New(cron.WithSeconds())

	id, err := expr.AddFunc("@every 1s", func() {
		t.Log("任务执行")
	})

	assert.NoError(t, err)
	t.Log("任务 id", id)

	expr.Start()
	time.Sleep(5 * time.Second)
	// 注意：不代表整个调度立马停止，而是不再进行新任务的调度，已调度的任务会让其执行完毕
	ctx := expr.Stop()

	// 正在执行的任务执行完毕后会发出信号
	<-ctx.Done()
	t.Log("已经没有新的任务在执行")
}

type JobFunc func()

func (j JobFunc) Run() {
	j()
}
