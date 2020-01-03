package limit

import (
	"github.com/gpmgo/gopm/modules/log"
	"sync"
)

//停牌桶限流器

type ConnLimiter struct {
	// 并发数
	concurrentConn int
	// 停牌桶数 每有一个链接，往channel里加一个值 每释放一个 则出一个
	bucket    chan int
	closeOnce sync.Once
}

//传入参数 并发数
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

//获取连接 返回是否还有令牌
func (l *ConnLimiter) GetConn() bool {
	if len(l.bucket) >= l.concurrentConn {
		log.Error("reached the rate limitation.")
		return false
	}
	l.bucket <- 1
	return true
}

//释放令牌
func (l *ConnLimiter) ReleaseConn() {
	<-l.bucket
	log.Info("release one connection")
}
