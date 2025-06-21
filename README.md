# p2p

A simple peer to peer network that communicates over tcp. This is currently being
expanded upon to demonstrate various distributed data structures and mostly a pet
projects with peer-to-peer techniques.


## Getting started

This project uses Go modules and relies on the `golang.org/x/sys` package for
handling UNIX signals. Build the p2p binary with the standard Go toolchain:

```
$ git clone https://github.com/cpurta/p2p.git
$ cd p2p/src
$ go build -o p2p
$ mv p2p ../bin
```

## Running a single Peer node

To run a node for this peer2peer network just run the binary:

```
$ ./bin/p2p
```

You can test that the node is running and accepting requests by sending a TCP request
to the server running locally:

```
$ echo -n "PING" | nc 127.0.0.1 8888
PONG
```

## LICENSE

MIT
