// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"

	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreatePetRequestJSON(
	req CreatePetReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeCreatePetCategoriesRequestJSON(
	req CreatePetCategoriesReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeCreatePetFriendsRequestJSON(
	req CreatePetFriendsReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeCreatePetOwnerRequestJSON(
	req CreatePetOwnerReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeUpdatePetRequestJSON(
	req UpdatePetReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}