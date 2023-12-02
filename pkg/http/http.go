package http

import (
	"bytes"
	"context"
	"github.com/ZBIGBEAR/openai-go/common"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"

	"time"
)

type Config struct {
	Authorization string
	RetryMax      int
	RetryWaitMin  time.Duration
	RetryWaitMax  time.Duration
	Timeout       time.Duration
}

type Option func(cfg *Config)

func TimeOutOption(timeout time.Duration) Option {
	return func(cfg *Config) {
		cfg.Timeout = timeout
	}
}

type Http interface {
	Get(ctx context.Context, header map[string]string, url string, body []byte) (*http.Response, error)
	Post(ctx context.Context, header map[string]string, url string, body []byte) (*http.Response, error)
}

type innerHttp struct {
	cfg    *Config
	log    common.Logger
	client *retryablehttp.Client
}

var _ Http = &innerHttp{}

func New(cfg *Config, opts ...Option) Http {
	for i := range opts {
		opts[i](cfg)
	}

	client := retryablehttp.NewClient()
	client.RetryMax = cfg.RetryMax
	client.RetryWaitMin = cfg.RetryWaitMin
	client.RetryWaitMax = cfg.RetryWaitMax
	client.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return false, err
	}
	client.Logger = nil
	return &innerHttp{
		cfg:    cfg,
		log:    common.NewLog(),
		client: client,
	}
}

func (i *innerHttp) Get(ctx context.Context, header map[string]string, url string, body []byte) (*http.Response, error) {
	return i.do(ctx, Get, header, url, body)
}

func (i *innerHttp) Post(ctx context.Context, header map[string]string, url string, body []byte) (*http.Response, error) {
	return i.do(ctx, Post, header, url, body)
}

func (i *innerHttp) do(ctx context.Context, method string, header map[string]string, url string, body []byte) (*http.Response, error) {
	begin := time.Now()
	defer func() {
		if i.log != nil {
			i.log.Infof(ctx, "[innerHttp.do] method:%s, url:%s, time cost:%fs\n", method, url, time.Since(begin).Seconds())
		}
	}()

	req, err := retryablehttp.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentType, "application/json")
	if i.cfg != nil {
		if i.cfg.Authorization != "" {
			req.Header.Set(Authorization, i.cfg.Authorization)
		}
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return i.client.Do(req)
}
