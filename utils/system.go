package utils

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	B           = 1
	KB          = 1024 * B
	MB          = 1024 * KB
	GB          = 1024 * MB
	nanoseconds = 1e9
)

type SystemInfo struct {
	OsInfo   OsInfo   `json:"os_info"`
	CpuInfo  CpuInfo  `json:"cpu_info"`
	RamInfo  RamInfo  `json:"ram_info"`
	DiskInfo DiskInfo `json:"disk_info"`
}

type OsInfo struct {
	GOOS         string `json:"type"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutine"`
}

type CpuInfo struct {
	UsedPercent  float64 `json:"used_percent"`
	Cores        int     `json:"cores"`
	LogicalCores int     `json:"logical_cores"`
}

type RamInfo struct {
	UsedMB      int `json:"used_mb"`
	TotalMB     int `json:"total_mb"`
	UsedPercent int `json:"used_percent"`
}

type DiskInfo struct {
	UsedMB      int `json:"used_mb"`
	UsedGB      int `json:"used_gb"`
	TotalMB     int `json:"total_mb"`
	TotalGB     int `json:"total_gb"`
	UsedPercent int `json:"used_percent"`
}

// GetHardwareInfo 获取运行所在服务器的配置及资源消耗
func GetHardwareInfo() (info SystemInfo, err error) {
	info.OsInfo = getOsInfo()
	info.CpuInfo, err = getCpuInfo()
	if err != nil {
		return
	}
	info.RamInfo, err = getRamInfo()
	if err != nil {
		return
	}
	info.DiskInfo, err = getDiskInfo()
	if err != nil {
		return
	}
	return
}

// GetContainerInfo 获取运行所在容器的配置及资源消耗
func GetContainerInfo() (info SystemInfo, err error) {
	info.OsInfo = getOsInfo()
	info.CpuInfo, err = getCgroupCpuInfo()
	if err != nil {
		return
	}
	info.RamInfo, err = getCgroupRamInfo()
	if err != nil {
		return
	}
	info.DiskInfo, err = getDiskInfo()
	if err != nil {
		return
	}
	return
}

// IsContainer 判断当前运行环境是否是容器
func IsContainer() (b bool) {
	_, err := os.Stat("/.dockerenv")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 获取运行环境操作系统的详细信息
func getOsInfo() (o OsInfo) {
	o.GOOS = runtime.GOOS
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

// 获取CPU资源占用的详细信息
func getCpuInfo() (c CpuInfo, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cores, err := cpu.Counts(true); err != nil {
		return c, err
	} else {
		c.LogicalCores = cores
	}
	if percents, err := cpu.Percent(time.Duration(200)*time.Millisecond, false); err != nil {
		return c, err
	} else {
		c.UsedPercent = percents[0]
	}
	return c, err
}

// 获取 Cgroup CPU资源占用的详细信息
func getCgroupCpuInfo() (c CpuInfo, err error) {
	if cpuacctUsage, err := ReadLines("/sys/fs/cgroup/cpuacct/cpuacct.usage"); err != nil {
		return c, err
	} else {
		if usage, err := strconv.ParseFloat(cpuacctUsage[0], 64); err != nil {
			return c, err
		} else {
			c.UsedPercent = usage / nanoseconds
		}
	}
	return c, err
}

// 获取内存资源占用的详细信息
func getRamInfo() (r RamInfo, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, err
}

// 获取 Cgroup 内存资源占用的详细信息
func getCgroupRamInfo() (r RamInfo, err error) {
	if usageInBytes, err := ReadLines("/sys/fs/cgroup/memory/memory.usage_in_bytes"); err != nil {
		return r, err
	} else {
		if used, err := strconv.Atoi(usageInBytes[0]); err != nil {
			return r, err
		} else {
			r.UsedMB = used / MB
		}
	}
	if limitInBytes, err := ReadLines("/sys/fs/cgroup/memory/memory.limit_in_bytes"); err != nil {
		return r, err
	} else {
		if total, err := strconv.Atoi(limitInBytes[0]); err != nil {
			return r, err
		} else {
			r.TotalMB = total / MB
		}
	}
	r.UsedPercent = r.UsedMB / r.TotalMB
	return r, err
}

// 获取硬盘资源占用的详细信息
func getDiskInfo() (d DiskInfo, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, err
}
