package command

import (
	"os/exec"
	"time"
)

// Cmd is an external command.
type Cmd struct {
	pool    *Pool
	cmdstr  string
	restart bool
	onExit  func(int)

	// in
	terminate chan struct{}

	isClosed bool

	err []byte

	CommandBase *exec.Cmd
}

func NewCmd(
	pool *Pool,
	cmdstr string,
	restart bool,
	onExit func(int),
) *Cmd {

	e := &Cmd{
		pool:      pool,
		cmdstr:    cmdstr,
		restart:   restart,
		onExit:    onExit,
		terminate: make(chan struct{}),
		isClosed:  false,
	}

	pool.Wg.Add(1)

	go e.run()

	return e
}

func (e *Cmd) Close() {
	close(e.terminate)
}

func (e *Cmd) IsClosed() bool {
	return e.isClosed
}

func (e *Cmd) GetErr() []byte {
	return e.err
}

func (e *Cmd) run() {
	defer e.pool.Wg.Done()

	for {
		ok := func() bool {
			c, ok := e.runInner()
			if !ok {
				return false
			}

			e.onExit(c)

			if !e.restart {
				<-e.terminate
				return false
			}

			select {
			case <-time.After(restartPause):
				return true
			case <-e.terminate:
				return false
			}
		}()
		if !ok {
			e.isClosed = true
			break
		}
	}
}
