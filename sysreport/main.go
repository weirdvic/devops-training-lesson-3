package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var report = Report{
		ProcInfo: getProcInfo(os.Getpid()),
		CPUInfo:  getCPUInfo(),
		MemInfo:  getMemInfo(),
	}

	reportJSON, err := json.Marshal(report)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	// Print the JSON representation
	fmt.Println(string(reportJSON))
}
