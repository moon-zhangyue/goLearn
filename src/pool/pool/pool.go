package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// 定义资源池结构体
type Pool struct {
	// 通过锁机制确保资源池的并发安全
	m sync.Mutex
	// 通过缓冲通道管理资源池，资源池大小即缓冲值
	resources chan io.Closer
	// 在资源池中注册新的资源
	factory func() (io.Closer, error)
	// 标识资源池是否关闭
	closed bool
}

var ErrPoolClosed = errors.New("资源池已关闭")

// 初始化资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("资源池容量需要大于0")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// 从资源池申请资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		// 资源池不为空则从中获取资源
		log.Println("Acquire:", "共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		// 资源池为空则调用 p.factory() 方法注册新资源
		log.Println("Acquire:", "新增资源")
		return p.factory()
	}
}

// 资源使用完成后释放
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	// 资源池已关闭则支持释放资源
	if p.closed {
		_ = r.Close()
		return
	}

	// 否则将资源归还到资源池
	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		_ = r.Close()
	}
}

// 关闭资源池
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	// 释放资源池和资源池中所有资源
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		_ = r.Close()
	}
}
