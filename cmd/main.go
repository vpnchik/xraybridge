package main

import (
	"fmt"
	"os"

	"github.com/vpnchik/xraybridge"
)

func main() {
	data, err := os.ReadFile("conf.json")
	if err != nil {
		fmt.Println(err)
	}
	xray := xraybridge.NewXrayInstance()

	err = xray.Run(data)
	sig := make(chan os.Signal, 1)
	<-sig
	xray.Stop()
}
