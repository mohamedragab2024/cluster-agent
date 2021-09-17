package models

type PodRowMetrics struct {
	Name          string `json:"name"`
	CpuUsageCores string `json:"cpuUsageCores"`
	MemoryUsage   string `json:"memoryUsage"`
}
