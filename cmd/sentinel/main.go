package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2pizzzza/sentinetAgent/internal/collector/metrics"
	"github.com/2pizzzza/sentinetAgent/internal/config"
	"github.com/2pizzzza/sentinetAgent/internal/core"
	"github.com/2pizzzza/sentinetAgent/pkg/logger"
	"github.com/redis/go-redis/v9"
)

func main() {

	cnf, err := config.New("config/config.yml")
	if err != nil {
		panic(err)
	}

	log := logger.New(cnf.Env)

	application := core.New(log, *cnf)

	linuxMetrics, err := metrics.NewLinuxMetrics()
	if err != nil {
		panic(err)
	}

	metricsCh := make(chan *metrics.Metrics)

	go linuxMetrics.StartCollecting(100*time.Millisecond, metricsCh)
	_ = metricsCh

	// go SaveMetrics(ctx, log, metricsCh, redisConn, postgresConn)

	go application.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	log.Info("Server is dead")
}


func SaveMetrics(ch chan *metrics.Metrics, doneCh chan struct{}){
	redis, err := redis.SearchFieldTypeVector("metrics")
	resultCh := make(chan int, 1)
if err != nil {
	fmt.Errorf("failde to save metrics to redis")
	return
}
select {
  doneCh: return
	resultCh <- ch
}
}
