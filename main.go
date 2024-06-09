package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	h, _ := host.Info()

	fmt.Printf("OS: %s, Platform: %s, Platform Version: %s\n", h.OS, h.Platform, h.PlatformVersion)

	c, _ := cpu.Info()

	for _, cp := range c {
		fmt.Printf("CPU #: %d, Model Name: %s, Mhz: %f, Cores: %d\n", cp.CPU, cp.ModelName, cp.Mhz, cp.Cores)
	}
}
