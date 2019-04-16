package funcs

import (
	"fmt"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/toolkits/nux"
)

/*
add by wch 2019-03-28: aiops
新增多核CPU展示逻辑
*/

func deltaTotalSinglecpu(cpu, oldcpu *nux.CpuUsage) uint64 {
	if procStatHistory[1] == nil {
		return 0
	}
	return cpu.Total - oldcpu.Total
}

func CpuIdleSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.Idle-oldcpu.Idle) * invQuotient
}

func CpuUserSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.User-oldcpu.User) * invQuotient
}

func CpuNiceSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.Nice-oldcpu.Nice) * invQuotient
}

func CpuSystemSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.System-oldcpu.System) * invQuotient
}

func CpuIowaitSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.Iowait-oldcpu.Iowait) * invQuotient
}

func CpuIrqSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.Irq-oldcpu.Irq) * invQuotient
}

func CpuSoftIrqSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.SoftIrq-oldcpu.SoftIrq) * invQuotient
}

func CpuStealSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.Steal-oldcpu.Steal) * invQuotient
}

func CpuGuestSinglecpu(cpu, oldcpu *nux.CpuUsage) float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotalSinglecpu(cpu, oldcpu)
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(cpu.Guest-oldcpu.Guest) * invQuotient
}

func MultiCpuMetrics() []*model.MetricValue {

	mc := []*model.MetricValue{}

	for i := 0; i < len(procStatHistory[0].Cpus); i++ {
		currentCpu := procStatHistory[0].Cpus[i]
		oldCpu := procStatHistory[1].Cpus[i]

		// modify by wjl:2019-03-28: aiops
		tags := fmt.Sprintf("cpu=cpu-%d", i+1)
		component := fmt.Sprintf("Cpu-%d", i+1)
		metricGroup := "MultiCpuMonitor"
		unitPer := "%"
		idle := GaugeValueAIOPS("cpu.idle", CpuIdleSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		user := GaugeValueAIOPS("cpu.user", CpuUserSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		nice := GaugeValueAIOPS("cpu.nice", CpuNiceSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		system := GaugeValueAIOPS("cpu.system", CpuSystemSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		iowait := GaugeValueAIOPS("cpu.wait", CpuIowaitSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		irq := GaugeValueAIOPS("cpu.hi", CpuIrqSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		softirq := GaugeValueAIOPS("cpu.cpu.si", CpuSoftIrqSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)
		steal := GaugeValueAIOPS("cpu.steal", CpuStealSinglecpu(currentCpu, oldCpu), component, metricGroup, unitPer, tags)

		mc = append(mc, idle, user, nice, system, iowait, irq, softirq, steal)

	}

	return mc
}
