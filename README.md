# IpInfo


## Usage
```
▶ ipinfo -h
 _ _____  _  __  _  ____  ____
| || ()_)| ||  \| || ===|/ () \
|_||_|   |_||_|\__||__|  \____/
Press CTRL+C to exit

Usage: ipinfo [options]
Options:
  -h, --help                            Show this help message and exit
  -v, --verbose                         Show full information
  -w, --workers                         Number of workers(default: 1)
  -i, --interval                        Interval between requests(default: 0.4)
```

## Example

```shell
> echo 8.8.8.8 | ipinfo
```
```
 _ _____  _  __  _  ____  ____
| || ()_)| ||  \| || ===|/ () \
|_||_|   |_||_|\__||__|  \____/
Press CTRL+C to exit

[*] 8.8.8.8
Country: United States
ASN: AS15169
ASN Org: GOOGLE
Time Zone: America/Los_Angeles
```
### help
```shell
 _ _____  _  __  _  ____  ____
| || ()_)| ||  \| || ===|/ () \
|_||_|   |_||_|\__||__|  \____/
Press CTRL+C to exit

Usage: ipinfo [options]
Options:
  -h, --help                            Show this help message and exit
  -v, --verbose                         Show full information
  -w, --workers                         Number of workers(default: 1)
  -i, --interval                        Interval between requests(default: 1.0)
````

```
> ipinfo -v -w 10 -i 0.2
 _ _____  _  __  _  ____  ____
| || ()_)| ||  \| || ===|/ () \
|_||_|   |_||_|\__||__|  \____/
Press CTRL+C to exit
```
for more information use `-h` or `--help` flag
## Installation

You can just [download a binary for Linux, Mac, Windows or FreeBSD and run it](https://github.com/beyrakIn/ipinfo).
Put the binary in your `$PATH` (e.g. in `/usr/local/bin`) to make it easy to use:
```
▶ git clone https://github.com/beyrakIn/ipinfo.git
▶ cd ipinfo 
▶ sudo mv ipinfo /usr/local/bin/
```
