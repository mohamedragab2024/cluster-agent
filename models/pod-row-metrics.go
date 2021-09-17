package models

type PodRowMetrics struct {
	Name          string `json:"name"`
	CpuUsageCores int64  `json:"cpuUsageCores"`
	MemoryUsage   int64  `json:"memoryUsage"`
}
