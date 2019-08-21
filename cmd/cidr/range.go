package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/google/subcommands"
)

type rangeCmd struct{}

func (*rangeCmd) Name() string     { return "range" }
func (*rangeCmd) Synopsis() string { return "Print address range of CIDR." }
func (*rangeCmd) Usage() string {
	return `range <CIDR>:
  Print address range of CIDR arg to stdout.
`
}

func (*rangeCmd) SetFlags(f *flag.FlagSet) {}

func (c *rangeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s %s", os.Args[0], c.Usage())
		return subcommands.ExitUsageError
	}
	_, network, err := net.ParseCIDR(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse CIDR\n")
		return subcommands.ExitUsageError
	}

	first, last := cidr.AddressRange(network)
	fmt.Printf("first: %s\nlast:  %s\n", first, last)
	return subcommands.ExitSuccess
}
