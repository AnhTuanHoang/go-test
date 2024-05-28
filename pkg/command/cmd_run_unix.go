//go:build !windows
// +build !windows

package command

import (
	"github.com/kballard/go-shellquote"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const (
	restartPause = 5 * time.Second
)

func (e *Cmd) runInner() (int, bool) {
	cmdparts, err := shellquote.Split(e.cmdstr)
	if err != nil {
		return 0, true
	}

	cmd := exec.Command(cmdparts[0], cmdparts[1:]...)

	cmd.Env = append([]string(nil), os.Environ()...)

	cmd.Stdout = os.Stdout
	stdErr, _ := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		return 0, true
	}
	e.CommandBase = cmd
	cmdDone := make(chan int)
	go func() {
		cmdDone <- func() int {
			slurp, errorRead := io.ReadAll(stdErr)
			if errorRead == nil {
				e.err = slurp
			}
			err := cmd.Wait()
			if err == nil {
				return 0
			}
			ee, ok := err.(*exec.ExitError)
			if !ok {
				return 0
			}
			return ee.ExitCode()
		}()
	}()

	select {
	case <-e.terminate:
		syscall.Kill(cmd.Process.Pid, syscall.SIGINT)
		<-cmdDone
		return 0, false

	case c := <-cmdDone:
		return c, true
	}
}
