package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/google/subcommands"
)

type nextCmd struct{}

func (*nextCmd) Name() string     { return "next" }
func (*nextCmd) Synopsis() string { return "Print the next network of CIDR." }
func (*nextCmd) Usage() string {
	return `next <CIDR> <PrefixLen>:
  Print the next network of CIDR with prefix length.
`
}

func (*nextCmd) SetFlags(f *flag.FlagSet) {}

func (c *nextCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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
	prefixLen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse PrefixLen\n")
		return subcommands.ExitUsageError
	}

	nextNet, bad := cidr.NextSubnet(network, prefixLen)
	if bad {
		fmt.Fprintf(os.Stderr, "not enough range for next network.\n")
		return subcommands.ExitUsageError
	}

	fmt.Println(nextNet)
	return subcommands.ExitSuccess
}
