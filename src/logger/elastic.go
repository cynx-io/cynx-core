package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"io"
)

var (
	elasticClient *elasticsearch.Client
)

func LogTrxElasticsearch(index string, entry TrxEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	Debug("Logging to Elasticsearch: ", string(data))
	res, err := elasticClient.Index(
		index,
		bytes.NewReader(data),
		elasticClient.Index.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Error("Error closing response body: ", err)
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}

func LogElasticsearch(index string, entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	Debug("Logging to Elasticsearch: ", string(data))
	res, err := elasticClient.Index(
		index,
		bytes.NewReader(data),
		elasticClient.Index.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Error("Error closing response body: ", err)
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}
