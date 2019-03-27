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
	"strings"
)

func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *model.MetricValue {
	mv := model.MetricValue{
		Metric: metric,
		Value:  val,
		Type:   dataType,
	}

	size := len(tags)

	if size > 0 {
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv
}

// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit
func NewMetricValueAIOPS(metric string, val interface{}, dataType string, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	mv := model.MetricValue{
		Metric:      metric,
		Value:       val,
		Type:        dataType,
		Component:   com,
		Metricgroup: mgroup,
		Unit:        unit,
	}

	size := len(tags)

	if size > 0 {
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv
}

func GaugeValue(metric string, val interface{}, tags ...string) *model.MetricValue {
	return NewMetricValue(metric, val, "GAUGE", tags...)
}

func CounterValue(metric string, val interface{}, tags ...string) *model.MetricValue {
	return NewMetricValue(metric, val, "COUNTER", tags...)
}

// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit
func GaugeValueAIOPS(metric string, val interface{}, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	return NewMetricValueAIOPS(metric, val, "GAUGE", com, mgroup, unit, tags...)
}

// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit
func CounterValueAIOPS(metric string, val interface{}, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	return NewMetricValueAIOPS(metric, val, "COUNTER", com, mgroup, unit, tags...)
}
