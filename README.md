# JSON Pretty Print - Golang

A quick and simple Golang tool to help you prettify your JSON.

```bash
go get cpl.li/go/jpp
```

Usage is straightforward:

```bash
# Get the help message with all flag description
jpp -help

# Change indent style and/or prefix
jpp -indent "    " -prefix " "

# Use a file as input
jpp -file /var/log/docker/output.json

# Use it as a pipe
cmd0 | cmd1 | jpp
```

Feel free to make whatever changes you want to the code.
