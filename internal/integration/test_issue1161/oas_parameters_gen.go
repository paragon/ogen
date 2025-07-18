// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// FooParamXyzGetParams is parameters of GET /foo/{param}/xyz operation.
type FooParamXyzGetParams struct {
	Param string
}

func unpackFooParamXyzGetParams(packed middleware.Parameters) (params FooParamXyzGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "param",
			In:   "path",
		}
		params.Param = packed[key].(string)
	}
	return params
}

func decodeFooParamXyzGetParams(args [1]string, argsEscaped bool, r *http.Request) (params FooParamXyzGetParams, _ error) {
	// Decode path: param.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "param",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Param = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "param",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
