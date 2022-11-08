package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const totalNumberOfDigits = 32

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("input IP address with mask")
		return
	}

	ipWithMask := arguments[1]
	input := strings.Split(ipWithMask, "/")

	ip := iPv4ToUint32(input[0])
	mask := input[1]

	digitMask, err := strconv.Atoi(mask)
	if err != nil {
		fmt.Println("failed to convert mask to number")
		return
	}

	countZeros := totalNumberOfDigits - digitMask

	countIPs := math.Pow(2, float64(countZeros))

	subnetMaskBinary := ""
	for i := 0; i < totalNumberOfDigits; i++ {
		if i < digitMask {
			subnetMaskBinary += "1"
			continue
		}

		subnetMaskBinary += "0"
	}

	subnetMaskDecimal, err := strconv.ParseInt(subnetMaskBinary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	firstIPAddress := ip & uint32(subnetMaskDecimal)
	lastIPAddress := ip | ^ uint32(subnetMaskDecimal)

	fmt.Println("Count IPs:", countIPs)
	fmt.Println("First IP address:", uInt32ToIPv4(firstIPAddress))
	fmt.Println("Last IP address:", uInt32ToIPv4(lastIPAddress))
}

func iPv4ToUint32(iPv4 string) uint32 {
	ipOctets := [4]uint64{}

	for i, v := range strings.SplitN(iPv4, ".", 4) {
		ipOctets[i], _ = strconv.ParseUint(v, 10, 32)
	}

	result := (ipOctets[0] << 24) | (ipOctets[1] << 16) | (ipOctets[2] << 8) | ipOctets[3]

	return uint32(result)
}

func uInt32ToIPv4(iPuInt32 uint32) (iP string) {
	iP = fmt.Sprintf("%d.%d.%d.%d",
		iPuInt32>>24,
		(iPuInt32&0x00FFFFFF)>>16,
		(iPuInt32&0x0000FFFF)>>8,
		iPuInt32&0x000000FF)
	return iP
}
