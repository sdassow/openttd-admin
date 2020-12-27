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
  -d, --debug            debug mode
  -h, --help             help for chat
  -n, --name string      client name (default "OpenTTD-Admin")
  -v, --version string   client version (default "1.10.3")
```

Sample output:
```
$ ./openttd-admin chat localhost:3977
Server admin password:
2020/12/27 17:33:28 connecting to localhost:3977...
2020/12/27 17:33:28 connected
> hello world
2020/12/27 17:33:43 1792> yeah, hi...
> ^C
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
