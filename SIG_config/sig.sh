#!/bin/bash

export IA=$(cat $SC/gen/ia)
export IAd=$(cat $SC/gen/ia | sed 's/_/\:/g')
export AS=$(cat $SC/gen/ia | cut --fields=2 --delimiter="-")
export ISD=$(cat $SC/gen/ia | cut --fields=1 --delimiter="-")
mkdir -p ${SC}/gen/ISD${ISD}/AS${AS}/sig${IA}-1/
go build -o ${SC}/bin/sig ${SC}/go/sig/main.go
sudo setcap cap_net_admin+eip ${SC}/bin/sig

sudo sysctl net.ipv4.conf.default.rp_filter=0
sudo sysctl net.ipv4.conf.all.rp_filter=0
sudo sysctl net.ipv4.ip_forward=1

sed -i "s/\${IA}/${IA}/g" ${SC}/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigA.config
sed -i "s/\${IAd}/${IAd}/g" ${SC}/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigA.config
sed -i "s/\${AS}/${AS}/g" ${SC}/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigA.config
sed -i "s/\${ISD}/${ISD}/g" ${SC}/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigA.config
export tunIP=$(echo $(ip a | grep "global tun") | cut --fields=2 --delimiter=" ")
sed -i "s/10.0.8.A/${tunIP}/g" ${SC}/gen/ISD${ISD}/AS${AS}/sig${IA}-1/sigA.config
