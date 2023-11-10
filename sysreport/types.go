package main

type ProcInfo struct {
	FDCount  int    `json:"fd_count"`
	VmRSS    string `json:"vm_rss"`
	ExecPath string `json:"exec_path"`
}

type CPUInfo struct {
	CPUCount  int    `json:"cpu_count"`
	ModelName string `json:"model_name"`
}

type MemInfo struct {
	MemTotal     string `json:"mem_total"`
	MemFree      string `json:"mem_free"`
	MemAvailable string `json:"mem_available"`
	Buffers      string `json:"buffers"`
	Cached       string `json:"cached"`
	SwapTotal    string `json:"swap_total"`
	SwapCached   string `json:"swap_cached"`
	SwapFree     string `json:"swap_free"`
}

type Report struct {
	ProcInfo ProcInfo `json:"process_info"`
	CPUInfo  CPUInfo  `json:"cpu_info"`
	MemInfo  MemInfo  `json:"memory_info"`
}
