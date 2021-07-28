package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type JobRunner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

//类的初始化函数 首字母大写
//对 JobRunner 类进行初始化并返回一个指向该类的指针
func New(d time.Duration) *JobRunner {
	return &JobRunner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//成员方法
func (r *JobRunner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *JobRunner) Start() error {
	// 接收系统中断信号通知
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *JobRunner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *JobRunner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
