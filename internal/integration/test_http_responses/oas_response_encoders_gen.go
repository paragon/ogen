// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeAnyContentTypeBinaryStringSchemaResponse(response *AnyContentTypeBinaryStringSchemaOKHeaders, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "Content-Type" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "Content-Type",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.ContentType))
			}); err != nil {
				return errors.Wrap(err, "encode Content-Type header")
			}
		}
	}
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if closer, ok := response.Response.Data.(io.Closer); ok {
		defer closer.Close()
	}
	if _, err := io.Copy(writer, response.Response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeAnyContentTypeBinaryStringSchemaDefaultResponse(response *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "Content-Type" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "Content-Type",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.ContentType))
			}); err != nil {
				return errors.Wrap(err, "encode Content-Type header")
			}
		}
	}
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	writer := w
	if closer, ok := response.Response.Data.(io.Closer); ok {
		defer closer.Close()
	}
	if _, err := io.Copy(writer, response.Response); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodeCombinedResponse(response CombinedRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CombinedOK:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *Combined2XXStatusCode:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		if st := http.StatusText(code); code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := new(jx.Encoder)
		e.Int(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	case *Combined5XXStatusCode:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		if st := http.StatusText(code); code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := new(jx.Encoder)
		e.Bool(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	case *CombinedDefStatusCode:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		if st := http.StatusText(code); code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := new(jx.Encoder)
		e.ArrStart()
		for _, elem := range response.Response {
			e.Str(elem)
		}
		e.ArrEnd()
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeHeaders200Response(response *Headers200OK, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "X-Test-Header" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Test-Header",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.XTestHeader))
			}); err != nil {
				return errors.Wrap(err, "encode X-Test-Header header")
			}
		}
	}
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeHeadersCombinedResponse(response HeadersCombinedRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *HeadersCombinedOK:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "X-Test-Header" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Test-Header",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.XTestHeader))
				}); err != nil {
					return errors.Wrap(err, "encode X-Test-Header header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *HeadersCombined4XX:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "X-Test-Header" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Test-Header",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.XTestHeader))
				}); err != nil {
					return errors.Wrap(err, "encode X-Test-Header header")
				}
			}
		}
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		if st := http.StatusText(code); code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	case *HeadersCombinedDef:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "X-Test-Header" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Test-Header",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.XTestHeader))
				}); err != nil {
					return errors.Wrap(err, "encode X-Test-Header header")
				}
			}
		}
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		if st := http.StatusText(code); code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeHeadersDefaultResponse(response *HeadersDefaultDef, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "X-Test-Header" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Test-Header",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.XTestHeader))
			}); err != nil {
				return errors.Wrap(err, "encode X-Test-Header header")
			}
		}
	}
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodeHeadersJSONResponse(response *HeadersJSONOK, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "X-Json-Custom-Header" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Json-Custom-Header",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				var enc jx.Encoder
				func(e *jx.Encoder) {

					if len(response.XJSONCustomHeader) != 0 {
						e.Raw(response.XJSONCustomHeader)
					}
				}(&enc)
				return e.EncodeValue(string(enc.Bytes()))
			}); err != nil {
				return errors.Wrap(err, "encode X-Json-Custom-Header header")
			}
		}
		// Encode "X-Json-Header" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Json-Header",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				var enc jx.Encoder
				func(e *jx.Encoder) {

					response.XJSONHeader.Encode(e)
				}(&enc)
				return e.EncodeValue(string(enc.Bytes()))
			}); err != nil {
				return errors.Wrap(err, "encode X-Json-Header header")
			}
		}
	}
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeHeadersPatternResponse(response *HeadersPattern4XX, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "X-Test-Header" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Test-Header",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.XTestHeader))
			}); err != nil {
				return errors.Wrap(err, "encode X-Test-Header header")
			}
		}
	}
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodeIntersectPatternCodeResponse(response IntersectPatternCodeRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *IntersectPatternCodeOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *IntersectPatternCode2XXStatusCode:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		code := response.StatusCode
		if code == 0 {
			// Set default status code.
			code = http.StatusOK
		}
		w.WriteHeader(code)
		if st := http.StatusText(code); code >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}

		e := new(jx.Encoder)
		e.Int(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		if code >= http.StatusInternalServerError {
			return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeMultipleGenericResponsesResponse(response MultipleGenericResponsesRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *NilInt:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NilString:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeOctetStreamBinaryStringSchemaResponse(response OctetStreamBinaryStringSchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if closer, ok := response.Data.(io.Closer); ok {
		defer closer.Close()
	}
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeOctetStreamEmptySchemaResponse(response OctetStreamEmptySchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if closer, ok := response.Data.(io.Closer); ok {
		defer closer.Close()
	}
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeOptionalHeadersResponse(response *OptionalHeadersOK, w http.ResponseWriter, span trace.Span) error {
	// Encoding response headers.
	{
		h := uri.NewHeaderEncoder(w.Header())
		// Encode "X-Optional" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Optional",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := response.XOptional.Get(); ok {
					return e.EncodeValue(conv.StringToString(val))
				}
				return nil
			}); err != nil {
				return errors.Wrap(err, "encode X-Optional header")
			}
		}
		// Encode "X-Required" header.
		{
			cfg := uri.HeaderParameterEncodingConfig{
				Name:    "X-Required",
				Explode: false,
			}
			if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(response.XRequired))
			}); err != nil {
				return errors.Wrap(err, "encode X-Required header")
			}
		}
	}
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	return nil
}

func encodeStreamJSONResponse(response StreamJSONRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *QueryData:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.NewStreamingEncoder(w, -1)
		response.Encode(e)
		if err := e.Close(); err != nil {
			return errors.Wrap(err, "flush streaming")
		}

		return nil

	case *Error:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeTextPlainBinaryStringSchemaResponse(response TextPlainBinaryStringSchemaOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if closer, ok := response.Data.(io.Closer); ok {
		defer closer.Close()
	}
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}
