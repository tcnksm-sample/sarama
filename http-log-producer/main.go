package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {

	brokers := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()

	http.HandleFunc("/", KafkaLogProducerWrapper(producer, HelloHandler))

	log.Printf("[INFO] start server on :8080")
	http.ListenAndServe(":8080", nil)
	return 0
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!\n"))
}

// KafkaLogger is HandlerWrapper function to send log to Apache kafka
func KafkaLogProducerWrapper(producer sarama.AsyncProducer, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s %s%s", r.UserAgent(), r.Method, r.URL.Host, r.URL.Path)

		// Should execute handler first
		fn(w, r)

		logEncoder := &AccessLogEncoder{
			Time:   strconv.Itoa(int(time.Now().Unix())),
			Method: r.Method,
			Host:   r.Host,
			Path:   r.RequestURI,
			IP:     r.RemoteAddr,
		}

		msg := &sarama.ProducerMessage{
			Topic: "important",
			Key:   sarama.StringEncoder("app_hello"),
			Value: logEncoder,
		}

		producer.Input() <- msg
	}
}

// AccessLog is entry for KafkaLogProducer
// Need to implement sarama.Encoder interface
type AccessLogEncoder struct {
	Time   string `json:"time"`
	Method string `json:"method"`
	Host   string `json:"host"`
	Path   string `json:"path"`
	IP     string `json:"ip"`
}

func (a *AccessLogEncoder) Encode() ([]byte, error) {
	return json.Marshal(a)
}

func (a *AccessLogEncoder) Length() int {
	encoded, _ := json.Marshal(a)
	return len(encoded)
}
