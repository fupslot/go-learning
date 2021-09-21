package main

import (
	"bytes"
	"errors"
	"fmt"
	cache "go-learn/cache/pkg"
	"log"
	"os/exec"
	"strings"
	"time"
)

func callPs(arg string) ([]string, error) {
	bin, err := exec.LookPath("ps")
	if err != nil {
		return nil, err
	}

	var args = []string{"-c", "-e", "-o", arg}
	cmd := exec.Command(bin, args...)

	var b bytes.Buffer

	cmd.Stderr = &b
	cmd.Stdout = &b

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	err = WaitTimeout(cmd, 3*time.Second)
	if err != nil {
		return nil, err
	}

	out := strings.Split(b.String(), "\n")
	fmt.Println(out)

	return out, nil
}

func WaitTimeout(c *exec.Cmd, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	done := make(chan error)
	go func() { done <- c.Wait() }()
	select {
	case err := <-done:
		timer.Stop()
		return err
	case <-timer.C:
		if err := c.Process.Kill(); err != nil {
			log.Printf("FATAL error killing process: %s", err)
			return err
		}
		// wait for the command to return after killing it
		<-done
		return errors.New("command timed out")
	}
}

var AppProcessFields = "pid,ppid,utime,stime,etime,state,rss,vsize,pagein,command"

func main() {
	fmt.Println(cache.BuildCacheKeys("example", "pkg", "cache"))

	out, err := callPs(AppProcessFields)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
