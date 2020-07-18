# Barcode keyboard emulator
Very small web service for use with a data collection terminal, emulate a keyboard in the any OS.

**Use:**
1. Run the utility on the host system and set focus to the window where you want to transfer the received barcode  
2. Open a web browser on the data collection terminal and go to address http://host_ip:8080  
3. Scan code in one of two fields  
  - "BarCode (Tap Enter)" (DEFAULT focus) - emulated tap "Enter" after input barcode data  
  - "BarCode only" - only barcode data  
  
**Build:**
```sh
go mod download
go build -ldflags "-s -w" -o build/
```