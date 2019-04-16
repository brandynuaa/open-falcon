package funcs

import (
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/nux"
	"log"
)

func NetMetricsAIOPS() []*model.MetricValue {
	return CoreNetMetrics(g.Config().Collector.IfacePrefix)
}

func CoreNetMetricsAIOPS(ifacePrefix []string) []*model.MetricValue {

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

		// add by wjl 2019-03-20: aiops
		component := netIf.Iface
		metricgroup := "NetMonitor"
		unit := "bps"
		// unitPer := "%"
		ret[idx*2+0] = CounterValueAIOPS("net.if.in.bytes", netIf.InBytes, component, metricgroup, unit, iface)
		ret[idx*2+1] = CounterValueAIOPS("net.if.out.bytes", netIf.OutBytes, component, metricgroup, unit, iface)
	}
	return ret
}

