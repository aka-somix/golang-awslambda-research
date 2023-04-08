package main

import (
	"fmt"
	"sync"
	"runtime"
	"github.com/aws/aws-lambda-go/lambda"
	"bufio"
	"os"
)

// Utiity functon for translate from Bytes to MegaBytes
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func Handler(event any) (string, error) {
	fmt.Printf("Event Received: %s \n", event)
	
	f, err := os.OpenFile("/tmp/test", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return "Error while creating file", nil
	}

	var wg sync.WaitGroup
	bytesWriter := bufio.NewWriter(f)
	mutex := sync.Mutex{}
	
	count := 0

	// This infinite loop will eventually saturate all the memory
	for {
		wg.Add(1)

		count++

		go func(id int) {
			defer wg.Done()
    	mutex.Lock()
    	defer mutex.Unlock()

				// Check and print current allocation
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				fmt.Printf("[%d] - Alloc = %v MiB", id, bToMb(m.Alloc))
				fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
				fmt.Printf("\tNumGC = %v\n", m.NumGC)
				
				bytesWriter.Write([]byte{91,92,93,94,95,96,97, 98, 99, 100, 101, 102,103,104,105,106}) 
		}(count)		
	}

	return "Memory did not broke :(", nil
}

func main() {
        lambda.Start(Handler)
}
