// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func KernelMetrics() (L []*model.MetricValue) {

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

	// L = append(L, GaugeValue("kernel.maxfiles", maxFiles))

	// maxProc, err := nux.KernelMaxProc()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// L = append(L, GaugeValue("kernel.maxproc", maxProc))

	allocateFiles, err := nux.KernelAllocateFiles()
	if err != nil {
		log.Println(err)
		return
	}

	L = append(L, GaugeValueAIOPS("system.openfile.use", allocateFiles, component, metricGroup, unit))
	// L = append(L, GaugeValue("kernel.files.allocated", allocateFiles))
	// L = append(L, GaugeValue("kernel.files.left", maxFiles-allocateFiles))

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

	return
}
