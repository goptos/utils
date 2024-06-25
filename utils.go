package utils

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

const DEBUG = true

func Debug(str string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(str, args...)
		return
	}
}

type Verbose struct {
	Level int
}

// var verbose = (*utils.Verbose).New(nil)
func (*Verbose) New() *Verbose {
	level, err := strconv.Atoi(os.Getenv("GOPTOS_VERBOSE"))
	if err != nil {
		level = 0
	}
	return &Verbose{Level: level}
}

func (_self *Verbose) Printf(level int, str string, args ...interface{}) {
	if _self.Level >= level {
		fmt.Printf(str, args...)
	}
}

func Assert(condition bool, msg string, skip int) {
	if !condition {
		_, file, line, _ := runtime.Caller(skip)
		panic(fmt.Sprintf("ASSERT [the following condition was not met in '%s' at line %d]: %s", file, line, msg))
	}
}

// Return items of array from index i to index j (inclusive of i and j).
func Pick[T any](a []T, i int, j int) []T {
	Assert(i >= 0, fmt.Sprintf("i >= 0 when calling 'func pick(a, i, j)', i:%d j:%d", i, j), 2)
	Assert(i <= j, fmt.Sprintf("i <= j when calling 'func pick(a, i, j)', i:%d j:%d", i, j), 2)
	Assert(j < len(a), fmt.Sprintf("j < len(a) when calling 'func pick(a, i, j)', i:%d j:%d", i, j), 2)
	return a[i : j+1]
}
