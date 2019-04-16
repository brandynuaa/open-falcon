package funcs

import (
	"bufio"
	"bytes"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/file"
	"github.com/toolkits/sys"
	"io"
	"regexp"
	"strings"
)

/*
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 52:54:00:e4:89:3a brd ff:ff:ff:ff:ff:ff
    inet 172.21.0.14/20 brd 172.21.15.255 scope global eth0
       valid_lft forever preferred_lft forever
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 56:84:7a:fe:97:99 brd ff:ff:ff:ff:ff:ff
    inet 172.17.42.1/16 scope global docker0
       valid_lft forever preferred_lft forever
 */
type NetInfo struct {
	AdminStatus int
	OperStatus int
	Mac string
	address string
}

func NetCardInfo(netname string)  (*NetInfo,error){

	cmdnet :="ip addr show "+netname
	bs, err := sys.CmdOutBytes("sh", "-c", cmdnet)
	if err != nil {
		return nil,err
	}
	ni := new(NetInfo)
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	for {
		line, err := file.ReadLine(reader)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil,err
		}

		lineString := string(line)
		if strings.Contains(lineString,netname+":"){
			net_arr := strings.Split(lineString,"mtu")
			if strings.Contains(net_arr[1],"UP"){
				ni.OperStatus = 0
			}else if strings.Contains(net_arr[1],"DOWN") {
				ni.OperStatus = 1
			}else {
				ni.OperStatus = 2
			}
		}
		if strings.Contains(lineString,"link/ether"){
			reg := regexp.MustCompile(`\s+`)
			link_arr := reg.Split(strings.TrimSpace(lineString), -1)
			ni.Mac = link_arr[1]
		}
		if strings.Contains(lineString,"inet"){
			reg := regexp.MustCompile(`\s+`)
			inet_arr := reg.Split(strings.TrimSpace(lineString), -1)
			ni.address = strings.Split(inet_arr[1],"/")[0]
			//println("==============================")
			//println(ni.address)
		}

	}
	return ni,nil
}

func NetCardMetrics() []*model.MetricValue {

	mv := []*model.MetricValue{}

	for _, ne := range g.Config().NetCards{

		nci,err := NetCardInfo(ne)
		if err != nil || nci == nil{
			return nil
		}

		component := ne
		metricGroup := "NetMonitor"
		unit := "_"

		status := GaugeValueAIOPS("net.if.admin.status",nci.AdminStatus, component, metricGroup,unit)
		operstatus := GaugeValueAIOPS("net.if.operation.status",nci.OperStatus, component, metricGroup,unit)
		mac := GaugeValueAIOPSString("net.if.mac",nci.Mac,component, metricGroup, unit)
		address := GaugeValueAIOPSString("net.if.address",nci.address, component, metricGroup,unit)

		mv = append(mv,status,operstatus,mac,address)
	}

	return mv
}
