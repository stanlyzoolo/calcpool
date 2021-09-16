package main

import (
	"github.com/stanlyzoolo/exprgen"
)

func main() {
	fmt.Println(prepareJob())

}

type jobMetadata map[string]interface{}

type Job struct {
	workerID int
	execFn   string
}

func prepareJob() string {
	return exprgen.Generate(10)
}

