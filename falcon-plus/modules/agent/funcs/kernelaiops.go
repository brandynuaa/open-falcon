package funcs

import (
	"fmt"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/toolkits/nux"
	"github.com/toolkits/sys"
	"log"
)

// add by wjl 2019-03-27:aiops Kernel Version
var kernelVersion string

func InitKernelVersion() {
	data, err := sys.CmdOutNoLn("uname", "-r")
	if err != nil {
		log.Println(err)
		kernelVersion = "uname -r error"
		return
	}
	kernelVersion = data
}

func KernelMetricsAIOPS() (L []*model.MetricValue) {

	maxFiles, err := nux.KernelMaxFiles()
	if err != nil {
		log.Println(err)
		return
	}

	// modify by wjl 2019-03-26: aiops
	component := "SystemInfo"
	metricGroup := "SysMonitor"
	unit := "_"
	L = append(L, GaugeValueAIOPS("system.openfile.max", maxFiles, component, metricGroup, unit))

	allocateFiles, err := nux.KernelAllocateFiles()
	if err != nil {
		log.Println(err)
		return
	}

	L = append(L, GaugeValueAIOPS("system.openfile.use", allocateFiles, component, metricGroup, unit))

	if kernelVersion != "" {
		L = append(L, GaugeValueAIOPSString("system.kernel", kernelVersion, component, metricGroup, unit))
	}

	// add by wjl 2019-03-27:aiops Kernel Version
	days, hours, mins, err := nux.SystemUptime()
	if err != nil {
		log.Println(err)
		return
	}
	L = append(L, GaugeValueAIOPSString("system.lastboot", fmt.Sprintf("%d days %d hours %d minutes", days, hours, mins), component, metricGroup, unit))

	load, err := nux.LoadAvg()
	if err != nil {
		log.Println(err)
		return nil
	}
	unit1 := "%"
	load1 := GaugeValueAIOPS("system.load.1min",load.Avg1min, component, metricGroup,unit1)
	load5 := GaugeValueAIOPS("system.load.5min",load.Avg5min,component,metricGroup,unit1)
	load15 := GaugeValueAIOPS("system.load.15min",load.Avg15min,component,metricGroup,unit1)

	L = append(L,load1,load5,load15)

	return
}

