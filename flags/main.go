package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namsral/flag"
)

type config struct {
	worker int
}

func (c *config) Init() error {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, "", "Path to config file")

	worker := flags.Int("worker", 1, "Number of workers")

	err := flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	c.worker = *worker

	return nil
}

func run(ctx context.Context, c *config, out io.Writer) error {
	c.Init()
	log.SetOutput(out)
	fmt.Printf("PID %d\n", os.Getpid())
	for {
		select {
		case <-ctx.Done():
			return nil
		case t := <-time.Tick(2 * time.Second):
			fmt.Fprintf(out, "%d worker %s\n", c.worker, t)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			fmt.Printf("Got SIGINT/SIGTERM, exiting.")
			cancel()
			os.Exit(1)
		case <-ctx.Done():
			os.Exit(1)
		}
	}()

	c := &config{}

	if err := run(ctx, c, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
