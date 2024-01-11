package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Создание метрик
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_counter",
			Help: "This is my counter metric",
		},
	)
	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "my_gauge",
			Help: "This is my gauge metric",
		},
	)
	histogram := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "my_histogram",
			Help:    "This is my histogram metric",
			Buckets: []float64{1, 2, 3, 4, 5},
		},
	)

	// Регистрация метрик в реестре
	prometheus.MustRegister(counter, gauge, histogram)

	// Запуск HTTP сервера для сбора метрик
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8080", nil)
	}()

	// Генерация метрик каждую секунду
	for {
		counter.Inc()
		gauge.Set(rand.Float64() * 100)
		histogram.Observe(rand.Float64() * 5)
		time.Sleep(time.Second)
	}
}
