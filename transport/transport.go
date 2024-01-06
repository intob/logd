package transport

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/swissinfo-ch/logd/auth"
	"github.com/swissinfo-ch/logd/msg"
	"github.com/swissinfo-ch/logd/pack"
)

const (
	PingPeriod            = time.Second
	KickAfterMissingPings = 10
	bufferSize            = 2048
)

type Sub struct {
	raddr    *net.UDPAddr
	lastPing time.Time
}

type Transporter struct {
	In          chan []byte
	Out         chan []byte
	subs        map[string]*Sub
	mu          sync.Mutex
	readSecret  []byte
	writeSecret []byte
}

func NewTransporter() *Transporter {
	return &Transporter{
		In:   make(chan []byte),
		Out:  make(chan []byte, 10),
		subs: make(map[string]*Sub),
		mu:   sync.Mutex{},
	}
}

func (t *Transporter) SetReadSecret(secret []byte) {
	t.readSecret = secret
}

func (t *Transporter) SetWriteSecret(secret []byte) {
	t.writeSecret = secret
}

func (t *Transporter) Listen(ctx context.Context, laddr string) {
	l, err := net.ResolveUDPAddr("udp", laddr)
	if err != nil {
		panic(fmt.Errorf("resolve laddr err: %w", err))
	}
	conn, err := net.ListenUDP("udp", l)
	if err != nil {
		panic(fmt.Errorf("listen udp err: %w", err))
	}
	defer conn.Close()
	fmt.Println("listening udp on", conn.LocalAddr())
	go t.readFromConn(ctx, conn)
	go t.writeToConn(ctx, conn)
	<-ctx.Done()
	fmt.Println("stopped listening udp")
}

func (t *Transporter) readFromConn(ctx context.Context, conn *net.UDPConn) {
	var buf []byte
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf = make([]byte, bufferSize)
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			n, raddr, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}
			sum, timeBytes, payload, err := pack.UnpackSignedMsg(buf[:n])
			if err != nil {
				fmt.Println("unpack msg err:", err)
				continue
			}
			payloadStr := string(payload)
			if payloadStr == "tail" || payloadStr == "ping" {
				valid, err := auth.Verify(t.readSecret, sum, timeBytes, payload)
				if !valid || err != nil {
					fmt.Printf("%s unauthorised: %s\r\n", raddr.IP.String(), err)
					continue
				}
				if string(payload) == "tail" {
					go t.handleTailer(conn, raddr)
					continue
				}
				if string(payload) == "ping" {
					go t.handlePing(raddr)
					continue
				}
			}
			valid, err := auth.Verify(t.writeSecret, sum, timeBytes, payload)
			if !valid || err != nil {
				fmt.Printf("%s unauthorised: %s\r\n", raddr.IP.String(), err)
				continue
			}
			t.In <- payload
		}
	}
}

func (t *Transporter) handleTailer(conn *net.UDPConn, raddr *net.UDPAddr) {
	t.mu.Lock()
	t.subs[raddr.AddrPort().String()] = &Sub{
		raddr:    raddr,
		lastPing: time.Now(),
	}
	t.mu.Unlock()
	// notify sub directly
	payload, err := pack.PackMsg(&msg.Msg{
		Fn:  "logd",
		Msg: "tailing logs...",
	})
	if err != nil {
		fmt.Println("pack msg err:", err)
	}
	_, err = conn.WriteToUDP(payload, raddr)
	if err != nil {
		fmt.Printf("write udp err: (%s) %s\r\n", raddr, err)
	}
}

func (t *Transporter) handlePing(raddr *net.UDPAddr) {
	t.mu.Lock()
	sub := t.subs[raddr.AddrPort().String()]
	if sub != nil {
		sub.lastPing = time.Now()
	}
	t.mu.Unlock()
}

func (t *Transporter) writeToConn(ctx context.Context, conn *net.UDPConn) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-t.Out:
			for raddr, sub := range t.subs {
				if sub.lastPing.Before(time.Now().Add(-(PingPeriod * KickAfterMissingPings))) {
					t.kickSub(conn, sub, raddr)
					continue
				}
				go func(msg []byte, sub *Sub, raddr string) {
					_, err := conn.WriteToUDP(msg, sub.raddr)
					if err != nil {
						fmt.Printf("write udp err: (%s) %s\r\n", raddr, err)
					}
				}(msg, sub, raddr)
			}
		}
	}
}

// kickSub removes sub from map & broadcasts kick
func (t *Transporter) kickSub(conn *net.UDPConn, sub *Sub, raddr string) {
	t.mu.Lock()
	delete(t.subs, raddr)
	t.mu.Unlock()
	t.broadcast(fmt.Sprintf("kicked %s, no ping received, timed out", raddr))
	fmt.Printf("kicked %s, no ping received, timed out\r\n", raddr)
	// notify sub directly
	payload, err := pack.PackMsg(&msg.Msg{
		Fn:  "logd",
		Msg: "you've been kicked, ping timed out",
	})
	if err != nil {
		fmt.Println("pack msg err:", err)
	}
	_, err = conn.WriteToUDP(payload, sub.raddr)
	if err != nil {
		fmt.Printf("write udp err: (%s) %s\r\n", raddr, err)
	}
}

func (t *Transporter) broadcast(m string) error {
	payload, err := pack.PackMsg(&msg.Msg{
		Fn:  "logd",
		Msg: m,
	})
	if err != nil {
		return fmt.Errorf("pack msg err: %w", err)
	}
	t.Out <- payload
	return nil
}
