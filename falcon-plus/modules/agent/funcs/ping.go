package funcs

import (
	"bufio"
	"bytes"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/file"
	"github.com/toolkits/sys"
	"io"
	"strconv"
	"strings"
)

type PingInfos struct {
	Packets float64
	ReceivedP float64
	LossP float64
	Responsetime float64

	Min float64
	Max float64
	Avg float64
}

/*
ping -c 2 -q 140.143.79.72 | grep -E "packets|rtt"
2 packets transmitted, 2 received, 0% packet loss, time 999ms
rtt min/avg/max/mdev = 0.510/0.538/0.566/0.028 m
 */
func PingInfo(serviceip string)  (*PingInfos,error){

	cmdsting := "ping -c 4 -q "+serviceip+" | grep -E 'packets|rtt'"
	bs, err := sys.CmdOutBytes("sh", "-c", cmdsting)
	if err != nil {
		return nil,err
	}
	pi := new(PingInfos)
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	for{
		line, err := file.ReadLine(reader)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil,err
		}
		lineString := string(line)
		if strings.Contains(lineString,"packets transmitted"){
			packets_arr := strings.Split(lineString,",")
			pi.Packets,_ = strconv.ParseFloat(strings.Trim(packets_arr[0]," packets transmitted"),64)
			pi.ReceivedP,_=strconv.ParseFloat(strings.Trim(packets_arr[1]," received"),64)
			pi.LossP = pi.Packets - pi.ReceivedP

		}
		if strings.Contains(lineString,"rtt") {
			rtt_arr := strings.Split(lineString,"=")
			rttsum_arr := strings.Split(rtt_arr[1],"/")
			pi.Min,_ = strconv.ParseFloat(rttsum_arr[0],64)
			pi.Avg,_ =strconv.ParseFloat(rttsum_arr[1],64)
			pi.Max,_ = strconv.ParseFloat(rttsum_arr[2],64)
			pi.Responsetime = pi.Avg

		}
	}

	return pi,nil
}

func PingMetrics() []*model.MetricValue  {

	pi,err := PingInfo(g.Config().ServiceIp)
	if err != nil || pi == nil{
		return nil
	}
	var jitterF float64
	var statusI int
	if pi.ReceivedP == 0 {
		statusI = 0
		jitterF = 0
	}else {
		statusI = 1
		jitterF = (pi.Max - pi.Min)/3
	}

	component := "ping"
	metricGroup := "PingMonitor"
	unit := "_"
	unit2 := "ms"

	status := GaugeValueAIOPS("net.ping.status",statusI, component, metricGroup, unit)
	responsetime := GaugeValueAIOPS("net.ping.responsetime",pi.Responsetime, component, metricGroup, unit2)
	lost := GaugeValueAIOPS("net.ping.lost",pi.Packets - pi.ReceivedP, component, metricGroup, unit)
	jitter := GaugeValueAIOPS("net.ping.jitter",jitterF, component, metricGroup, unit2)

	return []*model.MetricValue{status,responsetime,lost,jitter}
}

