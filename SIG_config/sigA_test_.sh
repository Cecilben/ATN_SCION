#!/bin/bash

sudo modprobe dummy
sudo ip link set name dummy11 dev dummy0
sudo ip addr add 172.16.0.11/32 brd + dev dummy11 label dummy11:0
sudo ip rule add to 172.16.12.0/24 lookup 11 prio 11
$SC/bin/sig -config=/home/ubuntu/go/src/github.com/scionproto/scion/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigA.config > $SC/logs/sig${IA}-1.log 2>&1 &
sudo ip link add client type dummy
sudo ip addr add 172.16.11.1/24 brd + dev client label client:0
