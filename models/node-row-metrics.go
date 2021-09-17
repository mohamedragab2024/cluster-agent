package models

type NodeRowMetrics struct {
	Name                    string `json:"name"`
	TotalCpuCores           int64  `json:"totalCpuCores"`
	CpuUsageCores           int64  `json:"cpuUsageCores"`
	CpuUsagePercentage      string `json:"cpuUsagePercentage"`
	MemoryUsage             int64  `json:"memoryUsage"`
	TotalMemory             int64  `json:"totalMemory"`
	MemoryUsagePercentage   string `json:"memoryUsagePercentage"`
	Pods                    string `json:"pods"`
	Architecture            string `json:"architecture"`
	KubeletVersion          string `json:"kubeletVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	IpAddress               string `json:"ipAddress"`
	HostName                string `json:"hostName"`
}
