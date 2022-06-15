// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeDefaultTestRequest(r *http.Request, span trace.Span) (
	req DefaultTest,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request DefaultTest
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeFoobarPostRequest(r *http.Request, span trace.Span) (
	req OptPet,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptPet
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeOneofBugRequest(r *http.Request, span trace.Span) (
	req OneOfBugs,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request OneOfBugs
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePetCreateRequest(r *http.Request, span trace.Span) (
	req OptPet,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptPet
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePetUpdateNameAliasPostRequest(r *http.Request, span trace.Span) (
	req OptPetName,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptPetName
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePetUpdateNamePostRequest(r *http.Request, span trace.Span) (
	req OptString,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request OptString
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			request.Reset()
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := (validate.String{
						MinLength:    6,
						MinLengthSet: true,
						MaxLength:    0,
						MaxLengthSet: false,
						Email:        false,
						Hostname:     false,
						Regex:        nil,
					}).Validate(string(request.Value)); err != nil {
						return errors.Wrap(err, "string")
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodePetUploadAvatarByIDRequest(r *http.Request, span trace.Span) (
	req PetUploadAvatarByIDReq,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/octet-stream":
		return PetUploadAvatarByIDReq{Data: r.Body}, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestFloatValidationRequest(r *http.Request, span trace.Span) (
	req TestFloatValidation,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request TestFloatValidation
		buf := new(bytes.Buffer)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, close, err
		}

		if written == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}

		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestFormURLEncodedRequest(r *http.Request, span trace.Span) (
	req TestForm,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseForm(); err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}
		form := r.PostForm

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotIDVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.ID.SetTo(reqDotIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotUUIDVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						reqDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.UUID.SetTo(reqDotUUIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.Description = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"description\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "array",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var reqDotArrayVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							reqDotArrayVal = c
							return nil
						}(); err != nil {
							return err
						}
						req.Array = append(req.Array, reqDotArrayVal)
						return nil
					})
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"array\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "object",
				Style:   uri.QueryStyleForm,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotObjectVal TestFormObject
					if err := func() error {
						return reqDotObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					req.Object.SetTo(reqDotObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"object\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "deepObject",
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotDeepObjectVal TestFormDeepObject
					if err := func() error {
						return reqDotDeepObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					req.DeepObject.SetTo(reqDotDeepObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"deepObject\"")
				}
			}
		}

		return req, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestMultipartRequest(r *http.Request, span trace.Span) (
	req TestForm,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotIDVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.ID.SetTo(reqDotIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotUUIDVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						reqDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.UUID.SetTo(reqDotUUIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					req.Description = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"description\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "array",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var reqDotArrayVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							reqDotArrayVal = c
							return nil
						}(); err != nil {
							return err
						}
						req.Array = append(req.Array, reqDotArrayVal)
						return nil
					})
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"array\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "object",
				Style:   uri.QueryStyleForm,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotObjectVal TestFormObject
					if err := func() error {
						return reqDotObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					req.Object.SetTo(reqDotObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"object\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "deepObject",
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotDeepObjectVal TestFormDeepObject
					if err := func() error {
						return reqDotDeepObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					req.DeepObject.SetTo(reqDotDeepObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"deepObject\"")
				}
			}
		}

		return req, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestMultipartUploadRequest(r *http.Request, span trace.Span) (
	req TestMultipartUploadReq,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "orderId",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotOrderIdVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotOrderIdVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.OrderId.SetTo(reqDotOrderIdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"orderId\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "userId",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}

			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var reqDotUserIdVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						reqDotUserIdVal = c
						return nil
					}(); err != nil {
						return err
					}
					req.UserId.SetTo(reqDotUserIdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"userId\"")
				}
			}
		}
		if err := func() error {
			files, ok := r.MultipartForm.File["fileName"]
			if !ok || len(files) < 1 {
				return errors.New("file is not set")
			}
			header := files[0]

			f, err := header.Open()
			if err != nil {
				return errors.Wrap(err, "open")
			}
			closers = append(closers, f)

			req.FileName = ht.MultipartFile{
				Name: header.Filename,
				File: f,
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "file \"fileName\"")
		}

		return req, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}