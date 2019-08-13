package main

import (
	"fmt"
	"log"
	"time"

	"github.com/goburrow/modbus"
)

func main() {
	handler1 := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler1.BaudRate = 9600
	handler1.DataBits = 8
	handler1.Parity = "N"
	handler1.StopBits = 1
	handler1.SlaveId = 1
	handler1.Timeout = 5 * time.Second
	errin := handler1.Connect()
	if errin != nil {
		log.Println("port is busy")
	} else {
		//log.Println("succes open port")
	}
	client1 := modbus.NewClient(handler1)

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			//query
			data1, err := client1.ReadHoldingRegisters(2, 8)
			if err != nil {
				reply := fmt.Sprintf("%x", data1)
				log.Println("========================================================")
				log.Printf("reply:%s", reply)
				log.Println("========================================================")
				fmt.Println()
			}
			//query data
		}
	}
}
