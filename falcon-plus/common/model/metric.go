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

// modify by wjl 2019-03-20: aiops
// add System, Component, Metricgroup, Datatype, Unit, keep Tags
package model

import (
	"fmt"

	MUtils "github.com/open-falcon/falcon-plus/common/utils"
)

type MetricValue struct {
	System      string      `json:"system"`
	Endpoint    string      `json:"endpoint"`
	Component   string      `json:"component"`
	Metricgroup string      `json:"metricgroup"`
	Metric      string      `json:"metric"`
	Datatype    string      `json:"datatype"`
	Value       interface{} `json:"value"`
	Step        int64       `json:"step"`
	Unit        string      `json:"unit"`
	Type        string      `json:"countertype"`
	Tags        string      `json:"tags"`
	Timestamp   int64       `json:"timestamp"`
}

func (this *MetricValue) String() string {
	return fmt.Sprintf(
		"<System:%s, Endpoint:%s, Component:%s, Metricgroup:%s, Metric:%s, Datatype:%s, Type:%s, Tags:%s, Step:%d, Unit:%s, Time:%d, Value:%v>",
		this.System,
		this.Endpoint,
		this.Component,
		this.Metricgroup,
		this.Metric,
		this.Datatype,
		this.Type,
		this.Tags,
		this.Step,
		this.Unit,
		this.Timestamp,
		this.Value,
	)
}

// Same As `MetricValue`
type JsonMetaData struct {
	System      string      `json:"system"`
	Component   string      `json:"component"`
	Metricgroup string      `json:"metricgroup"`
	Metric      string      `json:"metric"`
	Datatype    string      `json:"datatype"`
	Endpoint    string      `json:"endpoint"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
	Value       interface{} `json:"value"`
	Unit        string      `json:"unit"`
	CounterType string      `json:"countertype"`
	Tags        string      `json:"tags"`
}

func (t *JsonMetaData) String() string {
	return fmt.Sprintf("<JsonMetaData System:%s, Component:%s, Metricgroup:%s, Endpoint:%s, Metric:%s, Datatype:%s, Tags:%s, DsType:%s, Step:%d, Unit:%s, Value:%v, Timestamp:%d>",
		t.System, t.Component, t.Metricgroup, t.Endpoint, t.Metric, t.Datatype, t.Tags, t.CounterType, t.Step, t.Unit, t.Value, t.Timestamp)
}

type MetaData struct {
	System      string            `json:"system"`
	Component   string            `json:"component"`
	Metricgroup string            `json:"metricgroup"`
	Metric      string            `json:"metric"`
	Datatype    string            `json:"datatype"`
	Endpoint    string            `json:"endpoint"`
	Timestamp   int64             `json:"timestamp"`
	Step        int64             `json:"step"`
	Value       float64           `json:"value"`
	Unit        string            `json:"unit"`
	CounterType string            `json:"countertype"`
	Tags        map[string]string `json:"tags"`
}

func (t *MetaData) String() string {
	return fmt.Sprintf("<MetaData System:%s, Component:%s, Metricgroup:%s, Datatype:%s, Endpoint:%s, Metric:%s, Timestamp:%d, Step:%d, Value:%f, Unit:%s, CounterType:%s, Tags:%v>",
		t.System, t.Component, t.Metricgroup, t.Datatype, t.Endpoint, t.Metric, t.Timestamp, t.Step, t.Value, t.Unit, t.CounterType, t.Tags)
}

func (t *MetaData) PK() string {
	return MUtils.PK(t.Endpoint, t.Metric, t.Tags)
}
