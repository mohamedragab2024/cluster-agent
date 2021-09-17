package models

type ClusterMetricsCache struct {
	TotalCpuCores    int64  `json:"totalCpuCores"`
	TotalCpuUsage    int64  `json:"totalCpuUsage"`
	TotalMemory      int64  `json:"totalMemory"`
	TotalMemoryUsage int64  `json:"totalMemoryUsage"`
	CpuPercentage    string `json:"cpuPercentage"`
	MemoryPercentage string `json:"memoryPercentage"`
	NodesCount       int64  `json:"nodesCount"`
	Provider         string `json:"provider"`
}
