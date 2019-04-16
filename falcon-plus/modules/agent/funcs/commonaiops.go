package funcs

import (
	"github.com/open-falcon/falcon-plus/common/model"
	"strings"
)


// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit
func NewMetricValueAIOPS(metric string, val interface{}, dataType string, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	mv := model.MetricValue{
		Metric:      metric,
		Value:       val,
		Type:        dataType,
		Component:   com,
		Datatype:    "Float",
		Metricgroup: mgroup,
		Unit:        unit,
	}

	size := len(tags)

	if size > 0 {
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv
}

// add by wjl 2019-03-27: aiops
// add component, metricgroup, unit
func NewMetricValueAIOPSString(metric string, val interface{}, dataType string, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	mv := model.MetricValue{
		Metric:      metric,
		Value:       val,
		Type:        dataType,
		Component:   com,
		Datatype:    "String",
		Metricgroup: mgroup,
		Unit:        unit,
	}

	size := len(tags)

	if size > 0 {
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv
}

// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit, datatype=Float
func GaugeValueAIOPS(metric string, val interface{}, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	return NewMetricValueAIOPS(metric, val, "GAUGE", com, mgroup, unit, tags...)
}

// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit, datatype=String
func GaugeValueAIOPSString(metric string, val interface{}, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	return NewMetricValueAIOPSString(metric, val, "GAUGE", com, mgroup, unit, tags...)
}

// add by wjl 2019-03-20: aiops
// add component, metricgroup, unit
func CounterValueAIOPS(metric string, val interface{}, com string, mgroup string, unit string, tags ...string) *model.MetricValue {
	return NewMetricValueAIOPS(metric, val, "COUNTER", com, mgroup, unit, tags...)
}

