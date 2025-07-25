// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Bird) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Bird) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("petType")
		e.Str(s.PetType)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("species")
		e.Str(s.Species)
	}
	{
		if s.CanTalk.Set {
			e.FieldStart("canTalk")
			s.CanTalk.Encode(e)
		}
	}
}

var jsonFieldsNameOfBird = [4]string{
	0: "petType",
	1: "name",
	2: "species",
	3: "canTalk",
}

// Decode decodes Bird from json.
func (s *Bird) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Bird to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "petType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.PetType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"petType\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "species":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Species = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"species\"")
			}
		case "canTalk":
			if err := func() error {
				s.CanTalk.Reset()
				if err := s.CanTalk.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"canTalk\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Bird")
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
				if fieldIdx < len(jsonFieldsNameOfBird) {
					name = jsonFieldsNameOfBird[fieldIdx]
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
func (s *Bird) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Bird) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Car) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Car) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("vehicleType")
		e.Str(s.VehicleType)
	}
	{
		e.FieldStart("make")
		e.Str(s.Make)
	}
	{
		e.FieldStart("model")
		e.Str(s.Model)
	}
	{
		if s.Doors.Set {
			e.FieldStart("doors")
			s.Doors.Encode(e)
		}
	}
}

var jsonFieldsNameOfCar = [4]string{
	0: "vehicleType",
	1: "make",
	2: "model",
	3: "doors",
}

// Decode decodes Car from json.
func (s *Car) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Car to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "vehicleType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.VehicleType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleType\"")
			}
		case "make":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Make = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"make\"")
			}
		case "model":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Model = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"model\"")
			}
		case "doors":
			if err := func() error {
				s.Doors.Reset()
				if err := s.Doors.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"doors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Car")
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
				if fieldIdx < len(jsonFieldsNameOfCar) {
					name = jsonFieldsNameOfCar[fieldIdx]
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
func (s *Car) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Car) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Cat) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Cat) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("petType")
		e.Str(s.PetType)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("breed")
		e.Str(s.Breed)
	}
	{
		if s.MeowVolume.Set {
			e.FieldStart("meowVolume")
			s.MeowVolume.Encode(e)
		}
	}
}

var jsonFieldsNameOfCat = [4]string{
	0: "petType",
	1: "name",
	2: "breed",
	3: "meowVolume",
}

// Decode decodes Cat from json.
func (s *Cat) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Cat to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "petType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.PetType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"petType\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "breed":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Breed = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"breed\"")
			}
		case "meowVolume":
			if err := func() error {
				s.MeowVolume.Reset()
				if err := s.MeowVolume.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"meowVolume\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Cat")
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
				if fieldIdx < len(jsonFieldsNameOfCat) {
					name = jsonFieldsNameOfCat[fieldIdx]
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
func (s *Cat) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Cat) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Dog) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Dog) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("petType")
		e.Str(s.PetType)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("breed")
		e.Str(s.Breed)
	}
	{
		if s.BarkLoudness.Set {
			e.FieldStart("barkLoudness")
			s.BarkLoudness.Encode(e)
		}
	}
}

var jsonFieldsNameOfDog = [4]string{
	0: "petType",
	1: "name",
	2: "breed",
	3: "barkLoudness",
}

// Decode decodes Dog from json.
func (s *Dog) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Dog to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "petType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.PetType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"petType\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "breed":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Breed = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"breed\"")
			}
		case "barkLoudness":
			if err := func() error {
				s.BarkLoudness.Reset()
				if err := s.BarkLoudness.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"barkLoudness\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Dog")
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
				if fieldIdx < len(jsonFieldsNameOfDog) {
					name = jsonFieldsNameOfDog[fieldIdx]
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
func (s *Dog) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Dog) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *EmailNotification) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *EmailNotification) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("notificationType")
		e.Str(s.NotificationType)
	}
	{
		e.FieldStart("recipient")
		e.Str(s.Recipient)
	}
	{
		e.FieldStart("subject")
		e.Str(s.Subject)
	}
	{
		if s.Body.Set {
			e.FieldStart("body")
			s.Body.Encode(e)
		}
	}
}

