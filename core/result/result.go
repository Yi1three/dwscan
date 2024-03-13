package result

import (
	"fmt"
	"sync"
)

// string是IP，键值是port，port又是个map，每个port对应多个vulnerability
type Result struct {
	sync.RWMutex
	Vuls map[string]map[int][]*Vulnerability
}

type Vulnerability struct {
	VulType     string
	Description string //描述包含造成漏洞的参数之类的，全丢在这打印出去得了
}

func NewResult() *Result {
	Vuls := make(map[string]map[int][]*Vulnerability)
	return &Result{Vuls: Vuls}
}

// 一个漏洞记录由ip, port, Vulnerability struct构成。
func (r *Result) AddVul(ip string, port int, vulnerability *Vulnerability) {
	r.Lock()
	defer r.Unlock()

	if r.Vuls[ip] == nil {
		r.Vuls[ip] = make(map[int][]*Vulnerability)
		r.Vuls[ip][port] = append(r.Vuls[ip][port], vulnerability)
		return
	}

	r.Vuls[ip][port] = append(r.Vuls[ip][port], vulnerability)
}

// test function
func (r *Result) PrintMapValues() {
	for ip, ports := range r.Vuls {
		fmt.Printf("IP: %s\n", ip)
		for port, vulnerabilities := range ports {
			fmt.Printf("  Port: %d\n", port)
			for _, vul := range vulnerabilities {
				fmt.Printf("    Vulnerability Type: %s, Description: %s\n", vul.VulType, vul.Description)
			}
		}
	}
}
