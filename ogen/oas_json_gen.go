// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes uuid.UUID as json.
func (o OptUUID) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	json.EncodeUUID(e, o.Value)
}

// Decode decodes uuid.UUID from json.
func (o *OptUUID) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUUID to nil")
	}
	o.Set = true
	v, err := json.DecodeUUID(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUUID) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUUID) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Task) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Task) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		e.FieldStart("site_url")
		json.EncodeURI(e, s.SiteURL)
	}
	{
		e.FieldStart("condition")
		e.Str(s.Condition)
	}
	{
		e.FieldStart("duration_day")
		e.Int32(s.DurationDay)
	}
}

var jsonFieldsNameOfTask = [4]string{
	0: "id",
	1: "site_url",
	2: "condition",
	3: "duration_day",
}

// Decode decodes Task from json.
func (s *Task) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Task to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "site_url":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeURI(d)
				s.SiteURL = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_url\"")
			}
		case "condition":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Condition = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"condition\"")
			}
		case "duration_day":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int32()
				s.DurationDay = int32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"duration_day\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Task")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001110,
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
				if fieldIdx < len(jsonFieldsNameOfTask) {
					name = jsonFieldsNameOfTask[fieldIdx]
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
func (s *Task) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Task) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
