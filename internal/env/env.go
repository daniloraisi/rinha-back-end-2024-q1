package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	ErrRequired              = fmt.Errorf("variável de ambiente requerida, mas em branco")
	ErrInvalidDurationString = fmt.Errorf("variável de ambiente deve ser um `time.Duration` válido")
	ErrNotBool               = fmt.Errorf("variável de ambiente deve ser uma expressão boolean")
	ErrNotUint               = fmt.Errorf("variável de ambiente deve ser UINT")
)

type Env struct {
	name  string
	value string
	err   error
}

func (e Env) String() (string, error) {
	return e.value, e.err
}

func (e Env) TimeDuration() (time.Duration, error) {
	var d time.Duration
	if e.err != nil {
		return d, e.err
	}

	d, err := time.ParseDuration(e.value)
	if err != nil {
		return d, fmt.Errorf("%s: %w: %s", e.name, ErrInvalidDurationString, err.Error())
	}

	return d, nil
}

func (e Env) Bool() (bool, error) {
	var b bool
	if e.err != nil {
		return b, e.err
	}

	b, err := strconv.ParseBool(e.value)
	if err != nil {
		return b, fmt.Errorf("%s: %w", e.name, ErrNotBool)
	}

	return b, nil
}

func (e Env) Uint(base, bitSize int) (uint64, error) {
	var i uint64
	if e.err != nil {
		return i, e.err
	}

	i, err := strconv.ParseUint(e.value, base, bitSize)
	if err != nil {
		return i, fmt.Errorf("%s: %w", e.name, ErrNotUint)
	}

	return i, nil
}

func (e Env) Required() Env {
	if e.value == "" {
		e.err = fmt.Errorf("%s: %w", e.name, ErrRequired)
	}

	return e
}

func (e Env) Default(s string) Env {
	if e.value == "" {
		e.value = s
	}

	return e
}

func Get(name string) Env {
	return Env{
		name:  name,
		value: os.Getenv(name),
	}
}

func Panic[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
