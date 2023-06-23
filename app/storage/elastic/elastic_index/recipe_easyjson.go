// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package elastic_index

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonAf04a44aDecodeOtusRecipeAppStorageElasticElasticIndex(in *jlexer.Lexer, out *Recipe) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "cooking_time":
			out.CookingTime = int16(in.Int16())
		case "calories":
			if in.IsNull() {
				in.Skip()
				out.Calories = nil
			} else {
				if out.Calories == nil {
					out.Calories = new(int32)
				}
				*out.Calories = int32(in.Int32())
			}
		case "proteins":
			if in.IsNull() {
				in.Skip()
				out.Proteins = nil
			} else {
				if out.Proteins == nil {
					out.Proteins = new(int32)
				}
				*out.Proteins = int32(in.Int32())
			}
		case "fats":
			if in.IsNull() {
				in.Skip()
				out.Fats = nil
			} else {
				if out.Fats == nil {
					out.Fats = new(int32)
				}
				*out.Fats = int32(in.Int32())
			}
		case "carbohydrates":
			if in.IsNull() {
				in.Skip()
				out.Carbohydrates = nil
			} else {
				if out.Carbohydrates == nil {
					out.Carbohydrates = new(int32)
				}
				*out.Carbohydrates = int32(in.Int32())
			}
		case "version":
			out.Version = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAf04a44aEncodeOtusRecipeAppStorageElasticElasticIndex(out *jwriter.Writer, in Recipe) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"cooking_time\":"
		out.RawString(prefix)
		out.Int16(int16(in.CookingTime))
	}
	{
		const prefix string = ",\"calories\":"
		out.RawString(prefix)
		if in.Calories == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Calories))
		}
	}
	{
		const prefix string = ",\"proteins\":"
		out.RawString(prefix)
		if in.Proteins == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Proteins))
		}
	}
	{
		const prefix string = ",\"fats\":"
		out.RawString(prefix)
		if in.Fats == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Fats))
		}
	}
	{
		const prefix string = ",\"carbohydrates\":"
		out.RawString(prefix)
		if in.Carbohydrates == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Carbohydrates))
		}
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Recipe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAf04a44aEncodeOtusRecipeAppStorageElasticElasticIndex(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Recipe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAf04a44aEncodeOtusRecipeAppStorageElasticElasticIndex(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Recipe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAf04a44aDecodeOtusRecipeAppStorageElasticElasticIndex(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Recipe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAf04a44aDecodeOtusRecipeAppStorageElasticElasticIndex(l, v)
}
