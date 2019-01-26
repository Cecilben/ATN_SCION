//The code was developed by refering the hello world, sensorfetch app and master code of latency
// https://github.com/perrig/scionlab/blob/master/sensorapp/sensorfetcher/sensorfetcher.go
// https://github.com/netsec-ethz/scion-homeworks0/blob/master/latency/timestamp_server.go
// https://github.com/netsec-ethz/scion-apps/tree/master/helloworld
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
)

// Check just ensures the error is nil, or complains and quits
//This block has been refered from sensorfetcher
func Check(e error) {
	if e != nil {
		fmt.Println("Fatal error. Exiting.", "err", e)
	}
}

//This block has been refered from sensor fetcher
func printUsage() {
	fmt.Println(" -saddr ServerSCIONAddress -caddr ClientSCIONAddress")
	fmt.Println("The SCION address is specified as ISD-AS,[IP Address]:Port")
	fmt.Println("Example SCION address 1-1,[127.0.0.1]:42002")
}

func main() {

	var clientAddress string
	var serverAddress string
	var err error
	var client_local *snet.Addr
	var remote_server *snet.Addr
	var scion_udpconnection *snet.Conn

	//Fetches the argument from command line
	//Refered from sensor fetch
	flag.StringVar(&clientAddress, "caddr", "", "Client SCION Address")
	flag.StringVar(&serverAddress, "saddr", "", "Server SCION Address")
	flag.Parse()

	//SCION UDP socket creation
	//Refered from hello world
	client_local, err = snet.AddrFromString(clientAddress)
	Check(err)
	remote_server, err = snet.AddrFromString(serverAddress)
	Check(err)

	dpath := "/run/shm/dispatcher/default.sock"
	snet.Init(client_local.IA, sciond.GetDefaultSCIONDPath(nil), dpath)

	scion_udpconnection, err = snet.DialSCION("udp4", client_local, remote_server)
	check(err)

	receivePacketBuffer := make([]byte, 2500) //Intiating a dynamic array of respective size
	sendPacketBuffer := make([]byte, 16)      //Intiating a dynamic array of respective size

	random_seed := rand.NewSource(time.Now().Unix())

	var total int64 = 0
	gen_id := rand.New(random_seed).Uint64()         //generating a random ID using rand function
	n := binary.PutUvarint(sendPacketBuffer, gen_id) //Using binary encoder to add the random number
	sendPacketBuffer[n] = 0

	time_sent := time.Now()
	_, err = scion_udpconnection.Write(sendPacketBuffer)
	check(err)

	_, _, err = scion_udpconnection.ReadFrom(receivePacketBuffer)
	check(err)
	time_recevied := time.now()

	return_id, n := binary.Uvarint(receivePacketBuffer)

	if return_id == gen_id {
		diff := time_received.sub(time_sent) // Difference calculated in seconds
		total = diff
	}

	var final_diff float64 = float64(total)

	fmt.Printf("\nClient Address: %s\nServer Address: %s\n", clientAddress, serverAddress)
	fmt.Println("Timestamp in Seconds")
	//Printing latency in seconds
	fmt.Printf("\tRTT: %.3fs\n", final_diff)
	fmt.Printf("\tLatency: %.3fs\n", final_diff)
}
