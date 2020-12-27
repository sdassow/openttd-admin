# OpenTTD Admin

Golang based library and command line tool for OpenTTD server administration.

Currently a simple chat to communicate with players on the server is implemented:
```
$ ./openttd-admin chat -h
Chat with other players on a server

Usage:
  ./openttd-admin chat [address] [flags]

Flags:
  -c, --config string    openttd config file (default "~/.config/openttd/openttd.cfg")
  -h, --help             help for chat
  -n, --name string      client name (default "OpenTTD-Admin")
  -v, --version string   client version (default "1.10.3")
```

Sample output:
```
$ ./openttd-admin chat localhost:3977
Server admin password:
2020/12/27 16:36:30 connecting to localhost:3977...
# &{Version:1 Frequencies:[{Bool:true Num:0 Flags:63} {Bool:true Num:1 Flags:65} {Bool:true Num:2 Flags:65} {Bool:true Num:3 Flags:61} {Bool:true Num:4 Flags:61} {Bool:true Num:5 Flags:64} {Bool:true Num:6 Flags:64} {Bool:true Num:7 Flags:1} {Bool:true Num:8 Flags:64} {Bool:true Num:9 Flags:64}] Bool:false}
2020/12/27 16:36:30 connected
# &{ClientId:1 HostName: ClientName: ClientLang:NetlangAny JoinDate:0000-01-01 00:00:00 +0000 UTC ClientPlayers:255}
# &{ClientId:7 HostName:::1 ClientName:Player ClientLang:NetlangAny JoinDate:2198-06-24 00:00:00 +0000 UTC ClientPlayers:0}
# &{Index:0 CompanyName:Player Transport ManagerName:Player Colour:12 IsPasswordProtected:false InauguratedYear:2198 IsAI:false QuartersOfBankruptcy:0 ShareOwners:[255 255 255 255]}
> hello world
1792> hello back
>
```

## Directories

 * cmd/ - admin commands
 * pkg/admin/ - admin protocol implementation

## References

 - Server admin port documentation: https://wiki.openttd.org/en/Development/Server%20Admin%20Port%20Development
 - Game source code: https://github.com/OpenTTD/OpenTTD
 - Java based network admin library: https://dev.openttdcoop.org/projects/joan

## License

MIT
