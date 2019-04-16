package funcs

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/file"
	"github.com/toolkits/sys"
	"io"
	"regexp"
	"strconv"
	"strings"
)
type PortInfo struct {
	/*
	ss -op state all | grep 1988
	tcp    ESTAB  0   0   127.0.0.1:58074   127.0.0.1:1988  users:(("falcon-graph",pid=20082,fd=11))
	 */
	Protocol string
	Status   string
	Users    string
}

type ListenInfo struct {
	Listenstatus int
	Pid          float64
	Program		 string
}

func (this *PortInfo) String() string {
	return fmt.Sprintf(
		"<Protocol:%s, Status:%s, Users:%s>",
		this.Protocol,
		this.Status,
		this.Users,
	)
}

func NumPortInfos(reader *bufio.Reader) int{
	count :=0
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return count
}

func PortInfos(port string) ([]*PortInfo,error){
	var err error
	var bs []byte
	cmdinfo := "ss -op state all | grep "+port
	bs, err = sys.CmdOutBytes("sh", "-c", cmdinfo)
	if err != nil {
		return nil,err
	}

	reader := bufio.NewReader(bytes.NewBuffer(bs))
	pontiffs := []*PortInfo{}

	for {

		line, err := file.ReadLine(reader)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil,err
		}
		lineStr := string(line)
		reg := regexp.MustCompile(`\s+`)
		portinfo_arr := reg.Split(strings.TrimSpace(lineStr), -1)

		pi := new(PortInfo)
		pi.Protocol = portinfo_arr[0]
		pi.Status = portinfo_arr[1]
		pi.Users = portinfo_arr[6]
		if strings.Contains(portinfo_arr[5],port) || strings.Contains(portinfo_arr[5],":::*"){
			pontiffs = append(pontiffs,pi)
		}
	}
	return pontiffs,nil
}

func PortInfoMetrics() []*model.MetricValue {

	ps,_ := PortInfos(g.Config().PortCollect)
	pi,err := PortListenInfo(g.Config().PortCollect)
	if err != nil || pi == nil{
		return nil
	}

	estabcount := 0
	closewaitcount := 0
	timewaitcount := 0
	synrecvcount := 0
	for i :=0; i< len(ps);i++{
		pi := ps[i]
		switch pi.Status {
		case "ESTAB":
			estabcount++
		case "CLOSE-WAIT":
			closewaitcount++
		case "TIME-WAIT":
			timewaitcount++
		case "SYN-RECV":
			synrecvcount++
		}
	}

	component := "port"
	metricGroup := "PortMonitor"
	unit := "_"

	listen := GaugeValueAIOPS("net.port.listen",pi.Listenstatus, component, metricGroup, unit)
	pid := GaugeValueAIOPS("net.port.pid",pi.Pid, component, metricGroup, unit)
	program := GaugeValueAIOPSString("net.port.program",pi.Program, component, metricGroup, unit)
	establish := GaugeValueAIOPS("net.port.establish.count",estabcount, component, metricGroup, unit)
	closewait := GaugeValueAIOPS("net.port.closewait.count",closewaitcount, component, metricGroup, unit)
	timewait := GaugeValueAIOPS("net.port.timewait.count",timewaitcount, component, metricGroup, unit)
	synrecv := GaugeValueAIOPS("net.port.synrecv.count",synrecvcount, component, metricGroup, unit)

	return []*model.MetricValue{listen,pid,program,establish,closewait,timewait,synrecv}
}


func PortListenInfo(port string) (*ListenInfo,error){
	li := new(ListenInfo)
	cmdlisten := "ss -op state all | grep "+port+" |grep LISTEN"
	out, err := sys.CmdOutBytes("sh", "-c", cmdlisten)
	if err != nil {
		return nil,err
	}
	reader := bufio.NewReader(bytes.NewBuffer(out))
	retcode, err := file.ReadLine(reader)
	if err != nil {
		return nil, err
	}
	// (("falcon-agent",pid=20191,fd=3))  截取服务名和pid
	msgarr := strings.Split(string(retcode),"users:")
	tmparr := strings.Split(strings.TrimRight(strings.TrimLeft(msgarr[1],"(("),"))"),",")
	li.Listenstatus = 1
	li.Program = strings.Trim(tmparr[0],"\"")
	li.Pid,_ = strconv.ParseFloat(strings.Trim(tmparr[1],"pid="),64)

	return li,nil
}

func main() {
	ps,_ := PortInfos("1988")
	println(len(ps))
	for _,in := range ps{
		println(in.String())
	}

	println("----------------------------------------")
	s := "tcp    LISTEN     0      128                 :::1988                 :::*        users:((\"falcon-agent\",pid=20191,fd=3))"
	msgarr := strings.Split(s,"users:")
	tmparr := strings.Split(strings.TrimRight(strings.TrimLeft(msgarr[1],"(("),"))"),",")
	println(strings.Trim(tmparr[0],"\""))
	println(strings.Trim(tmparr[1],"pid="))


}
