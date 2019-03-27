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
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/nux"
	"log"
)

func NetMetrics() []*model.MetricValue {
	return CoreNetMetrics(g.Config().Collector.IfacePrefix)
}

func CoreNetMetrics(ifacePrefix []string) []*model.MetricValue {

	netIfs, err := nux.NetIfs(ifacePrefix)
	if err != nil {
		log.Println(err)
		return []*model.MetricValue{}
	}

	cnt := len(netIfs)
	// modify by wjl 2019-03-20: aiops
	ret := make([]*model.MetricValue, cnt*2)

	for idx, netIf := range netIfs {
		iface := "iface=" + netIf.Iface
		// ret[idx*26+0] = CounterValue("net.if.in.bytes", netIf.InBytes, iface)
		// ret[idx*26+1] = CounterValue("net.if.in.packets", netIf.InPackages, iface)
		// ret[idx*26+2] = CounterValue("net.if.in.errors", netIf.InErrors, iface)
		// ret[idx*26+3] = CounterValue("net.if.in.dropped", netIf.InDropped, iface)
		// ret[idx*26+4] = CounterValue("net.if.in.fifo.errs", netIf.InFifoErrs, iface)
		// ret[idx*26+5] = CounterValue("net.if.in.frame.errs", netIf.InFrameErrs, iface)
		// ret[idx*26+6] = CounterValue("net.if.in.compressed", netIf.InCompressed, iface)
		// ret[idx*26+7] = CounterValue("net.if.in.multicast", netIf.InMulticast, iface)
		// ret[idx*26+8] = CounterValue("net.if.out.bytes", netIf.OutBytes, iface)
		// ret[idx*26+9] = CounterValue("net.if.out.packets", netIf.OutPackages, iface)
		// ret[idx*26+10] = CounterValue("net.if.out.errors", netIf.OutErrors, iface)
		// ret[idx*26+11] = CounterValue("net.if.out.dropped", netIf.OutDropped, iface)
		// ret[idx*26+12] = CounterValue("net.if.out.fifo.errs", netIf.OutFifoErrs, iface)
		// ret[idx*26+13] = CounterValue("net.if.out.collisions", netIf.OutCollisions, iface)
		// ret[idx*26+14] = CounterValue("net.if.out.carrier.errs", netIf.OutCarrierErrs, iface)
		// ret[idx*26+15] = CounterValue("net.if.out.compressed", netIf.OutCompressed, iface)
		// ret[idx*26+16] = CounterValue("net.if.total.bytes", netIf.TotalBytes, iface)
		// ret[idx*26+17] = CounterValue("net.if.total.packets", netIf.TotalPackages, iface)
		// ret[idx*26+18] = CounterValue("net.if.total.errors", netIf.TotalErrors, iface)
		// ret[idx*26+19] = CounterValue("net.if.total.dropped", netIf.TotalDropped, iface)
		// ret[idx*26+20] = GaugeValue("net.if.speed.bits", netIf.SpeedBits, iface)
		// ret[idx*26+21] = CounterValue("net.if.in.percent", netIf.InPercent, iface)
		// ret[idx*26+22] = CounterValue("net.if.out.percent", netIf.OutPercent, iface)
		// ret[idx*26+23] = CounterValue("net.if.in.bits", netIf.InBytes*8, iface)
		// ret[idx*26+24] = CounterValue("net.if.out.bits", netIf.OutBytes*8, iface)
		// ret[idx*26+25] = CounterValue("net.if.total.bits", netIf.TotalBytes*8, iface)

		// add by wjl 2019-03-20: aiops
		component := netIf.Iface
		metricgroup := "NetMonitor"
		unit := "bps"
		// unitPer := "%"
		ret[idx*2+0] = CounterValueAIOPS("net.if.in.bytes", netIf.InBytes, component, metricgroup, unit, iface)
		// ret[idx*26+1] = CounterValueAIOPS("net.if.in.packets", netIf.InPackages, component, metricgroup, unit, iface)
		// ret[idx*26+2] = CounterValueAIOPS("net.if.in.errors", netIf.InErrors, component, metricgroup, unit, iface)
		// ret[idx*26+3] = CounterValueAIOPS("net.if.in.dropped", netIf.InDropped, component, metricgroup, unit, iface)
		// ret[idx*26+4] = CounterValueAIOPS("net.if.in.fifo.errs", netIf.InFifoErrs, component, metricgroup, unit, iface)
		// ret[idx*26+5] = CounterValueAIOPS("net.if.in.frame.errs", netIf.InFrameErrs, component, metricgroup, unit, iface)
		// ret[idx*26+6] = CounterValueAIOPS("net.if.in.compressed", netIf.InCompressed, component, metricgroup, unit, iface)
		// ret[idx*26+7] = CounterValueAIOPS("net.if.in.multicast", netIf.InMulticast, component, metricgroup, unit, iface)
		ret[idx*2+1] = CounterValueAIOPS("net.if.out.bytes", netIf.OutBytes, component, metricgroup, unit, iface)
		// ret[idx*26+9] = CounterValueAIOPS("net.if.out.packets", netIf.OutPackages, component, metricgroup, unit, iface)
		// ret[idx*26+10] = CounterValueAIOPS("net.if.out.errors", netIf.OutErrors, component, metricgroup, unit, iface)
		// ret[idx*26+11] = CounterValueAIOPS("net.if.out.dropped", netIf.OutDropped, component, metricgroup, unit, iface)
		// ret[idx*26+12] = CounterValueAIOPS("net.if.out.fifo.errs", netIf.OutFifoErrs, component, metricgroup, unit, iface)
		// ret[idx*26+13] = CounterValueAIOPS("net.if.out.collisions", netIf.OutCollisions, component, metricgroup, unit, iface)
		// ret[idx*26+14] = CounterValueAIOPS("net.if.out.carrier.errs", netIf.OutCarrierErrs, component, metricgroup, unit, iface)
		// ret[idx*26+15] = CounterValueAIOPS("net.if.out.compressed", netIf.OutCompressed, component, metricgroup, unit, iface)
		// ret[idx*26+16] = CounterValueAIOPS("net.if.total.bytes", netIf.TotalBytes, component, metricgroup, unit, iface)
		// ret[idx*26+17] = CounterValueAIOPS("net.if.total.packets", netIf.TotalPackages, component, metricgroup, unit, iface)
		// ret[idx*26+18] = CounterValueAIOPS("net.if.total.errors", netIf.TotalErrors, component, metricgroup, unit, iface)
		// ret[idx*26+19] = CounterValueAIOPS("net.if.total.dropped", netIf.TotalDropped, component, metricgroup, unit, iface)
		// ret[idx*26+20] = GaugeValueAIOPS("net.if.speed.bits", netIf.SpeedBits, component, metricgroup, unit, iface)
		// ret[idx*26+21] = CounterValueAIOPS("net.if.in.percent", netIf.InPercent, component, metricgroup, unitPer, iface)
		// ret[idx*26+22] = CounterValueAIOPS("net.if.out.percent", netIf.OutPercent, component, metricgroup, unitPer, iface)
		// ret[idx*26+23] = CounterValueAIOPS("net.if.in.bits", netIf.InBytes*8, component, metricgroup, unit, iface)
		// ret[idx*26+24] = CounterValueAIOPS("net.if.out.bits", netIf.OutBytes*8, component, metricgroup, unit, iface)
		// ret[idx*26+25] = CounterValueAIOPS("net.if.total.bits", netIf.TotalBytes*8, component, metricgroup, unit, iface)
	}
	return ret
}
