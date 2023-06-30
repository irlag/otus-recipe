// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson6938619eDecodeOtusRecipeAppModels(in *jlexer.Lexer, out *RecipeUpdated) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson6938619eEncodeOtusRecipeAppModels(out *jwriter.Writer, in RecipeUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipeUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6938619eEncodeOtusRecipeAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipeUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6938619eEncodeOtusRecipeAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipeUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6938619eDecodeOtusRecipeAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipeUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6938619eDecodeOtusRecipeAppModels(l, v)
}
