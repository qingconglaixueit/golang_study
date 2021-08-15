package main

import "github.com/wonderivan/logger"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		logger.Info("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	logger.Info("Successfully got connection")
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	logger.Info("New connection coming: %d", c)
}
