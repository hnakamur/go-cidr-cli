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

type previousCmd struct{}

func (*previousCmd) Name() string     { return "previous" }
func (*previousCmd) Synopsis() string { return "Print the previous network of CIDR." }
func (*previousCmd) Usage() string {
	return `previous <CIDR> <PrefixLen>:
  Print the previous network of CIDR with prefix length.
`
}

func (*previousCmd) SetFlags(f *flag.FlagSet) {}

func (c *previousCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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

	nextNet, bad := cidr.PreviousSubnet(network, prefixLen)
	if bad {
		fmt.Fprintf(os.Stderr, "not enough range for previous network.\n")
		return subcommands.ExitUsageError
	}

	fmt.Println(nextNet)
	return subcommands.ExitSuccess
}
