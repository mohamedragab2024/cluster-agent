package models

type NodeRowMetrics struct {
	Name                    string `json:"name"`
	TotalCpuCors            string `json:"totalCpuCors"`
	CpuUsageCors            string `json:"cpuUsageCors"`
	CpuUsagePercentage      string `json:"cpuUsagePercentage"`
	MemoryUsage             string `json:"memoryUsage"`
	TotalMemory             string `json:"totalMemory"`
	MemoryUsagePercentage   string `json:"memoryUsagePercentage"`
	Pods                    string `json:"pods"`
	Architecture            string `json:"architecture"`
	KubeletVersion          string `json:"kubeletVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	IpAddress               string `json:"ipAddress"`
	HostName                string `json:"hostName"`
}
