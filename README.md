go-cidr-cli
===========

A command line tool for some computation of a network CIDR.

## How to install

```
go get -u github.com/hnakamur/go-cidr-cli/cmd/cidr
```

## Usage

```
$ cidr
Usage: cidr <flags> <subcommand> <subcommand args>

Subcommands:
        commands         list all command names
        contains         Check a network contains an IP.
        flags            describe all known top-level flags
        help             describe subcommands and their syntax
        next             Print the next network of CIDR.
        previous         Print the previous network of CIDR.
        range            Print address range of CIDR.


Use "cidr flags" for a list of top-level flags

$ cidr help contains
contains [-quiet] <CIDR> <IP>:
  Print whether a network contains an IP.
  -quiet
        print nothing and exit with 0 if contained, 1 otherwise

$ cidr help range
range <CIDR>:
  Print address range of CIDR arg to stdout.

$ cidr help next
next <CIDR> <PrefixLen>:
  Print the next network of CIDR with prefix length.

$ cidr help previous
previous <CIDR> <PrefixLen>:
  Print the previous network of CIDR with prefix length.
```

## License

MIT License
