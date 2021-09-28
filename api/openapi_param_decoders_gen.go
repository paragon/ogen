// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

func DecodeFoobarGetParams(r *http.Request) (FoobarGetParams, error) {
	var params FoobarGetParams
	{
		param := r.URL.Query().Get("inlinedParam")
		if len(param) == 0 {
			return params, fmt.Errorf("query param 'inlinedParam' is empty")
		}

		v, err := conv.ToInt64(param)
		if err != nil {
			return params, fmt.Errorf("parse query param 'inlinedParam': %w", err)
		}

		params.Query.InlinedParam = v
	}
	{
		param := r.URL.Query().Get("skip")
		if len(param) == 0 {
			return params, fmt.Errorf("query param 'skip' is empty")
		}

		v, err := conv.ToInt32(param)
		if err != nil {
			return params, fmt.Errorf("parse query param 'skip': %w", err)
		}

		params.Query.Skip = v
	}

	return params, nil
}

func DecodePetGetParams(r *http.Request) (PetGetParams, error) {
	var params PetGetParams
	{
		c, err := r.Cookie("token")
		if err != nil {
			return params, fmt.Errorf("get cookie 'token': %w", err)
		}

		param := c.Value
		if len(param) == 0 {
			return params, fmt.Errorf("cookie param 'token' is empty")
		}

		v, err := conv.ToString(param)
		if err != nil {
			return params, fmt.Errorf("parse cookie param 'token': %w", err)
		}

		params.Cookie.Token = v
	}
	{
		param := r.Header.Values("x-scope")
		if len(param) == 0 {
			return params, fmt.Errorf("header param 'x-scope' is empty")
		}

		v, err := conv.ToStringArray(param)
		if err != nil {
			return params, fmt.Errorf("parse header param 'x-scope': %w", err)
		}

		params.Header.XScope = v
	}
	{
		param := r.URL.Query().Get("petID")
		if len(param) == 0 {
			return params, fmt.Errorf("query param 'petID' is empty")
		}

		v, err := conv.ToInt64(param)
		if err != nil {
			return params, fmt.Errorf("parse query param 'petID': %w", err)
		}

		params.Query.PetID = v
	}

	return params, nil
}

func DecodePetGetByNameParams(r *http.Request) (PetGetByNameParams, error) {
	var params PetGetByNameParams
	{
		param := chi.URLParam(r, "name")
		if len(param) == 0 {
			return params, fmt.Errorf("path param 'name' is empty")
		}

		v, err := conv.ToString(param)
		if err != nil {
			return params, fmt.Errorf("parse path param 'name': %w", err)
		}

		params.Path.Name = v
	}

	return params, nil
}