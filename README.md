# IpInfo


```
> echo 8.8.8.8 | ipinfo
 _ _____  _  __  _  ____  ____
| || ()_)| ||  \| || ===|/ () \
|_||_|   |_||_|\__||__|  \____/
Press CTRL+C to exit

Enter IP address: IP: 8.8.8.8
Country: United States
ASN: AS15169
ASN Org: GOOGLE
Time Zone: America/Los_Angeles
```

for verbose mode user `-v` or `--verbose` flag
```
> echo 8.8.8.8 | ipinfo -v
 _ _____  _  __  _  ____  ____
| || ()_)| ||  \| || ===|/ () \
|_||_|   |_||_|\__||__|  \____/
Press CTRL+C to exit

Enter IP address: IP: 8.8.8.8
IP Decimal: 134744072
Country: United States
Country ISO: US
Country EU: false
Region Name: California
Region Code: CA
City: Los Angeles
Latitude: 34.054400
Longitude: -118.244100
Time Zone: America/Los_Angeles
ASN: AS15169
ASN Org: GOOGLE
User Agent:
  Product: Mozilla
  Version: 5.0
  Raw Value: Mozilla/5.0 (Android 4.4; Mobile; rv:41.0) Gecko/41.0 Firefox/41.0

```
you can also give ip list as input
```
> cat ips.txt | ipinfo
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
