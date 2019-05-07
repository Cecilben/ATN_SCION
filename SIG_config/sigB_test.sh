#!/bin/bash

sudo modprobe dummy
sudo ip link add dummy12 type dummy
sudo ip addr add 172.16.0.12/32 brd + dev dummy12 label dummy12:0
sudo ip rule add to 172.16.11.0/24 lookup 12 prio 12
$SC/bin/sig -config=/home/ubuntu/go/src/github.com/scionproto/scion/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigB.config > $SC/logs/sig${IA}-1.log 2>&1 &
sudo ip link add server type dummy
sudo ip addr add 172.16.12.1/24 brd + dev server label server:0

mkdir $SC/WWW
echo "Hello World!" > $SC/WWW/hello.html
cd $SC/WWW/ && python3 -m http.server --bind 172.16.12.1 8081 &
