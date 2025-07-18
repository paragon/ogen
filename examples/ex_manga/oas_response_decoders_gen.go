// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeGetBookResponse(resp *http.Response) (res GetBookRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Book
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetBookForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPageCoverImageResponse(resp *http.Response) (res GetPageCoverImageRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("image/*", ct):
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := GetPageCoverImageOK{Data: bytes.NewReader(b)}
			var wrapper GetPageCoverImageOKHeaders
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Content-Type" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Content-Type",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							wrapper.ContentType = c
							return nil
						}); err != nil {
							return err
						}
					} else {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Content-Type header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetPageCoverImageForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPageImageResponse(resp *http.Response) (res GetPageImageRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("image/*", ct):
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := GetPageImageOK{Data: bytes.NewReader(b)}
			var wrapper GetPageImageOKHeaders
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Content-Type" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Content-Type",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							wrapper.ContentType = c
							return nil
						}); err != nil {
							return err
						}
					} else {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Content-Type header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetPageImageForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPageThumbnailImageResponse(resp *http.Response) (res GetPageThumbnailImageRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("image/*", ct):
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := GetPageThumbnailImageOK{Data: bytes.NewReader(b)}
			var wrapper GetPageThumbnailImageOKHeaders
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Content-Type" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Content-Type",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							wrapper.ContentType = c
							return nil
						}); err != nil {
							return err
						}
					} else {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Content-Type header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetPageThumbnailImageForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeSearchResponse(resp *http.Response) (res SearchRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response SearchOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &SearchForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeSearchByTagIDResponse(resp *http.Response) (res SearchByTagIDRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response SearchByTagIDOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &SearchByTagIDForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
