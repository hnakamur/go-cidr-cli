package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/google/subcommands"
)

type containsCmd struct {
	quiet bool
}

func (*containsCmd) Name() string     { return "contains" }
func (*containsCmd) Synopsis() string { return "Check a network contains an IP." }
func (*containsCmd) Usage() string {
	return `contains [-quiet] <CIDR> <IP>:
  Print whether a network contains an IP.
`
}

func (c *containsCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.quiet, "quiet", false, "print nothing and exit with 0 if contained, 1 otherwise")
}

func (c *containsCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s %s", os.Args[0], c.Usage())
		return subcommands.ExitUsageError
	}
	_, network, err := net.ParseCIDR(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse CIDR\n")
		return subcommands.ExitUsageError
	}

	ip := net.ParseIP(args[1])
	if ip == nil {
		fmt.Fprintf(os.Stderr, "failed to parse IP\n")
		return subcommands.ExitUsageError
	}

	if !network.Contains(ip) {
		if !c.quiet {
			fmt.Println("does not contain")
		}
		return subcommands.ExitFailure
	}

	if !c.quiet {
		fmt.Println("contains")
	}
	return subcommands.ExitSuccess
}
