package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/m1sol/api-tester/internal/config"
	"io"
	"net/http"
	"time"
)

type Executor struct {
	client *http.Client
}

type Result struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	LatencyMs  int64
	Err        error
}

func NewExecutor() *Executor {
	return &Executor{
		client: &http.Client{},
	}
}

func (e *Executor) Execute(ctx context.Context, req config.RequestConfig) Result {
	if e.client == nil {
		e.client = &http.Client{}
	}
	bodyBytes, _ := json.Marshal(req.Body)
	fmt.Println("Request body JSON:", string(bodyBytes))
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, bytes.NewReader(bodyBytes))
	if err != nil {
		return Result{Err: err}
	}

	// Добавляем headers
	httpReq.Header.Add("Content-Type", req.Headers.ContentType)
	//if req.Headers.Authorization != "" {
	//	httpReq.Header.Add("Authorization", req.Headers.Authorization)
	//}

	// Засекаем время запроса
	start := time.Now()
	resp, err := e.client.Do(httpReq)
	latency := time.Since(start).Milliseconds()

	if err != nil {
		return Result{Err: err, LatencyMs: latency}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Err: err, LatencyMs: latency}
	}

	return Result{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
		LatencyMs:  latency,
		Err:        nil,
	}
}
