// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"mime"
	"mime/multipart"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeCreateAnswerRequest(
	req *CreateAnswerRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateChatCompletionRequest(
	req *CreateChatCompletionRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateClassificationRequest(
	req *CreateClassificationRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateCompletionRequest(
	req *CreateCompletionRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateEditRequest(
	req *CreateEditRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateEmbeddingRequest(
	req *CreateEmbeddingRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateFileRequest(
	req *CreateFileRequestMultipart,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "purpose" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "purpose",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Purpose))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.File.WriteMultipart("file", w); err != nil {
			return errors.Wrap(err, "write \"file\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeCreateFineTuneRequest(
	req *CreateFineTuneRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateImageRequest(
	req *CreateImageRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateImageEditRequest(
	req *CreateImageEditRequestMultipart,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "prompt" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "prompt",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Prompt))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "n" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "n",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.N.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "size" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "size",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Size.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "response_format" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "response_format",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.ResponseFormat.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.User.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.Image.WriteMultipart("image", w); err != nil {
			return errors.Wrap(err, "write \"image\"")
		}
		if val, ok := request.Mask.Get(); ok {
			if err := val.WriteMultipart("mask", w); err != nil {
				return errors.Wrap(err, "write \"mask\"")
			}
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeCreateImageVariationRequest(
	req *CreateImageVariationRequestMultipart,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "n" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "n",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.N.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "size" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "size",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Size.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "response_format" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "response_format",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.ResponseFormat.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.User.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.Image.WriteMultipart("image", w); err != nil {
			return errors.Wrap(err, "write \"image\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeCreateModerationRequest(
	req *CreateModerationRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateSearchRequest(
	req *CreateSearchRequest,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateTranscriptionRequest(
	req *CreateTranscriptionRequestMultipart,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "model" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "model",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Model))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "prompt" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "prompt",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Prompt.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "response_format" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "response_format",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.ResponseFormat.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "temperature" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "temperature",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Temperature.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "language" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "language",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Language.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.File.WriteMultipart("file", w); err != nil {
			return errors.Wrap(err, "write \"file\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeCreateTranslationRequest(
	req *CreateTranslationRequestMultipart,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "model" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "model",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Model))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "prompt" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "prompt",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Prompt.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "response_format" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "response_format",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.ResponseFormat.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "temperature" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "temperature",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Temperature.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.File.WriteMultipart("file", w); err != nil {
			return errors.Wrap(err, "write \"file\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}
