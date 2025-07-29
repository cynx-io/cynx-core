package logger

import (
	"context"
	"fmt"
	coreContext "github.com/cynx-io/cynx-core/src/context"
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
	trxIndexName   string
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
	trxIndexName = "trx-" + cfg.ServiceName
}

func newLogEntry(ctx context.Context, logType string, args ...interface{}) LogEntry {
	return LogEntry{
		Timestamp: time.Now(),
		UserId:    coreContext.GetUserId(ctx),
		Username:  coreContext.GetKey(ctx, coreContext.KeyUsername),
		RequestId: coreContext.GetKeyOrEmpty(ctx, coreContext.KeyRequestId),
		Type:      logType,
		Message:   fmt.Sprint(args...),
	}
}

func Debug(ctx context.Context, args ...interface{}) {
	l.Debugln(args...)
	entry := newLogEntry(ctx, "DEBUG", args...)
	go func() {
		ctxBg := context.Background()
		if err := LogElasticsearch(ctxBg, debugIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Warn(ctx context.Context, args ...interface{}) {
	l.Warnln(args...)
	entry := newLogEntry(ctx, "WARN", args...)
	go func() {
		ctxBg := context.Background()
		if err := LogElasticsearch(ctxBg, warnIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Info(ctx context.Context, args ...interface{}) {
	l.Infoln(args...)
	entry := newLogEntry(ctx, "INFO", args...)
	go func() {
		ctxBg := context.Background()
		if err := LogElasticsearch(ctxBg, infoIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Error(ctx context.Context, args ...interface{}) {
	l.Errorln(args...)
	entry := newLogEntry(ctx, "ERROR", args...)
	go func() {
		ctxBg := context.Background()
		if err := LogElasticsearch(ctxBg, errorIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}

func Fatal(ctx context.Context, args ...interface{}) {
	l.Fatalln(args...)
	entry := newLogEntry(ctx, "FATAL", args...)
	go func() {
		ctxBg := context.Background()
		if err := LogElasticsearch(ctxBg, fatalIndexName, entry); err != nil {
			l.Error("Failed to log to Elasticsearch: ", err)
		}
	}()
}
