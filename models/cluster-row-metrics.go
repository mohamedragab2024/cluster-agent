package models

type ClusterMetricsCache struct {
	TotalCpuCores    string `json:"totalCpuCores"`
	TotalCpuUsage    string `json:"totalCpuUsage"`
	TotalMemory      string `json:"totalMemory"`
	TotalMemoryUsage string `json:"totalMemoryUsage"`
	CpuPercentage    string `json:"cpuPercentage"`
	MemoryPercentage string `json:"memoryPercentage"`
	NodesCount       int64  `json:"nodesCount"`
	Provider         string `json:"provider"`
}
