package httpclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/logger"
	"github.com/arshabbir/utils/model"
)

type httpClient struct {
	client *http.Client
	l      logger.Logger
	conf   config.Config
}

type HttpClient interface {
	Do(method string, url string, body io.Reader) ([]byte, error)
}

func NewHttpClient(conf config.Config, l logger.Logger) HttpClient {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(conf.HttpConnectinTimeout) * time.Second,
		}).Dial,
		TLSHandshakeTimeout: time.Duration(conf.HttpConnectinTimeout) * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * time.Duration(conf.HttpRequestTimeout),
		Transport: netTransport,
	}

	return &httpClient{client: netClient, l: l, conf: conf}
}

func (h *httpClient) Do(method string, url string, body io.Reader) ([]byte, error) {
	// Form the request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		h.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "httpClient", Message: err.Error()})
		return nil, err
	}

	// Perform the client.Do
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Http Error code : %d", resp.StatusCode)
	}
	// Return the
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return data, nil
}
