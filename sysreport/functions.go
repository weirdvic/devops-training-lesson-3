package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// getProcInfo retrieves information about a process given its PID.
//
// pid: the process ID.
// Returns a ProcInfo struct containing information about the process.
func getProcInfo(pid int) ProcInfo {
	// Initialize ProcInfo struct
	var procInfo ProcInfo

	// Get file descriptors directory path
	fdDir := fmt.Sprintf("/proc/%d/fd", pid)

	// Read the directory to get file descriptors
	files, err := os.ReadDir(fdDir)
	if err != nil {
		log.Println(err)
		procInfo.FDCount = 0
	}

	// Set the number of file descriptors
	procInfo.FDCount = len(files)

	// Get the process status file path
	statusFilePath := fmt.Sprintf("/proc/%d/status", pid)

	// Read the process status file content
	statusContent, err := os.ReadFile(statusFilePath)
	if err != nil {
		log.Println(err)
		procInfo.VmRSS = "unknown"
	}

	// Parse the process status file content to get VmRSS
	for _, line := range strings.Split(string(statusContent), "\n") {
		if strings.HasPrefix(line, "VmRSS:") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				procInfo.VmRSS = parts[1] + " kB"
			}
		} else {
			procInfo.VmRSS = "unknown"
		}
	}

	// Get the executable path
	procInfo.ExecPath, err = os.Readlink(fmt.Sprintf("/proc/%d/exe", pid))
	if err != nil {
		log.Println(err)
		procInfo.ExecPath = "unknown"
	}

	// Return the ProcInfo struct
	return procInfo
}

// getCPUInfo retrieves CPU information from the /proc/cpuinfo file.
//
// It returns a CPUInfo struct containing the model name and the number of CPU cores.
func getCPUInfo() CPUInfo {
	// Open the /proc/cpuinfo file
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Println("Error opening /proc/cpuinfo:", err)
		return CPUInfo{}
	}
	defer file.Close()

	// Create a CPUInfo struct to store the data
	var cpuInfo CPUInfo

	// Create a map to store key-value pairs
	cpuInfoMap := make(map[string]string)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and parse the /proc/cpuinfo file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			cpuInfoMap[key] = value
		}
	}

	// Extract specific fields of interest
	if model, exists := cpuInfoMap["model name"]; exists {
		cpuInfo.ModelName = model
	}

	// Calculate the number of CPU cores
	lastProcessorID, err := strconv.Atoi(cpuInfoMap["processor"])
	if err != nil {
		log.Fatal(err)
	}
	cpuInfo.CPUCount = lastProcessorID + 1
	return cpuInfo
}

// getMemInfo returns the memory information of the system.
//
// It opens the /proc/meminfo file and reads its contents to extract
// specific fields of interest, such as MemTotal, MemFree, MemAvailable,
// Buffers, Cached, SwapTotal, SwapCached, and SwapFree. It returns a
// MemInfo struct containing the extracted information.
//
// Returns:
// - MemInfo: a struct containing the extracted memory information.
func getMemInfo() MemInfo {
	// Open the /proc/meminfo file
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Println("Error opening /proc/meminfo:", err)
		return MemInfo{}
	}
	defer file.Close()

	// Create a MemInfo struct to store the data
	var memInfo MemInfo

	// Create a map to store key-value pairs
	memInfoMap := make(map[string]string)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and parse the /proc/meminfo file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			memInfoMap[key] = value
		}
	}

	// Extract specific fields of interest
	if mem_total, exists := memInfoMap["MemTotal"]; exists {
		memInfo.MemTotal = mem_total
	}
	if mem_free, exists := memInfoMap["MemFree"]; exists {
		memInfo.MemFree = mem_free
	}
	if mem_available, exists := memInfoMap["MemAvailable"]; exists {
		memInfo.MemAvailable = mem_available
	}
	if buffers, exists := memInfoMap["Buffers"]; exists {
		memInfo.Buffers = buffers
	}
	if cached, exists := memInfoMap["Cached"]; exists {
		memInfo.Cached = cached
	}
	if swap_total, exists := memInfoMap["SwapTotal"]; exists {
		memInfo.SwapTotal = swap_total
	}
	if swap_cached, exists := memInfoMap["SwapCached"]; exists {
		memInfo.SwapCached = swap_cached
	}
	if swap_free, exists := memInfoMap["SwapFree"]; exists {
		memInfo.SwapFree = swap_free
	}

	return memInfo
}
