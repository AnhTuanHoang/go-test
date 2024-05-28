package command

import (
	"sync"
)

type Pool struct {
	Wg sync.WaitGroup
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) Close() {
	p.Wg.Wait()
}
