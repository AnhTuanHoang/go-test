package ffprobe

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrTagNotFound = errors.New("tag not found")

type Tags map[string]interface{}

func (t Tags) GetInt(tag string) (int64, error) {
	v, found := t[tag]
	if !found || v == nil {
		return 0, ErrTagNotFound
	}

	switch v := v.(type) {
	case string:
		return valToInt64(v)
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	}

	str := fmt.Sprintf("%v", v)
	return valToInt64(str)
}

func (t Tags) GetString(tag string) (string, error) {
	v, found := t[tag]
	if !found || v == nil {
		return "", ErrTagNotFound
	}
	return valToString(v), nil
}

func (t Tags) GetFloat(tag string) (float64, error) {
	v, found := t[tag]
	if !found || v == nil {
		return 0, ErrTagNotFound
	}

	switch v := v.(type) {
	case string:
		return valToFloat64(v)
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	}

	str := fmt.Sprintf("%v", v)
	return valToFloat64(str)
}

func valToInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("int64 parsing error (%v): %w", str, err)
	}
	return val, nil
}

func valToString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int64:
		return strconv.FormatInt(v, 10)
	}

	return fmt.Sprintf("%v", v)
}

func valToFloat64(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("float64 parsing error (%v): %w", str, err)
	}
	return val, nil
}
