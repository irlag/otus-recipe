// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package responses

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

func easyjson3c6d3af4DecodeOtusRecipeAppApiResponses(in *jlexer.Lexer, out *RecipeUpdateOkResponse) {
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
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "cooking_time":
			out.CookingTime = int(in.Int())
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
func easyjson3c6d3af4EncodeOtusRecipeAppApiResponses(out *jwriter.Writer, in RecipeUpdateOkResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
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
		out.Int(int(in.CookingTime))
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
func (v RecipeUpdateOkResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c6d3af4EncodeOtusRecipeAppApiResponses(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipeUpdateOkResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c6d3af4EncodeOtusRecipeAppApiResponses(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipeUpdateOkResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c6d3af4DecodeOtusRecipeAppApiResponses(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipeUpdateOkResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c6d3af4DecodeOtusRecipeAppApiResponses(l, v)
}
