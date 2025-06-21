package main

import (
	"net"
	"sync"
	"testing"
)

func TestHandleRequest_Ping(t *testing.T) {
	peers = 1
	p := NewPeer(1)
	p.PeerLock = &sync.Mutex{}

	client, server := net.Pipe()
	go p.Handler.HandleRequest(server)

	if _, err := client.Write([]byte("PING")); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := client.Read(buf)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	if string(buf[:n]) != "PONG" {
		t.Fatalf("expected PONG got %q", string(buf[:n]))
	}
}

func TestHandleRequest_Echo(t *testing.T) {
	peers = 1
	p := NewPeer(1)
	p.PeerLock = &sync.Mutex{}

	client, server := net.Pipe()
	go p.Handler.HandleRequest(server)

	msg := "ECHO"
	if _, err := client.Write([]byte(msg)); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := client.Read(buf)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	if string(buf[:n]) != msg {
		t.Fatalf("expected %q got %q", msg, string(buf[:n]))
	}
}
