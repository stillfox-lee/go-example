package main

import (
	"fmt"
	"time"
)

/*
embed 和 interface 结合使用
同时需要注意 BaseTask 的 receiver 使用
*/

type Task interface {
	IsTimeout() bool
}

type BaseTask struct {
	Number  int
	StartAt time.Time
}

func (t *BaseTask) IsTimeout() bool {
	return time.Now().After(t.StartAt)
}

type MapTask struct {
	BaseTask
	SourceFilename string
	InterFilesname map[int]string
}

func (t *MapTask) MapTaskFun() {
	fmt.Println("MapTaskFun")
}

func testIsTimeout(t Task) {
	if t.IsTimeout() {
		fmt.Println("timeout")
	}
}

func main() {
	task := &MapTask{
		BaseTask: BaseTask{
			Number:  1,
			StartAt: time.Now(),
		},
		SourceFilename: "testTemp.txt",
	}
	valueTask := MapTask{
		BaseTask: BaseTask{
			Number:  2,
			StartAt: time.Now(),
		},
		SourceFilename: "testTemp.txt",
	}
	testIsTimeout(task)
	testIsTimeout(&valueTask)

	task.MapTaskFun()
}
