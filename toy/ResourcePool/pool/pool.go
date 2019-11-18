package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个goroutine间共享的资源，该资源需实现io.Closer接口
type Pool struct {
	m         sync.Mutex                //保证多个goroutine访问资源池时，池内的值是安全的。
	resources chan io.Closer            //io.Closer接口类型的资源
	factory   func() (io.Closer, error) //工厂函数，当池需要资源时创建新的资源，需使用者实现
	closed    bool                      //标志资源池是否关闭
}

// ErrPoolClosed 表示请求一个已经关闭的资源池
var ErrPoolClosed = errors.New("Pool has been closed")

// New 创建一个资源池，需规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将资源归还池中
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 销毁资源池
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	// 清理资源之前关闭通道，否则死锁
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
