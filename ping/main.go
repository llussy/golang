package main

import (
	"fmt"
	"os"
	"time"

	"github.com/paulstuart/ping"
	"github.com/sirupsen/logrus"
)

const PingTimes = 10
const IntervalPackage = 50

var PingLive bool = true


func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
}
func main() {

	loss := Ping("10.21.8.105",1)
	if loss > 90.0 {
		logrus.Errorf("ping liveness false, ignore this data")
	}

}

func Ping(address string, timeout int) float64{
	var FailedCount float64
	for i := 1; i <= PingTimes; i++ {
		err := ping.Pinger(address,timeout)
		if err != nil {
			logrus.Error("ping error ",err)
			FailedCount++
		}
		time.Sleep(time.Duration(IntervalPackage) * time.Millisecond)
	}
	fmt.Println(FailedCount)
	return (FailedCount / PingTimes) * 100

}