package goenvvar

import (
	"fmt"
	"os"
	"strconv"
)

type Env struct {
	key   string
	value string
	ok    bool
}

func Get(key string) *Env {
	env := Env{}

	val, ok := os.LookupEnv(key)

	env.key = key
	env.value = val
	env.ok = ok

	return &env
}

func (e *Env) Required() *Env {
	if e.key == "" {
		panic("Get() should be used before Required()")
	}

	if !e.ok {
		panic(fmt.Sprintf("%s: Variable not set \n", e.key))
	}

	return e
}

func (e *Env) DefaultValue(value string) *Env {
	if !e.ok {
		e.value = value
	}

	return e
}

func (e *Env) AsInteger() int {
	val, err := strconv.Atoi(e.value)

	if err != nil {
		panic(fmt.Sprintf("%s: Cannot convert to integer", e.key))
	}

	return val
}

func (e *Env) AsPortNumber() int {
	val := e.AsInteger()

	if val > 65535 {
		panic(fmt.Sprintf("%s: Cannot assign a port number greater than 65535", e.key))
	}

	return val
}

func (e *Env) AsBoolean() bool {
	if e.value == "true" || e.value == "1" {
		return true
	} else if e.value == "false" || e.value == "0" {
		return false
	} else {
		panic(fmt.Sprintf("%s: Cannot convert to boolean", e.key))
	}
}

func (e *Env) AsString() string {
	return e.value
}
