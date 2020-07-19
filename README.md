# Barcode keyboard emulator
Very small web service for use with a data collection terminal, emulate a keyboard in the any OS.

**Usage:**
```sh
# Use default params (all host ip and port 9090)
barcode-keyboard-emulator
```
```sh
# Change address and port
barcode-keyboard-emulator -addr=192.168.0.2 -port=9091
```

1. Run the utility on the host system and set focus to the window where you want to transfer the received barcode  
2. Open a web browser on the data collection terminal and go to address http://host_ip:9090 (use -addr and -port commandline parameters for change this)  
3. Scan code in one of two fields  
  - "BarCode (Tap Enter)" (DEFAULT focus) - emulated tap "Enter" after input barcode data  
  - "BarCode only" - only barcode data  


  
**Build:**
```sh
go mod download
go build -ldflags "-s -w" -o build/
```