var jsonFieldsNameOfEmailNotification = [4]string{
	0: "notificationType",
	1: "recipient",
	2: "subject",
	3: "body",
}

// Decode decodes EmailNotification from json.
func (s *EmailNotification) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode EmailNotification to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "notificationType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.NotificationType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notificationType\"")
			}
		case "recipient":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Recipient = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"recipient\"")
			}
		case "subject":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Subject = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"subject\"")
			}
		case "body":
			if err := func() error {
				s.Body.Reset()
				if err := s.Body.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"body\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode EmailNotification")
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
				if fieldIdx < len(jsonFieldsNameOfEmailNotification) {
					name = jsonFieldsNameOfEmailNotification[fieldIdx]
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
func (s *EmailNotification) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *EmailNotification) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Motorcycle) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Motorcycle) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("vehicleType")
		e.Str(s.VehicleType)
	}
	{
		e.FieldStart("make")
		e.Str(s.Make)
	}
	{
		e.FieldStart("model")
		e.Str(s.Model)
	}
	{
		if s.EngineSize.Set {
			e.FieldStart("engineSize")
			s.EngineSize.Encode(e)
		}
	}
}

var jsonFieldsNameOfMotorcycle = [4]string{
	0: "vehicleType",
	1: "make",
	2: "model",
	3: "engineSize",
}

// Decode decodes Motorcycle from json.
func (s *Motorcycle) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Motorcycle to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "vehicleType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.VehicleType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleType\"")
			}
		case "make":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Make = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"make\"")
			}
		case "model":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Model = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"model\"")
			}
		case "engineSize":
			if err := func() error {
				s.EngineSize.Reset()
				if err := s.EngineSize.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"engineSize\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Motorcycle")
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
				if fieldIdx < len(jsonFieldsNameOfMotorcycle) {
					name = jsonFieldsNameOfMotorcycle[fieldIdx]
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
func (s *Motorcycle) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Motorcycle) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Notification as json.
func (s Notification) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s Notification) encodeFields(e *jx.Encoder) {
	switch s.Type {
	case NotificationEmailNotification, NotificationMailNotification:
		switch s.Type {
		case NotificationEmailNotification:
			e.FieldStart("notificationType")
			e.Str("email")
		case NotificationMailNotification:
			e.FieldStart("notificationType")
			e.Str("mail")
		}
		{
			s := s.EmailNotification
			{
				e.FieldStart("recipient")
				e.Str(s.Recipient)
			}
			{
				e.FieldStart("subject")
				e.Str(s.Subject)
			}
			{
				if s.Body.Set {
					e.FieldStart("body")
					s.Body.Encode(e)
				}
			}
		}
	case NotificationSMSNotification, NotificationTextNotification:
		switch s.Type {
		case NotificationSMSNotification:
			e.FieldStart("notificationType")
			e.Str("sms")
		case NotificationTextNotification:
			e.FieldStart("notificationType")
			e.Str("text")
		}
		{
			s := s.SMSNotification
			{
				e.FieldStart("phoneNumber")
				e.Str(s.PhoneNumber)
			}
			{
				e.FieldStart("message")
				e.Str(s.Message)
			}
		}
	case NotificationMobileNotification, NotificationPushNotification:
		switch s.Type {
		case NotificationMobileNotification:
			e.FieldStart("notificationType")
			e.Str("mobile")
		case NotificationPushNotification:
			e.FieldStart("notificationType")
			e.Str("push")
		}
		{
			s := s.PushNotification
			{
				e.FieldStart("deviceId")
				e.Str(s.DeviceId)
			}
			{
				e.FieldStart("title")
				e.Str(s.Title)
			}
			{
				if s.Body.Set {
					e.FieldStart("body")
					s.Body.Encode(e)
				}
			}
			{
				if s.Badge.Set {
					e.FieldStart("badge")
					s.Badge.Encode(e)
				}
			}
		}
	}
}

