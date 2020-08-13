package main

import "strconv"

type Expression interface {
	String() string
}

type Int struct {
	token Token
	Value int64
}

func (i Int) String() string {
	return strconv.FormatInt(i.Value, 10)
}

type Float struct {
	token Token
	Value float64
}

func (f Float) String() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}
