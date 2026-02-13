package handlers

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jcostabe/go-demo-3/model"
	"github.com/sirupsen/logrus"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	logrus.Infof("Alive method")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
	logrus.Infof("Elapsed time of /isAlive: %v", time.Since(start))

}

func GetHostInfo(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	logrus.Infof("Host info")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	host, _ := os.Hostname()
	io.WriteString(w, host)
	logrus.Infof("Elapsed time of /info: %v", time.Since(start))

}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	cfg := model.DefaultConfiguration()

	logrus.Infof("Version info")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, cfg.ServiceConfig.Version)
	logrus.Infof("Elapsed time of /version: %v", time.Since(start))

}

func RandomError(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	code := http.StatusOK
	clientConnections.WithLabelValues(cfg.ServiceConfig.Name).Inc()
	defer func() { recordMetrics(r.URL.Path, r.Method, code, start) }()

	switch randErr := rand.Intn(3); randErr {
	case 1:
		code = http.StatusNotFound
		w.WriteHeader(code)
		w.Write([]byte("404 - Not found"))
	case 2:
		code = http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write([]byte("500 - Internal server error"))
	default:
		w.WriteHeader(code)
		w.Write([]byte("200 - It works fine!"))

	}

	logrus.Infof("Elapsed time of %s: %v", r.URL.Path, time.Since(start))
}