// Decode decodes Notification from json.
func (s *Notification) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Notification to nil")
	}
	// Sum type discriminator.
	if typ := d.Next(); typ != jx.Object {
		return errors.Errorf("unexpected json type %q", typ)
	}

	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			if found {
				return d.Skip()
			}
			switch string(key) {
			case "notificationType":
				typ, err := d.Str()
				if err != nil {
					return err
				}
				switch typ {
				case "email":
					s.Type = NotificationEmailNotification
					found = true
				case "mail":
					s.Type = NotificationMailNotification
					found = true
				case "sms":
					s.Type = NotificationSMSNotification
					found = true
				case "text":
					s.Type = NotificationTextNotification
					found = true
				case "mobile":
					s.Type = NotificationMobileNotification
					found = true
				case "push":
					s.Type = NotificationPushNotification
					found = true
				default:
					return errors.Errorf("unknown type %s", typ)
				}
				return nil
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case NotificationEmailNotification, NotificationMailNotification:
		if err := s.EmailNotification.Decode(d); err != nil {
			return err
		}
	case NotificationSMSNotification, NotificationTextNotification:
		if err := s.SMSNotification.Decode(d); err != nil {
			return err
		}
	case NotificationMobileNotification, NotificationPushNotification:
		if err := s.PushNotification.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Notification) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Notification) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes float64 as json.
func (o OptFloat64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Float64(float64(o.Value))
}

// Decode decodes float64 from json.
func (o *OptFloat64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptFloat64 to nil")
	}
	o.Set = true
	v, err := d.Float64()
	if err != nil {
		return err
	}
	o.Value = float64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptFloat64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptFloat64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Pet as json.
func (s Pet) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s Pet) encodeFields(e *jx.Encoder) {
	switch s.Type {
	case PetCaninePet, PetDogPet:
		switch s.Type {
		case PetCaninePet:
			e.FieldStart("petType")
			e.Str("canine")
		case PetDogPet:
			e.FieldStart("petType")
			e.Str("dog")
		}
		{
			s := s.Dog
			{
				e.FieldStart("name")
				e.Str(s.Name)
			}
			{
				e.FieldStart("breed")
				e.Str(s.Breed)
			}
			{
				if s.BarkLoudness.Set {
					e.FieldStart("barkLoudness")
					s.BarkLoudness.Encode(e)
				}
			}
		}
	case PetCatPet, PetFelinePet:
		switch s.Type {
		case PetCatPet:
			e.FieldStart("petType")
			e.Str("cat")
		case PetFelinePet:
			e.FieldStart("petType")
			e.Str("feline")
		}
		{
			s := s.Cat
			{
				e.FieldStart("name")
				e.Str(s.Name)
			}
			{
				e.FieldStart("breed")
				e.Str(s.Breed)
			}
			{
				if s.MeowVolume.Set {
					e.FieldStart("meowVolume")
					s.MeowVolume.Encode(e)
				}
			}
		}
	case PetAvianPet, PetBirdPet:
		switch s.Type {
		case PetAvianPet:
			e.FieldStart("petType")
			e.Str("avian")
		case PetBirdPet:
			e.FieldStart("petType")
			e.Str("bird")
		}
		{
			s := s.Bird
			{
				e.FieldStart("name")
				e.Str(s.Name)
			}
			{
				e.FieldStart("species")
				e.Str(s.Species)
			}
			{
				if s.CanTalk.Set {
					e.FieldStart("canTalk")
					s.CanTalk.Encode(e)
				}
			}
		}
	}
}

// Decode decodes Pet from json.
func (s *Pet) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Pet to nil")
	}
	// Sum type discriminator.
	if typ := d.Next(); typ != jx.Object {
		return errors.Errorf("unexpected json type %q", typ)
	}

	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			if found {
				return d.Skip()
			}
			switch string(key) {
			case "petType":
				typ, err := d.Str()
				if err != nil {
					return err
				}
				switch typ {
				case "canine":
					s.Type = PetCaninePet
					found = true
				case "dog":
					s.Type = PetDogPet
					found = true
				case "cat":
					s.Type = PetCatPet
					found = true
				case "feline":
					s.Type = PetFelinePet
					found = true
				case "avian":
					s.Type = PetAvianPet
					found = true
				case "bird":
					s.Type = PetBirdPet
					found = true
				default:
					return errors.Errorf("unknown type %s", typ)
				}
				return nil
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case PetCaninePet, PetDogPet:
		if err := s.Dog.Decode(d); err != nil {
			return err
		}
	case PetCatPet, PetFelinePet:
		if err := s.Cat.Decode(d); err != nil {
			return err
		}
	case PetAvianPet, PetBirdPet:
		if err := s.Bird.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Pet) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Pet) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PushNotification) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PushNotification) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("notificationType")
		e.Str(s.NotificationType)
	}
	{
		e.FieldStart("deviceId")
		e.Str(s.DeviceId)
	}
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		if s.Body.Set {
			e.FieldStart("body")
			s.Body.Encode(e)
		}
	}
	{
		if s.Badge.Set {
			e.FieldStart("badge")
			s.Badge.Encode(e)
		}
	}
}

