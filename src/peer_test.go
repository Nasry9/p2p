package main

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

// startTestPeer starts a peer listening on an ephemeral port.
// It returns the peer and the dial address.
func startTestPeer(t *testing.T) (*Peer, string) {
	peers = 10
	p := NewPeer(10)
	p.PeerLock = &sync.Mutex{}

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen failed: %v", err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()

	p.Port = fmt.Sprintf("%d", port)
	go p.Start()
	time.Sleep(100 * time.Millisecond)
	return p, fmt.Sprintf("127.0.0.1:%d", port)
}

func stopPeer(p *Peer, addr string) {
	p.Shutdown()
	net.Dial("tcp", addr)
	time.Sleep(100 * time.Millisecond)
}

func TestPeerPingEcho(t *testing.T) {
	p, addr := startTestPeer(t)
	defer stopPeer(p, addr)

	// ping
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("dial failed: %v", err)
	}
	conn.Write([]byte("PING"))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if string(buf[:n]) != "PONG" {
		t.Fatalf("expected PONG got %q", string(buf[:n]))
	}
	conn.Close()

	// echo
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("dial failed: %v", err)
	}
	conn.Write([]byte("ECHO"))
	n, err = conn.Read(buf)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if string(buf[:n]) != "ECHO" {
		t.Fatalf("expected ECHO got %q", string(buf[:n]))
	}
	conn.Close()
}

func TestAddRemovePeer(t *testing.T) {
	peers = 2
	p := NewPeer(2)
	p.PeerLock = &sync.Mutex{}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")

	if err := p.AddPeer("peer1", addr); err != nil {
		t.Fatalf("AddPeer failed: %v", err)
	}
	if len(p.Peers) != 1 {
		t.Fatalf("expected 1 peer got %d", len(p.Peers))
	}

	if err := p.RemovePeer("peer1"); err != nil {
		t.Fatalf("RemovePeer failed: %v", err)
	}
	if len(p.Peers) != 0 {
		t.Fatalf("expected 0 peers got %d", len(p.Peers))
	}
}
