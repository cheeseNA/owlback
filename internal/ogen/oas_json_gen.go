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

// Encode encodes TaskRequest as json.
func (o OptTaskRequest) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes TaskRequest from json.
func (o *OptTaskRequest) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTaskRequest to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTaskRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTaskRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TaskRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TaskRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("site_url")
		json.EncodeURI(e, s.SiteURL)
	}
	{
		e.FieldStart("duration_day")
		e.Int(s.DurationDay)
	}
	{
		e.FieldStart("condition_query")
		e.Str(s.ConditionQuery)
	}
	{
		e.FieldStart("is_public")
		e.Bool(s.IsPublic)
	}
}

var jsonFieldsNameOfTaskRequest = [4]string{
	0: "site_url",
	1: "duration_day",
	2: "condition_query",
	3: "is_public",
}

// Decode decodes TaskRequest from json.
func (s *TaskRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TaskRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "site_url":
			requiredBitSet[0] |= 1 << 0
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
		case "duration_day":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.DurationDay = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"duration_day\"")
			}
		case "condition_query":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.ConditionQuery = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"condition_query\"")
			}
		case "is_public":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Bool()
				s.IsPublic = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_public\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TaskRequest")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
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
				if fieldIdx < len(jsonFieldsNameOfTaskRequest) {
					name = jsonFieldsNameOfTaskRequest[fieldIdx]
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
func (s *TaskRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TaskRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TaskResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TaskResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("site_url")
		json.EncodeURI(e, s.SiteURL)
	}
	{
		e.FieldStart("duration_day")
		e.Int(s.DurationDay)
	}
	{
		e.FieldStart("condition_query")
		e.Str(s.ConditionQuery)
	}
	{
		e.FieldStart("is_public")
		e.Bool(s.IsPublic)
	}
	{
		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{
		e.FieldStart("created_at")
		json.EncodeDateTime(e, s.CreatedAt)
	}
	{
		e.FieldStart("created_by")
		json.EncodeUUID(e, s.CreatedBy)
	}
	{
		e.FieldStart("updated_at")
		json.EncodeDateTime(e, s.UpdatedAt)
	}
}

var jsonFieldsNameOfTaskResponse = [8]string{
	0: "site_url",
	1: "duration_day",
	2: "condition_query",
	3: "is_public",
	4: "id",
	5: "created_at",
	6: "created_by",
	7: "updated_at",
}

// Decode decodes TaskResponse from json.
func (s *TaskResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TaskResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "site_url":
			requiredBitSet[0] |= 1 << 0
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
		case "duration_day":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.DurationDay = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"duration_day\"")
			}
		case "condition_query":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.ConditionQuery = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"condition_query\"")
			}
		case "is_public":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Bool()
				s.IsPublic = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_public\"")
			}
		case "id":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "created_at":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.CreatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "created_by":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.CreatedBy = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_by\"")
			}
		case "updated_at":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.UpdatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TaskResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b11111111,
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
				if fieldIdx < len(jsonFieldsNameOfTaskResponse) {
					name = jsonFieldsNameOfTaskResponse[fieldIdx]
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
func (s *TaskResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TaskResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}