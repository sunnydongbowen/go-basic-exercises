package test

// pprof http-web
import (
	"net/http"
	_ "net/http/pprof" // 开启pprof
	"time"
)

func getMMsg() string {
	time.Sleep(1 * time.Second)
	return "hello Go"
}

func homeHander(w http.ResponseWriter, r *http.Request) {
	msg := getMMsg()
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/", homeHander)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
