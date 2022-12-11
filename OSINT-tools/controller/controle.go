package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"

	"cybercops.in/utils"
)

// API root msg
func Root(w http.ResponseWriter, r *http.Request) {
	utils.Result(w, utils.Message{Message: "welcome to the cybercops.in API service", Status: true, Data: nil})
}

// API IPTracer msg
func Ipinfo(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query()["ip"]
	fmt.Println(ip[0])
	//if not installed
	//sudo apt install traceroute

	//init go channels for simultaneous task
	var chans [1]chan bool
	for i := range chans {
		chans[i] = make(chan bool)
	}
	traceout := ""
	shodan := ""
	//channel one
	// go func() { //channel 0 is traceroute going to happens
	// 	cmd := exec.Command("traceroute", ip[0])
	// 	fmt.Println("traceroute is happening")
	// 	out, err := cmd.StdoutPipe()
	// 	if err != nil {
	// 		chans[0] <- false
	// 		return
	// 	}
	// 	cmd.Start()
	// 	scanner := bufio.NewScanner(out)
	// 	for scanner.Scan() {
	// 		m := scanner.Text()
	// 		fmt.Println(m)
	// 		traceout += m + "\n"
	// 	}
	// 	fmt.Println("[OUTPUT]:", traceout)
	// 	chans[0] <- true
	// }()
	//channel two
	go func() { //channel 0 is traceroute going to happens
		cmd := exec.Command("shodan", "host", ip[0])
		fmt.Println("shodan is happening")
		out, err := cmd.StdoutPipe()
		if err != nil {
			chans[0] <- false
			return
		}
		cmd.Start()
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
			shodan += m + "\n"
		}
		fmt.Println("[OUTPUT]:", shodan)
		chans[0] <- true
	}()
	<-chans[0]
	
	TR := Ipresult{}
	TR.Traceroute = traceout
	TR.Shodan = shodan
	utils.Result(w, utils.Message{Message: "get the ip info", Status: true, Data: TR})
}

type Ipresult struct {
	Traceroute string `json:"traceroute"`
	Shodan     string `json:"shodan"`
}

// API Domine name tracer msg
func DomainCheck(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query()["domain"]
	fmt.Println(domain[0])
	//if not installed
	//sudo apt install traceroute

	//init go channels for simultaneous task
	var chans [1]chan bool
	for i := range chans {
		chans[i] = make(chan bool)
	}
	traceout := ""

	//channel one
	go func() { //channel 0 is traceroute going to happens
		cmd := exec.Command("python3", "main.py", domain[0])
		fmt.Println("traceroute is happening")
		out, err := cmd.StdoutPipe()

		if err != nil {
			chans[0] <- false
			return
		}
		cmd.Start()
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
			traceout += m + "\n"
		}
		fmt.Println("[OUTPUT]:", traceout)
		chans[0] <- true
	}()
	<-chans[0]

	TR := Ipresult{}
	TR.Traceroute = traceout
	utils.Result(w, utils.Message{Message: "get the domain info", Status: true, Data: TR})
}

// API Domine name tracer msg
func SearchKey(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["q"]
	fmt.Println(query[0])
	//if not installed
	//sudo apt install traceroute

	//init go channels for simultaneous task
	var chans [1]chan bool
	for i := range chans {
		chans[i] = make(chan bool)
	}
	traceout := ""

	//channel one
	go func() { //channel 0 is traceroute going to happens
		cmd := exec.Command("python3", "tor.py", query[0])
		fmt.Println("tor to find the urls!! is happening")
		out, err := cmd.StdoutPipe()

		if err != nil {
			chans[0] <- false
			return
		}
		cmd.Start()
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
			traceout += m + "\n"
		}
		fmt.Println("[OUTPUT]:", traceout)
		chans[0] <- true
	}()
	<-chans[0]

	utils.Result(w, utils.Message{Message: "get the domain info", Status: true, Data: traceout})
}
