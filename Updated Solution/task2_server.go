//The code was developed by refering the hello world, sensorfetch app and master code of latency
// https://github.com/perrig/scionlab/blob/master/sensorapp/sensorfetcher/sensorfetcher.go
// https://github.com/netsec-ethz/scion-homeworks0/blob/master/latency/timestamp_server.go
// https://github.com/netsec-ethz/scion-apps/tree/master/helloworld

package main

import (
	"flag"
	"fmt"

	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
)

//This block has been refered from sensor fetch
func Check(e error) {
	if e != nil {
		fmt.Println("Fatal error. Exiting.", "err", e)
	}
}

//This block has been refered from sensor fetch
func printUsage() {
	fmt.Println("\ntimestamp_server -saddr ServerSCIONAddress")
	fmt.Println("\tListens for incoming connections and responds back to them right away")
	fmt.Println("\tExample SCION address 1-1,[127.0.0.1]:42002\n")
}

func main() {

	var serverAddress string
	var err error
	var server *snet.Addr
	var scion_udpconnection *snet.Conn

	// Fetch arguments from command line
	flag.StringVar(&serverAddress, "saddr", "", "Server Address")
	flag.Parse()

	serverAddress, err = snet.AddrFromString(serverAddress)
	check(err)

	dpath := "/run/shm/dispatcher/default.sock"
	snet.Init(server.IA, sciond.GetDefaultSCIONDPath(nil), dpath)

	scion_udpconnection, err = snet.ListenSCION("udp4", server) //Listens for UDP connection to respond
	check(err)

	receivePacketBuffer := make([]byte, 2500)

	for {
		a, clientAddress, err := scion_udpconnection.ReadFrom(receivePacketBuffer)
		check(err)
		//time_recieved := time.Now().Unix()                                             // Time of recipte is stored
		//b := binary.PutVarint(receivePacketBuffer[b:], time_recieved)                  //encoding the time of reciept to packet
		_, err = scion_udpconnection.WriteTo(receivePacketBuffer[:a], clientAddress) //appending the recieved packet with the time stamp
		check(err)

		fmt.Println("Received Scion connection from", clientAddress)
	}
}
