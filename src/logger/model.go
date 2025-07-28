package logger

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
)

type LoggerConfig struct {
	ServiceName      string
	ElasticsearchURL []string
	Level            logrus.Level
}

type TrxEntry struct {
	Timestamp     time.Time       `json:"timestamp"`
	UserId        *int32          `json:"userId"`
	Username      *string         `json:"username"`
	UserType      *int32          `json:"userType"`
	RequestId     string          `json:"requestId"`
	RequestOrigin string          `json:"requestOrigin"` // frontend website ex www.makeadle.com
	IpAddress     string          `json:"ipAddress"`     // client ip
	Endpoint      string          `json:"endpoint"`      // POST /api/v1/resource
	Host          string          `json:"host"`          // e.g. api.myservice.com
	Referer       string          `json:"referer,omitempty"`
	UserAgent     string          `json:"userAgent,omitempty"` // browser or bot details
	Type          string          `json:"type,omitempty"`      // e.g. "request", "response"
	Body          json.RawMessage `json:"body,omitempty"`      // request or response body
}

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	UserId    *int32    `json:"userId"`
	Username  *string   `json:"username"`
	UserType  *int32    `json:"userType"`
	RequestId string    `json:"requestId"`
	Type      string    `json:"type,omitempty"`    // e.g. "debug", "warn", "error", "info"
	Message   string    `json:"message,omitempty"` // message
}
