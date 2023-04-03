[![Go Test](https://github.com/cmwylie19/sbctl/actions/workflows/unit-tests.yaml/badge.svg)](https://github.com/cmwylie19/sbctl/actions/workflows/unit-tests.yaml)

# Starburst Config Tool (WIP)

_Tool to simplify Starburst configuration for data practioners._

**TOC**
- [Installation](#installation)
- [Usage](#usage)
- [Build](#build-image-with-sbctl)

## Installation
TODO

## Usage 

This tool can install/uninstall the Trino CLI. More features to come.   

FYI - This go application will be compiled into a binary called `sbctl`, there will be a version compiled for windows, mac, and linux (RPM Based). Once this occurs, there will be **no** dependecy in the user having go. For now until it is compile, there is a dependency.

Run the application

```bash
go run main.go
```

Build and run the binary 

```bash
# from parent directory
GOOS=<linux|darwin|windows> go build -o sbctl .

./sbctl help
```


output

```bash
A tool to simplify and automate installations and configurations of
Starburst clients, libraries, and connections for data customers.

The tool is designed to unify the configuration into a single source of truth
in order to simplify the process of configuring, debugging, and generally managing
the Starburst client, libraries, and connections.

Usage:
  sbctl [command]

Available Commands:
  cli         Install/Uninstall the Trino CLI
  help        Help about any command

Flags:
  -h, --help   help for sbctl

Use "sbctl [command] --help" for more information about a command.
```

Look at the Trino CLI options

```bash
./sbctl cli -h
```

output

```bash
Install/Uninstall the Trino CLI which provides a terminal-based, 
interactive shell for running queries.

The CLI is a self-executing JAR file, which means it acts
like a normal UNIX executable. The CLI uses the Trino client
REST API over HTTP/HTTPS to communicate with the coordinator
on the cluster. 

More info can be found here: https://starburstdata.github.io/latest/client/cli.html

Examples:
  sbctl cli --os=linux --mode=install

  sbctl cli --os=mac --mode=uninstall

  sbctl cli --os=windows --mode=install

Usage:
  sbctl cli [flags]

Flags:
  -h, --help          help for cli
      --mode string   Options: install or uninstall
      --os string     Target Operating System to install the Trino CLI on. Options: mac, linux, or windows.
```


Install the CLI on a Mac (other OS's are a WIP)

```bash
./sbctl cli --mode=install --os=mac
```

output

```bash
2023/04/03 15:49:11 Downloaded and renamed file to: /usr/local/bin/trino

Success!

Usage:
 trino --version
```


## Build Image with sbctl

```bash
GOOS=<linux|darwin|windows> go build -o sbctl .
mv sbctl build/sbctl
```