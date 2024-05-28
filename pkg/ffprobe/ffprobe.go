package ffprobe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)
func GetFileContent(path string) (data *ProbeData, err error) {
	args := append([]string{
		"-loglevel", "error",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
	})
	args = append(args, path)

	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFn()
	cmd := exec.CommandContext(ctx, ffprobeBinPath, args...)
	cmd.SysProcAttr = nil
	return runProbe(cmd)
}

func runProbe(cmd *exec.Cmd) (data *ProbeData, err error) {
	var outputBuf bytes.Buffer
	var stdErr bytes.Buffer

	cmd.Stdout = &outputBuf
	cmd.Stderr = &stdErr

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running %s [%s] %w", "ffprobe", stdErr.String(), err)
	}

	if stdErr.Len() > 0 {
		return nil, fmt.Errorf("ffprobe error: %s", stdErr.String())
	}

	data = &ProbeData{}
	err = json.Unmarshal(outputBuf.Bytes(), data)
	if err != nil {
		return data, fmt.Errorf("error parsing ffprobe output: %w", err)
	}

	if data.Format == nil {
		return data, fmt.Errorf("no format data found in ffprobe output")
	}

	return data, nil
}

