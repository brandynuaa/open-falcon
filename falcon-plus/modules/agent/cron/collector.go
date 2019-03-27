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

package cron

import (
	"encoding/json"
	"fmt"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/funcs"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"time"
)

func InitDataHistory() {
	for {
		funcs.UpdateCpuStat()
		funcs.UpdateDiskStats()
		time.Sleep(g.COLLECT_INTERVAL)
	}
}

func Collect() {

	if !g.Config().Transfer.Enabled {
		return
	}

	if len(g.Config().Transfer.Addrs) == 0 {
		return
	}

	for _, v := range funcs.Mappers {
		go collect(int64(v.Interval), v.Fs)
	}
}

func collect(sec int64, fns []func() []*model.MetricValue) {
	t := time.NewTicker(time.Second * time.Duration(sec))
	defer t.Stop()
	for {
		<-t.C

		// add by wjl 2019-03-20: aipos
		// add system value from config
		system := g.System()

		hostname, err := g.Hostname()
		if err != nil {
			continue
		}

		mvs := []*model.MetricValue{}
		// modify by wjl 2019-03-22: aiops
		// ignoreMetrics := g.Config().IgnoreMetrics

		for _, fn := range fns {
			items := fn()
			if items == nil {
				continue
			}

			if len(items) == 0 {
				continue
			}

			for _, mv := range items {
				// if b, ok := ignoreMetrics[mv.Metric]; ok && b {
				// 	continue
				// } else {
				mvs = append(mvs, mv)
				// }
			}
		}

		now := time.Now().Unix()
		for j := 0; j < len(mvs); j++ {
			mvs[j].Step = sec
			mvs[j].System = system // add by wjl 2019-03-20: aipos
			mvs[j].Endpoint = hostname
			mvs[j].Timestamp = now
			mvs[j].Datatype = "Float" // add by wjl 2019-03-20: aipos
		}

		// g.SendToTransfer(mvs)

		for _, b := range mvs {
			// postform := fmt.Sprintf("[{\"endpoint\":\"%s\", \"metric\":\"%s\", \"counterType\":\"%s\", \"tags\":\"%s\", \"step\":%d, \"timestamp\":%d, \"value\":%v}]",
			// 	b.Endpoint,
			// 	b.Metric,
			// 	b.Type,
			// 	b.Tags,
			// 	b.Step,
			// 	b.Timestamp,
			// 	b.Value)
			// fmt.Println(postform)
			postform := toJson(b)
			fmt.Println(postform)
			g.SyncProducer(postform, g.Config().KafkaCfg)
		}

	}
}

func toJson(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
