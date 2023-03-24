# Starburst Config Tool (WIP)

_Tool to simplify Starburst configuration for data practioners._

**TOC**
- [Installation](#installation)
- [Usage](#usage)
- [Build](#build)

## Installation
TODO

## Usage 

This tool can install/uninstall the Trino CLI. More features to come soon.   

FYI - This go application will be compiled into a binary called `subctl`, there will be a version compiled for windows, mac, and linux (RPM Based). Once this occurs, there will be **no** dependecy in the user having go. For now until it is compile, there is a dependency.

Run the tool

```bash
go run main.go
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
  cli         Install the Trino CLI
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for sbctl

Use "sbctl [command] --help" for more information about a command.
```

Look at the Trino CLI options

```bash
go run main.go cli -h
```

output

```bash
Install the Trino CLI which provides a terminal-based, 
	interactive shell for running queries.
	
	The CLI is a self-executing JAR file, which means it acts
	like a normal UNIX executable. The CLI uses the Trino client
	REST API over HTTP/HTTPS to communicate with the coordinator
	on the cluster.

Usage:
  sbctl cli [flags]

Flags:
  -h, --help          help for cli
      --mode string   Options: install or uninstall (default "install")
      --os string     Target Operating System to install the Trino CLI on. Options: mac, linux, or windows. (default "mac")
```


Install the CLI on a Mac (other OS's are a WIP)

```bash
go run main.go cli --mode=install --os=mac
```

output

```bash
openjdk version "1.8.0_292"
OpenJDK Runtime Environment (AdoptOpenJDK)(build 1.8.0_292-b10)
OpenJDK 64-Bit Server VM (AdoptOpenJDK)(build 25.292-b10, mixed mode)

Downloaded and renamed file to: /usr/local/bin/trino
```


## Build

```bash
GOOS=<linux|darwin|windows> go build -o sbctl .
mv sbctl build/sbctl
```