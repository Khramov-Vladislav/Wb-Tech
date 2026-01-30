package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"microservice/config"
	"microservice/internal/cache"
	"microservice/internal/db"
	"microservice/internal/kafka"
	"microservice/internal/services"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Database
	dbConn := db.Connect(cfg.GetDSN())
	defer dbConn.Close()
	dbRepo := db.NewPostgresOrderRepo(dbConn)

	// Cache
	orderCache := cache.NewCache(cfg.CacheTTL)
	cache.InitCacheFromDB(dbConn, orderCache)

	// Order service
	orderService := services.NewOrderService(dbRepo, orderCache)

	// Context (graceful)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Kafka consumer
	consumer := kafka.NewConsumer(
		cfg.KafkaBrokers,
		cfg.KafkaTopic,
		cfg.KafkaGroupID,
		dbRepo,
		orderCache,
	)
	defer consumer.Close()

	go consumer.Start(ctx)

	// Kafka producer
	producer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	defer producer.Close()

	// HTTP handlers
	mux := http.NewServeMux()

	// GET /order/{order_uid}
	mux.HandleFunc(cfg.OrderURLPrefix, func(w http.ResponseWriter, r *http.Request) {
		orderUID := strings.TrimPrefix(r.URL.Path, cfg.OrderURLPrefix)
		if orderUID == "" {
			http.Error(w, "order_uid не указан", http.StatusBadRequest)
			return
		}

		order, err := orderService.GetOrder(orderUID)
		if err != nil {
			http.Error(w, "заказ не найден", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(order); err != nil {
			log.Printf("Ошибка кодирования JSON: %v", err)
		}
	})

	// Static frontend
	mux.Handle("/", http.FileServer(http.Dir(cfg.StaticDir)))

	// HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	go func() {
		log.Println("HTTP-сервер запущен на http://localhost:" + cfg.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска HTTP сервера: %v", err)
		}
	}()

	// ========================
	// Graceful shutdown
	// ========================
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Получен сигнал остановки, завершаем сервис...")

	// Stop kafka consumer
	cancel()

	// Stop HTTP
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Ошибка при остановке HTTP сервера: %v", err)
	} else {
		log.Println("HTTP-сервер корректно остановлен")
	}

	log.Println("Сервис завершил работу")
}
