[![Go Test](https://github.com/cmwylie19/sbctl/actions/workflows/unit-tests.yaml/badge.svg)](https://github.com/cmwylie19/sbctl/actions/workflows/unit-tests.yaml)

**Video Demo**   
  
[![Watch the video](https://i9.ytimg.com/vi_webp/0yL-a401YyM/mq2.webp?sqp=CLCZr6IG-oaymwEmCMACELQB8quKqQMa8AEB-AH-CYAC0AWKAgwIABABGGUgZShlMA8=&rs=AOn4CLDSRRUciVjIIZmFWBHMlJSIu0LNPg)](https://www.youtube.com/watch?v=0yL-a401YyM)


# Starburst Config Tool (WIP)

_Tool to simplify Starburst configuration for data practioners._

**TOC**
- [Prereqs](#prereqs)
- [Installation](#installation)
- [Usage](#usage)
- [Build](#build-image-with-sbctl)
- [Releases](#releases)

## Prereqs

The CLI requires a Java virtual machine available on the path. It can be used with Java version 8 and higher. 

Find if Java is installed in your path:

```bash
➜ java -version 
# output
openjdk version "1.8.0_292"
OpenJDK Runtime Environment (AdoptOpenJDK)(build 1.8.0_292-b10)
OpenJDK 64-Bit Server VM (AdoptOpenJDK)(build 25.292-b10, mixed mode)
```

## Installation

To install, go to the [releases page](https://github.com/cmwylie19/sbctl/releases) of the repository and download the binary for your operating system. It is recommended to add the binary to your path. 

**Special Linux instructions**

You may need to change permissions so that the current user is allowed to create an executable `trino` in the `/usr/local/bin` directory, if you do not want to do that, then you can run the commands as root.

```bash
sudo chown $(whoami) /usr/local/bin
```


## Usage 

This tool can install/uninstall the Trino CLI. More features to come.   

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


Install the CLI on a Mac

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


Uninstall the CLI on a Mac

```bash
./sbctl cli --mode=uninstall --os=mac
```

output

```bash
2023/04/03 16:04:30 Binary file deleted successfully!
```


## Build Image with sbctl

```bash
GOOS=<linux|darwin|windows> go build -o sbctl .
mv sbctl build/sbctl
```

## Releases

Each release should include binary files for the target operating systems with their respective sha256 checksums. The below snippet will create the the releases and checksums


```bash
# Intel Mac 
GOARCH=amd64 GOOS=darwin go build -o sbctl-darwin-amd64 . 
sha256sum sbctl-darwin-amd64 > sbctl-darwin-amd64.sha256

# Silicon Mac
GOARCH=arm64 GOOS=darwin go build -o sbctl-darwin-arm64 . 
sha256sum sbctl-darwin-arm64 > sbctl-darwin-arm64.sha256

# Linux 
GOARCH=amd64 GOOS=linux go build -o sbctl-linux-amd64 . 
sha256sum sbctl-linux-amd64 > sbctl-linux-amd64.sha256

# Windows 
GOARCH=amd64 GOOS=windows go build -o sbctl-windows-amd64 . 
sha256sum sbctl-windows-amd64 > sbctl-windows-amd64.sha256
```

Checksums can be checked by:

```bash
sha256sum -c <file.sha256>
```
