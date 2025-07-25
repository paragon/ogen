// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes AnyOfIntegerNumberString as json.
func (s AnyOfIntegerNumberString) Encode(e *jx.Encoder) {
	switch s.Type {
	case IntAnyOfIntegerNumberString:
		e.Int(s.Int)
	case Float64AnyOfIntegerNumberString:
		e.Float64(s.Float64)
	case StringAnyOfIntegerNumberString:
		e.Str(s.String)
	}
}

// Decode decodes AnyOfIntegerNumberString from json.
func (s *AnyOfIntegerNumberString) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AnyOfIntegerNumberString to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Number:
		num, err := d.Num()
		if err != nil {
			return errors.Wrap(err, "parse number")
		}
		if d := jx.DecodeBytes(num); num.IsInt() {
			v, err := d.Int()
			s.Int = int(v)
			if err != nil {
				return err
			}
			s.Type = IntAnyOfIntegerNumberString
		} else {
			v, err := d.Float64()
			s.Float64 = float64(v)
			if err != nil {
				return err
			}
			s.Type = Float64AnyOfIntegerNumberString
		}
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringAnyOfIntegerNumberString
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AnyOfIntegerNumberString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AnyOfIntegerNumberString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *IntegerNumber) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *IntegerNumber) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("plain")
		s.Plain.Encode(e)
	}
}

var jsonFieldsNameOfIntegerNumber = [1]string{
	0: "plain",
}

// Decode decodes IntegerNumber from json.
func (s *IntegerNumber) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IntegerNumber to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "plain":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Plain.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"plain\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode IntegerNumber")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfIntegerNumber) {
					name = jsonFieldsNameOfIntegerNumber[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *IntegerNumber) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IntegerNumber) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *JaegerAnyOf) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *JaegerAnyOf) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("medium")
		e.Str(s.Medium)
	}
	{
		e.FieldStart("sizeLimit")
		s.SizeLimit.Encode(e)
	}
}

var jsonFieldsNameOfJaegerAnyOf = [2]string{
	0: "medium",
	1: "sizeLimit",
}

// Decode decodes JaegerAnyOf from json.
func (s *JaegerAnyOf) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode JaegerAnyOf to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "medium":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Medium = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"medium\"")
			}
		case "sizeLimit":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.SizeLimit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sizeLimit\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode JaegerAnyOf")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfJaegerAnyOf) {
					name = jsonFieldsNameOfJaegerAnyOf[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *JaegerAnyOf) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *JaegerAnyOf) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes JaegerAnyOfSizeLimit as json.
func (s JaegerAnyOfSizeLimit) Encode(e *jx.Encoder) {
	switch s.Type {
	case IntJaegerAnyOfSizeLimit:
		e.Int(s.Int)
	case StringJaegerAnyOfSizeLimit:
		e.Str(s.String)
	}
}

// Decode decodes JaegerAnyOfSizeLimit from json.
func (s *JaegerAnyOfSizeLimit) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode JaegerAnyOfSizeLimit to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Number:
		v, err := d.Int()
		s.Int = int(v)
		if err != nil {
			return err
		}
		s.Type = IntJaegerAnyOfSizeLimit
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringJaegerAnyOfSizeLimit
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s JaegerAnyOfSizeLimit) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *JaegerAnyOfSizeLimit) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *OneUUID) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *OneUUID) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("owner")
		e.Str(s.Owner)
	}
	{
		e.FieldStart("version")
		e.Int32(s.Version)
	}
	{
		e.FieldStart("subscription_id")
		s.SubscriptionID.Encode(e)
	}
}

var jsonFieldsNameOfOneUUID = [3]string{
	0: "owner",
	1: "version",
	2: "subscription_id",
}

// Decode decodes OneUUID from json.
func (s *OneUUID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode OneUUID to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "owner":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Owner = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"owner\"")
			}
		case "version":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int32()
				s.Version = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"version\"")
			}
		case "subscription_id":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.SubscriptionID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"subscription_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode OneUUID")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfOneUUID) {
					name = jsonFieldsNameOfOneUUID[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *OneUUID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OneUUID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes OneUUIDSubscriptionID as json.
func (s OneUUIDSubscriptionID) Encode(e *jx.Encoder) {
	switch s.Type {
	case SubscriptionUUIDOneUUIDSubscriptionID:
		s.SubscriptionUUID.Encode(e)
	}
}

// Decode decodes OneUUIDSubscriptionID from json.
func (s *OneUUIDSubscriptionID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode OneUUIDSubscriptionID to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.String:
		if err := s.SubscriptionUUID.Decode(d); err != nil {
			return err
		}
		s.Type = SubscriptionUUIDOneUUIDSubscriptionID
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OneUUIDSubscriptionID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OneUUIDSubscriptionID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SubscriptionUUID as json.
func (s SubscriptionUUID) Encode(e *jx.Encoder) {
	switch s.Type {
	case UUIDv4SubscriptionUUID:
		s.UUIDv4.Encode(e)
	}
}

// Decode decodes SubscriptionUUID from json.
func (s *SubscriptionUUID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SubscriptionUUID to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.String:
		if err := s.UUIDv4.Decode(d); err != nil {
			return err
		}
		s.Type = UUIDv4SubscriptionUUID
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SubscriptionUUID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SubscriptionUUID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UUIDv4 as json.
func (s UUIDv4) Encode(e *jx.Encoder) {
	unwrapped := uuid.UUID(s)

	json.EncodeUUID(e, unwrapped)
}

// Decode decodes UUIDv4 from json.
func (s *UUIDv4) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UUIDv4 to nil")
	}
	var unwrapped uuid.UUID
	if err := func() error {
		v, err := json.DecodeUUID(d)
		unwrapped = v
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UUIDv4(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s UUIDv4) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UUIDv4) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
