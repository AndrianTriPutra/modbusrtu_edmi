# go-modbusRTU-powermeter

step by step
1. setting your powermeter
2. i set modbus address in powermeter "3 for kWH Tot,5 for WTot,7 for WBP,9 for LWBP"
3. i set slave id 1
4. connecting "converter usb to RS485" to powermeter and pc
5. don't forget permission "sudo chmod a+rw /dev/ttyU* "check with dmesg"
6. go run *go
7. done
