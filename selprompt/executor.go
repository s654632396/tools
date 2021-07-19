package selprompt

import (
	"fmt"
	"log"
	"os"
)

type ExecutorSession struct {
	// something todo
}

func (es *ExecutorSession) Executor(cmd string) {

	log.Println("execute CMD:", cmd)
	if cmd == "exit" || cmd == "quit" {
		fmt.Println("Bye~")
		os.Exit(0)
	}
}

func NewExecutorSession() *ExecutorSession {
	return &ExecutorSession{}
}
