// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type codeRecorder struct {
	http.ResponseWriter
	status int
}

func (c *codeRecorder) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

// handleOnlyFormRequest handles onlyForm operation.
//
// POST /onlyForm
func (s *Server) handleOnlyFormRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("onlyForm"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/onlyForm"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), OnlyFormOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: OnlyFormOperation,
			ID:   "onlyForm",
		}
	)
	request, close, err := s.decodeOnlyFormRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *OnlyFormOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    OnlyFormOperation,
			OperationSummary: "",
			OperationID:      "onlyForm",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *OnlyFormReq
			Params   = struct{}
			Response = *OnlyFormOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.OnlyForm(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.OnlyForm(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOnlyFormResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleOnlyMultipartFileRequest handles onlyMultipartFile operation.
//
// POST /onlyMultipartFile
func (s *Server) handleOnlyMultipartFileRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("onlyMultipartFile"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/onlyMultipartFile"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), OnlyMultipartFileOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: OnlyMultipartFileOperation,
			ID:   "onlyMultipartFile",
		}
	)
	request, close, err := s.decodeOnlyMultipartFileRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *OnlyMultipartFileOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    OnlyMultipartFileOperation,
			OperationSummary: "",
			OperationID:      "onlyMultipartFile",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *OnlyMultipartFileReq
			Params   = struct{}
			Response = *OnlyMultipartFileOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.OnlyMultipartFile(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.OnlyMultipartFile(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOnlyMultipartFileResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleOnlyMultipartFormRequest handles onlyMultipartForm operation.
//
// POST /onlyMultipartForm
func (s *Server) handleOnlyMultipartFormRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("onlyMultipartForm"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/onlyMultipartForm"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), OnlyMultipartFormOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: OnlyMultipartFormOperation,
			ID:   "onlyMultipartForm",
		}
	)
	request, close, err := s.decodeOnlyMultipartFormRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *OnlyMultipartFormOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    OnlyMultipartFormOperation,
			OperationSummary: "",
			OperationID:      "onlyMultipartForm",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *OnlyMultipartFormReq
			Params   = struct{}
			Response = *OnlyMultipartFormOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.OnlyMultipartForm(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.OnlyMultipartForm(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOnlyMultipartFormResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestFormURLEncodedRequest handles testFormURLEncoded operation.
//
// POST /testFormURLEncoded
func (s *Server) handleTestFormURLEncodedRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testFormURLEncoded"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/testFormURLEncoded"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TestFormURLEncodedOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TestFormURLEncodedOperation,
			ID:   "testFormURLEncoded",
		}
	)
	request, close, err := s.decodeTestFormURLEncodedRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TestFormURLEncodedOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TestFormURLEncodedOperation,
			OperationSummary: "",
			OperationID:      "testFormURLEncoded",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *TestForm
			Params   = struct{}
			Response = *TestFormURLEncodedOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TestFormURLEncoded(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.TestFormURLEncoded(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestFormURLEncodedResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestMultipartRequest handles testMultipart operation.
//
// POST /testMultipart
func (s *Server) handleTestMultipartRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testMultipart"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/testMultipart"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TestMultipartOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TestMultipartOperation,
			ID:   "testMultipart",
		}
	)
	request, close, err := s.decodeTestMultipartRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TestMultipartOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TestMultipartOperation,
			OperationSummary: "",
			OperationID:      "testMultipart",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *TestFormMultipart
			Params   = struct{}
			Response = *TestMultipartOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TestMultipart(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.TestMultipart(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestMultipartResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestMultipartUploadRequest handles testMultipartUpload operation.
//
// POST /testMultipartUpload
func (s *Server) handleTestMultipartUploadRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testMultipartUpload"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/testMultipartUpload"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TestMultipartUploadOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TestMultipartUploadOperation,
			ID:   "testMultipartUpload",
		}
	)
	request, close, err := s.decodeTestMultipartUploadRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TestMultipartUploadOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TestMultipartUploadOperation,
			OperationSummary: "",
			OperationID:      "testMultipartUpload",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *TestMultipartUploadReq
			Params   = struct{}
			Response = *TestMultipartUploadOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.TestMultipartUpload(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.TestMultipartUpload(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestMultipartUploadResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestReuseFormOptionalSchemaRequest handles testReuseFormOptionalSchema operation.
//
// POST /testReuseFormOptionalSchema
func (s *Server) handleTestReuseFormOptionalSchemaRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testReuseFormOptionalSchema"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/testReuseFormOptionalSchema"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TestReuseFormOptionalSchemaOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TestReuseFormOptionalSchemaOperation,
			ID:   "testReuseFormOptionalSchema",
		}
	)
	request, close, err := s.decodeTestReuseFormOptionalSchemaRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TestReuseFormOptionalSchemaOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TestReuseFormOptionalSchemaOperation,
			OperationSummary: "",
			OperationID:      "testReuseFormOptionalSchema",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = OptSharedRequestMultipart
			Params   = struct{}
			Response = *TestReuseFormOptionalSchemaOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TestReuseFormOptionalSchema(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.TestReuseFormOptionalSchema(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestReuseFormOptionalSchemaResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestReuseFormSchemaRequest handles testReuseFormSchema operation.
//
// POST /testReuseFormSchema
func (s *Server) handleTestReuseFormSchemaRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testReuseFormSchema"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/testReuseFormSchema"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TestReuseFormSchemaOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TestReuseFormSchemaOperation,
			ID:   "testReuseFormSchema",
		}
	)
	request, close, err := s.decodeTestReuseFormSchemaRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TestReuseFormSchemaOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TestReuseFormSchemaOperation,
			OperationSummary: "",
			OperationID:      "testReuseFormSchema",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *SharedRequestMultipart
			Params   = struct{}
			Response = *TestReuseFormSchemaOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TestReuseFormSchema(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.TestReuseFormSchema(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestReuseFormSchemaResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestShareFormSchemaRequest handles testShareFormSchema operation.
//
// POST /testShareFormSchema
func (s *Server) handleTestShareFormSchemaRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("testShareFormSchema"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/testShareFormSchema"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), TestShareFormSchemaOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#status
			// Span Status MUST be left unset if HTTP status code was in the 1xx, 2xx or 3xx ranges,
			// unless there was another error (e.g., network error receiving the response body; or 3xx codes with
			// max redirects exceeded), in which case status MUST be set to Error.
			code := statusWriter.status
			if code < 100 || code >= 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: TestShareFormSchemaOperation,
			ID:   "testShareFormSchema",
		}
	)
	request, close, err := s.decodeTestShareFormSchemaRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *TestShareFormSchemaOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    TestShareFormSchemaOperation,
			OperationSummary: "",
			OperationID:      "testShareFormSchema",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = TestShareFormSchemaReq
			Params   = struct{}
			Response = *TestShareFormSchemaOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.TestShareFormSchema(ctx, request)
				return response, err
			},
		)
	} else {
		err = s.h.TestShareFormSchema(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeTestShareFormSchemaResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
