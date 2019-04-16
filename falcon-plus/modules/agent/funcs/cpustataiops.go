package funcs

import (
	"github.com/open-falcon/falcon-plus/common/model"
	"runtime"
)

func CpuMetricsAIOPS() []*model.MetricValue {
	if !CpuPrepared() {
		return []*model.MetricValue{}
	}

	cpuIdleVal := CpuIdle()

	// modify by wjl:2019-03-20: aiops
	component := "Cpu"
	metricGroup := "CpuMonitor"
	unit := "_"
	unitPer := "%"
	idle := GaugeValueAIOPS("cpu.idle", cpuIdleVal, component, metricGroup, unitPer)
	user := GaugeValueAIOPS("cpu.user", CpuUser(), component, metricGroup, unitPer)
	nice := GaugeValueAIOPS("cpu.nice", CpuNice(), component, metricGroup, unitPer)
	system := GaugeValueAIOPS("cpu.system", CpuSystem(), component, metricGroup, unitPer)
	iowait := GaugeValueAIOPS("cpu.wait", CpuIowait(), component, metricGroup, unitPer)
	irq := GaugeValueAIOPS("cpu.hi", CpuIrq(), component, metricGroup, unitPer)
	softirq := GaugeValueAIOPS("cpu.cpu.si", CpuSoftIrq(), component, metricGroup, unitPer)
	steal := GaugeValueAIOPS("cpu.steal", CpuSteal(), component, metricGroup, unitPer)
	count := GaugeValueAIOPS("cpu.count", runtime.NumCPU(), component, metricGroup, unit)

	totalcpus := MultiCpuMetrics()

	return append(totalcpus, idle, user, nice, system, iowait, irq, softirq, steal, count)
}

