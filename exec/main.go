package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type os_exec struct {
	Cmd   string
	Args  []string
	Query string
}

func NewQuery(q string) *os_exec {
	return &os_exec{
		Cmd:   "osqueryi",
		Args:  []string{"--json", q},
		Query: q,
	}
}

func (*os_exec) Exec(query string) error {
	bin, err := exec.LookPath("osqueryi")
	if err != nil {
		return err
	}

	fmt.Printf("Osquery Path: %s\n", bin)

	args := []string{"osqueryi", "--json", query}
	env := os.Environ()

	err = syscall.Exec(bin, args, env)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	Username string `json:"username"`
}

type SystemInfo struct {
	ComputerName string `json:"computer_name"`
	CPUBrand     string `json:"cpu_brand"`
	CPUType      string `json:"cpu_type"`
	Model        string `json:"hardware_model"`
	Hostname     string `json:"hostname"`
	MemoryTotal  uint64 `json:"physical_memory,string"`
	UUID         string `json:"uuid"`
}

func (r *os_exec) Command(v interface{}) error {
	_, err := exec.LookPath("osqueryi")
	if err != nil {
		return err
	}

	cmd := exec.Command(r.Cmd, r.Args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return err
	}

	data := out.Bytes()
	fmt.Printf("%s", data)
	json.Unmarshal(data, &v)

	return nil
}

func main() {
	info := []SystemInfo{}
	q := NewQuery("select * from system_info;")

	if err := q.Command(&info); err != nil {
		panic(err)
	}

	if len(info) > 0 {
		fmt.Println(info[0].UUID)
		fmt.Println(info[0].MemoryTotal)
		fmt.Println(info[0].Model)
		fmt.Println(info[0].Hostname)
		fmt.Println(info[0].ComputerName)
		fmt.Println(info[0].CPUBrand)
	}
}
