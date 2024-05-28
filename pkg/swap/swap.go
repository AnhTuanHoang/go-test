package swap

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var ErrFileParse = errors.New("meomeo")

type Swap struct {
	Filename string
	Type     string
	Size     int
	Used     int
	Priority int
}

func ReadFileNoStat(filename string) ([]byte, error) {
	const maxBufferSize = 1024 * 1024
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := io.LimitReader(f, maxBufferSize)
	return io.ReadAll(reader)
}
func GetSwaps(fileName string) ([]*Swap, error) {
	data, err := ReadFileNoStat(fileName)
	if err != nil {
		return nil, err
	}
	return parseSwaps(data)
}

func parseSwaps(info []byte) ([]*Swap, error) {
	var swaps []*Swap
	scanner := bufio.NewScanner(bytes.NewReader(info))
	scanner.Scan() // ignore header line
	for scanner.Scan() {
		swapString := scanner.Text()
		parsedSwap, err := parseSwapString(swapString)
		if err != nil {
			return nil, err
		}
		swaps = append(swaps, parsedSwap)
	}

	err := scanner.Err()
	return swaps, err
}

func parseSwapString(swapString string) (*Swap, error) {
	var err error
	swapFields := strings.Fields(swapString)
	swapLength := len(swapFields)
	if swapLength < 5 {
		return nil, fmt.Errorf("%w: too few fields in swap string: %s", ErrFileParse, swapString)
	}
	swap := &Swap{
		Filename: swapFields[0],
		Type:     swapFields[1],
	}
	swap.Size, err = strconv.Atoi(swapFields[2])
	if err != nil {
		return nil, fmt.Errorf("%s: invalid swap size: %s: %w", ErrFileParse, swapFields[2], err)
	}
	swap.Used, err = strconv.Atoi(swapFields[3])
	if err != nil {
		return nil, fmt.Errorf("%s: invalid swap used: %s: %w", ErrFileParse, swapFields[3], err)
	}
	swap.Priority, err = strconv.Atoi(swapFields[4])
	if err != nil {
		return nil, fmt.Errorf("%s: invalid swap priority: %s: %w", ErrFileParse, swapFields[4], err)
	}

	return swap, nil
}
