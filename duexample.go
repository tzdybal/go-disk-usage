package main

import (
	"fmt"

	"github.com/tzdybal/go-disk-usage/du"
)

// KiB is a number of bytes in Kibibyte
const KiB = uint64(1024)

func main() {
	usage := du.NewDiskUsage("/")
	fmt.Println("Free:", usage.Free()/(KiB*KiB))
	fmt.Println("Available:", usage.Available()/(KiB*KiB))
	fmt.Println("Size:", usage.Size()/(KiB*KiB))
	fmt.Println("Used:", usage.Used()/(KiB*KiB))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
