/*

   Purpose:  This is a layer2 lan arp scanning tool.

   Author:   Matthew Rogers
   Date:     2020-05-17
   License:  GPLv2

   Revision:

*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/j-keck/arping"
)

func main() {

	//take arguments
	flag.Parse()
	scanCidr := flag.Arg(0)

	//check for bad CIDR
	_, ipv4Net, err := net.ParseCIDR(scanCidr)
	if err != nil {
		fmt.Println("Provide a valid CIDR Subnet as an argument")
		fmt.Println("Example: go_network_scanner 192.168.0.1/24")
		log.Fatal(err)
	}
	startTime := time.Now()
	//fmt.Println(ipv4Addr)
	fmt.Println("Net:", ipv4Net)
	startAddress, endAddress := cidr.AddressRange(ipv4Net)
	fmt.Println("Start/End:", startAddress, endAddress)
	fmt.Println("Count:", cidr.AddressCount(ipv4Net))

	//Go Through IP Address Range
	for ipAddress := startAddress; bytes.Compare(ipAddress, cidr.Inc(endAddress)) != 0; ipAddress = cidr.Inc(ipAddress) {

		arping.SetTimeout(100 * time.Millisecond)

		macAddress, _, _ := arping.Ping(ipAddress)

		if len(macAddress) > 0 {
			fmt.Println(ipAddress, macAddress)
		}

	}
	finishedTime := time.Now()
	fmt.Println("Total Time:", finishedTime.Sub(startTime))
}
