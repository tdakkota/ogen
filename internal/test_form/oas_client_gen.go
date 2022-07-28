// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.SyncInt64().Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// TestFormURLEncoded invokes testFormURLEncoded operation.
//
// POST /testFormURLEncoded
func (c *Client) TestFormURLEncoded(ctx context.Context, request TestForm) (res TestFormURLEncodedOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testFormURLEncoded"),
	}
	// Validate request before sending.

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TestFormURLEncoded",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.serverURL)
	u.Path += "/testFormURLEncoded"

	stage = "EncodeRequest"
	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeTestFormURLEncodedRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTestFormURLEncodedResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TestMultipart invokes testMultipart operation.
//
// POST /testMultipart
func (c *Client) TestMultipart(ctx context.Context, request TestForm) (res TestMultipartOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testMultipart"),
	}
	// Validate request before sending.

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TestMultipart",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.serverURL)
	u.Path += "/testMultipart"

	stage = "EncodeRequest"
	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeTestMultipartRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTestMultipartResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TestMultipartUpload invokes testMultipartUpload operation.
//
// POST /testMultipartUpload
func (c *Client) TestMultipartUpload(ctx context.Context, request TestMultipartUploadReqForm) (res TestMultipartUploadOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testMultipartUpload"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TestMultipartUpload",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.serverURL)
	u.Path += "/testMultipartUpload"

	stage = "EncodeRequest"
	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeTestMultipartUploadRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTestMultipartUploadResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TestShareFormSchema invokes testShareFormSchema operation.
//
// POST /testShareFormSchema
func (c *Client) TestShareFormSchema(ctx context.Context, request TestShareFormSchemaReq) (res TestShareFormSchemaOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testShareFormSchema"),
	}
	// Validate request before sending.
	switch request := request.(type) {
	case *SharedRequest:
		// Validation is not required for this type.
	case *SharedRequestForm:
		// Validation is not required for this type.
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TestShareFormSchema",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.serverURL)
	u.Path += "/testShareFormSchema"

	stage = "EncodeRequest"
	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeTestShareFormSchemaRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTestShareFormSchemaResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}