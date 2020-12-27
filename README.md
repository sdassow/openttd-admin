# OpenTTD Admin

Golang based library and command line tool for OpenTTD server administration.

Currently a simple remote console interface and a chat to communicate with players on the server is implemented:
```
$ ./openttd-admin rcon -c /tmp/testgame/openttd.cfg
2020/12/27 21:35:53 connecting to localhost:3977...
2020/12/27 21:35:53 connected
> getseed
net> [admin] Rcon command from 'OpenTTD-Admin' (1.10.3): 'getseed'
> Generation Seed: 3277715519
> ^C
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