var jsonFieldsNameOfPushNotification = [5]string{
	0: "notificationType",
	1: "deviceId",
	2: "title",
	3: "body",
	4: "badge",
}

// Decode decodes PushNotification from json.
func (s *PushNotification) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PushNotification to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "notificationType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.NotificationType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notificationType\"")
			}
		case "deviceId":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.DeviceId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"deviceId\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "body":
			if err := func() error {
				s.Body.Reset()
				if err := s.Body.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"body\"")
			}
		case "badge":
			if err := func() error {
				s.Badge.Reset()
				if err := s.Badge.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"badge\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PushNotification")
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
				if fieldIdx < len(jsonFieldsNameOfPushNotification) {
					name = jsonFieldsNameOfPushNotification[fieldIdx]
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
func (s *PushNotification) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PushNotification) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SMSNotification) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SMSNotification) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("notificationType")
		e.Str(s.NotificationType)
	}
	{
		e.FieldStart("phoneNumber")
		e.Str(s.PhoneNumber)
	}
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
}

var jsonFieldsNameOfSMSNotification = [3]string{
	0: "notificationType",
	1: "phoneNumber",
	2: "message",
}

// Decode decodes SMSNotification from json.
func (s *SMSNotification) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SMSNotification to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "notificationType":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.NotificationType = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notificationType\"")
			}
		case "phoneNumber":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.PhoneNumber = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"phoneNumber\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SMSNotification")
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
				if fieldIdx < len(jsonFieldsNameOfSMSNotification) {
					name = jsonFieldsNameOfSMSNotification[fieldIdx]
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
func (s *SMSNotification) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SMSNotification) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Vehicle as json.
func (s Vehicle) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s Vehicle) encodeFields(e *jx.Encoder) {
	switch s.Type {
	case CarVehicle:
		e.FieldStart("vehicleType")
		e.Str("Car")
		{
			s := s.Car
			{
				e.FieldStart("make")
				e.Str(s.Make)
			}
			{
				e.FieldStart("model")
				e.Str(s.Model)
			}
			{
				if s.Doors.Set {
					e.FieldStart("doors")
					s.Doors.Encode(e)
				}
			}
		}
	case MotorcycleVehicle:
		e.FieldStart("vehicleType")
		e.Str("Motorcycle")
		{
			s := s.Motorcycle
			{
				e.FieldStart("make")
				e.Str(s.Make)
			}
			{
				e.FieldStart("model")
				e.Str(s.Model)
			}
			{
				if s.EngineSize.Set {
					e.FieldStart("engineSize")
					s.EngineSize.Encode(e)
				}
			}
		}
	}
}

// Decode decodes Vehicle from json.
func (s *Vehicle) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Vehicle to nil")
	}
	// Sum type discriminator.
	if typ := d.Next(); typ != jx.Object {
		return errors.Errorf("unexpected json type %q", typ)
	}

	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			if found {
				return d.Skip()
			}
			switch string(key) {
			case "vehicleType":
				typ, err := d.Str()
				if err != nil {
					return err
				}
				switch typ {
				case "Car":
					s.Type = CarVehicle
					found = true
				case "Motorcycle":
					s.Type = MotorcycleVehicle
					found = true
				default:
					return errors.Errorf("unknown type %s", typ)
				}
				return nil
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case CarVehicle:
		if err := s.Car.Decode(d); err != nil {
			return err
		}
	case MotorcycleVehicle:
		if err := s.Motorcycle.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Vehicle) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Vehicle) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
