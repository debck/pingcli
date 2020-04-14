package cmd

import (
	"time"
    "os"
    "fmt"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)
var value string

var RootCmd = &cobra.Command{
	Use:   "pingcli [arg]",
	Short: "Ping CLI application.It accepts a hostname or an IP address as its argument,\nthen send ICMP echo request in a loop to the target while receiving messages.",
	Run: func(cmd *cobra.Command, args []string) {
		value = args[0]
		SetupCloseHandler()
		doEvery(2000*time.Millisecond, setupPing)
	},
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Packet->  sent:",total," |  Recived Successfully:", total-fail)
		fmt.Println("Loss%:",(fail/total)*100)
		os.Exit(0)
	}()
}


func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}