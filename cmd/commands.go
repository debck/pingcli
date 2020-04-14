package cmd

import (
	"time"
    "os"
    "log"
    "fmt"
    "net"
    "golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)
var total int 
var fail int
const Protocol = 1
var AddressListen = "0.0.0.0"


func setupPing(t time.Time) {
	ping := func(addr string){
		dst, dur, err := sendPing(addr)
		total++
        if err != nil {
			log.Printf("Ping %s (%s): %s\n", addr, dst, err)
			fail++
            return
		}
		log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
	}
	ping(value)
}




func sendPing(addr string) (*net.IPAddr, time.Duration, error) {

    c, err := icmp.ListenPacket("ip4:icmp", AddressListen)
    if err != nil {
        return nil, 0, err
    }
    defer c.Close()

    dst, err := net.ResolveIPAddr("ip4", addr)
    if err != nil {
        return nil, 0, err
    }

    // ICMP message
    m := icmp.Message{
        Type: ipv4.ICMPTypeEcho, Code: 0,
        Body: &icmp.Echo{
            ID: os.Getpid() & 0xffff, Seq: 1,
            Data: []byte(""),
        },
    }
    b, err := m.Marshal(nil)
    if err != nil {
        return dst, 0, err
    }

    // Send
    start := time.Now()
    n, err := c.WriteTo(b, dst)
    if err != nil {
        return dst, 0, err
    } else if n != len(b) {
        return dst, 0, fmt.Errorf("got %v; want %v", n, len(b))
    }

    // Reply
    reply := make([]byte, 1500)
    err = c.SetReadDeadline(time.Now().Add(10 * time.Second))
    if err != nil {
        return dst, 0, err
    }
    n, peer, err := c.ReadFrom(reply)
    if err != nil {
        return dst, 0, err
    }
    duration := time.Since(start)


    rm, err := icmp.ParseMessage(Protocol, reply[:n])
    if err != nil {
        return dst, 0, err
    }
    switch rm.Type {
    case ipv4.ICMPTypeEchoReply:
        return dst, duration, nil
    default:
        return dst, 0, fmt.Errorf("got %+v from %v; want echo reply", rm, peer)
    }
}

