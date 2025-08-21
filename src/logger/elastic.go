package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"io"
)

var (
	elasticClient *elasticsearch.Client
)

func LogTrxElasticsearch(ctx context.Context, entry TrxEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	res, err := elasticClient.Index(
		trxIndexName,
		bytes.NewReader(data),
		elasticClient.Index.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}

func LogElasticsearch(ctx context.Context, index string, entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

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
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}
