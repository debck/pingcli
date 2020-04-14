# pingCLI

> Ping CLI application allows user to send ICMP echo request in a loop to the target hostname or an IP address, while receiving messages and reporting loss and success % and RTT times for each message.

## Usage

- `pingcli --help`

```bash

Usage:
  pingcli [arg] [flags]

Flags:
  -h, --help   help for pingcli

Args:
    IP address / Hostname

```

## Features

- Send ICMP "echo requests" in an infinite loop.
- Reports loss% and RTT times for each message.
- Supports only IPv4 now.

### Examples

> Passing hostname as "google.com"

![2](https://user-images.githubusercontent.com/33368759/79215847-d9440180-7e69-11ea-8721-d89def5df59e.PNG)

> Passing IP address "192.168.0.1"

![3](https://user-images.githubusercontent.com/33368759/79215868-e06b0f80-7e69-11ea-848c-30cd04d65205.PNG)
