package logger

import (
	"context"
	"fmt"
	coreContext "github.com/cynxees/cynx-core/src/context"
	"github.com/elastic/go-elasticsearch"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"net/http"
	"time"
)

var (
	l              *logrus.Logger
	serviceName    string
	debugIndexName string
	infoIndexName  string
	warnIndexName  string
	errorIndexName string
	fatalIndexName string
)

func Init(cfg LoggerConfig) {
	l = logrus.New()
	l.SetFormatter(&ecslogrus.Formatter{
		DisableHTMLEscape: true,
		PrettyPrint:       true,
	})
	l.SetLevel(cfg.Level)

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: cfg.ElasticsearchURL,
		Transport: http.DefaultTransport,
	})
	if err != nil {
		panic("Failed to create Elasticsearch client: " + err.Error())
	}

	elasticClient = client
	serviceName = cfg.ServiceName
	debugIndexName = "debug-" + cfg.ServiceName
	infoIndexName = "info-" + cfg.ServiceName
	warnIndexName = "warn-" + cfg.ServiceName
	errorIndexName = "error-" + cfg.ServiceName
	fatalIndexName = "fatal-" + cfg.ServiceName
}

func Debug(ctx context.Context, args ...interface{}) {
	l.Debugln(args...)
	go func() {

		entry := LogEntry{
			Timestamp: time.Now(),
			UserId:    coreContext.GetUserId(ctx),
			Username:  coreContext.GetKey(ctx, coreContext.KeyUsername),
			RequestId: coreContext.GetKeyOrEmpty(ctx, coreContext.KeyRequestId),
			Type:      "DEBUG",
			Message:   fmt.Sprint(args...),
		}
		if err := LogElasticsearch(debugIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Warn(ctx context.Context, args ...interface{}) {
	l.Warnln(args...)
	go func() {

		entry := LogEntry{
			Timestamp: time.Now(),
			UserId:    coreContext.GetUserId(ctx),
			Username:  coreContext.GetKey(ctx, coreContext.KeyUsername),
			RequestId: coreContext.GetKeyOrEmpty(ctx, coreContext.KeyRequestId),
			Type:      "WARN",
			Message:   fmt.Sprint(args...),
		}
		if err := LogElasticsearch(warnIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Info(ctx context.Context, args ...interface{}) {
	l.Infoln(args...)
	go func() {

		entry := LogEntry{
			Timestamp: time.Now(),
			UserId:    coreContext.GetUserId(ctx),
			Username:  coreContext.GetKey(ctx, coreContext.KeyUsername),
			RequestId: coreContext.GetKeyOrEmpty(ctx, coreContext.KeyRequestId),
			Type:      "INFO",
			Message:   fmt.Sprint(args...),
		}
		if err := LogElasticsearch(infoIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Error(args ...interface{}) {
	l.Errorln(args...)
	go func() {
		entry := LogEntry{
			Timestamp: time.Now(),
			UserId:    coreContext.GetUserId(context.Background()),
			Username:  coreContext.GetKey(context.Background(), coreContext.KeyUsername),
			RequestId: coreContext.GetKeyOrEmpty(context.Background(), coreContext.KeyRequestId),
			Type:      "ERROR",
			Message:   fmt.Sprint(args...),
		}
		if err := LogElasticsearch(errorIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Fatal(args ...interface{}) {
	l.Fatalln(args...)
	go func() {
		entry := LogEntry{
			Timestamp: time.Now(),
			UserId:    coreContext.GetUserId(context.Background()),
			Username:  coreContext.GetKey(context.Background(), coreContext.KeyUsername),
			RequestId: coreContext.GetKeyOrEmpty(context.Background(), coreContext.KeyRequestId),
			Type:      "FATAL",
			Message:   fmt.Sprint(args...),
		}
		if err := LogElasticsearch(fatalIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
		panic(fmt.Sprint(args...))
	}()
}
