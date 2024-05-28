package gocron

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Locker interface {
	Lock(key string) (bool, error)
	Unlock(key string) error
}

type timeUnit int

const MAXJOBNUM = 10000

//go:generate stringer -type=timeUnit
const (
	seconds timeUnit = iota + 1
	minutes
	hours
	days
	weeks
)

var (
	loc    = time.Local
	locker Locker
)

// ChangeLoc change default the time location
func ChangeLoc(newLocation *time.Location) {
	loc = newLocation
	defaultScheduler.ChangeLoc(newLocation)
}

// SetLocker sets a locker implementation
func SetLocker(l Locker) {
	locker = l
}

func callJobFuncWithParams(jobFunc interface{}, params []interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(jobFunc)
	if len(params) != f.Type().NumIn() {
		return nil, ErrParamsNotAdapted
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in), nil
}

// for given function fn, get the name of function.
func getFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func getFunctionKey(funcName string) string {
	h := sha256.New()
	h.Write([]byte(funcName))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Jobs returns the list of Jobs from the defaultScheduler
func Jobs() []*Job {
	return defaultScheduler.Jobs()
}

func formatTime(t string) (hour, min, sec int, err error) {
	ts := strings.Split(t, ":")
	if len(ts) < 2 || len(ts) > 3 {
		return 0, 0, 0, ErrTimeFormat
	}

	if hour, err = strconv.Atoi(ts[0]); err != nil {
		return 0, 0, 0, err
	}
	if min, err = strconv.Atoi(ts[1]); err != nil {
		return 0, 0, 0, err
	}
	if len(ts) == 3 {
		if sec, err = strconv.Atoi(ts[2]); err != nil {
			return 0, 0, 0, err
		}
	}

	if hour < 0 || hour > 23 || min < 0 || min > 59 || sec < 0 || sec > 59 {
		return 0, 0, 0, ErrTimeFormat
	}

	return hour, min, sec, nil
}

// NextTick returns a pointer to a time that will run at the next tick
func NextTick() *time.Time {
	now := time.Now().Add(time.Second)
	return &now
}
