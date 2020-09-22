# GoScan 

A concurrent golang based network scanner. 

## Install Dependancies 

```bash
$ go get github.com/google/gopacket/pcap
$ go install github.com/google/gopacket/pcap
```

## Build 

```bash 
$ sudo go build -o goscan main.go network.go datagram.go
```

## Run 
```bash
$ GoScan <interface name> 
$ GoScan en0 
```

## Output 
```bash
Scanning on interface: en0
Polling for packets...
Writing ARP Requests.
Stopped polling for packets.
Host IP: 10.0.2.55 	    Type: net.IP
Host IP: 10.0.2.56 	    Type: net.IP
Host IP: 10.0.2.63      Type: net.IP
Host IP: 10.0.2.136     Type: net.IP
Host IP: 10.0.2.148     Type: net.IP
Host IP: 10.0.2.207     Type: net.IP
Host IP: 10.0.2.250     Type: net.IP
Host IP: 10.0.2.254     Type: net.IP
```