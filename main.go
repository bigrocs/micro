package main

import (
	_ "github.com/bigrocs/micro-plugins/registry/k8s"
	"github.com/micro/micro/v2/cmd"
)

func main() {
	cmd.Init()
}
