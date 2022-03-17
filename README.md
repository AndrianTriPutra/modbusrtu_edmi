# [golang] modbusrtu-powermeter EDMI

step by step
1. setting your powermeter
2. set modbus address in powermeter "3 for kWH Tot,5 for WTot,7 for WBP,9 for LWBP"
3. set slave id 1
4. connect "converter usb to RS485" to powermeter and pc
5. don't forget permission "sudo chmod a+rw /dev/ttyU* "check with dmesg"
6. install library goburrow with "go get github.com/goburrow/modbus"
note:
you must modified the library, so my file "basic.go" can work well
if you see my code, you can see how to start trace "goburrow/modbus"
7. and finally you can run my file go with "go run basic.go"

Thanks
