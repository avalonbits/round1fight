// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

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

func easyjsonC1cedd36DecodeGithubComAvalonbitsRound1fightEndpointsApi(in *jlexer.Lexer, out *personJSON) {
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
		case "apelido":
			out.Nickname = string(in.String())
		case "nome":
			out.Name = string(in.String())
		case "nascimento":
			if in.IsNull() {
				in.Skip()
				out.Birthday = nil
			} else {
				if out.Birthday == nil {
					out.Birthday = new(jsonDate)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Birthday).UnmarshalJSON(data))
				}
			}
		case "stack":
			if in.IsNull() {
				in.Skip()
				out.Stack = nil
			} else {
				in.Delim('[')
				if out.Stack == nil {
					if !in.IsDelim(']') {
						out.Stack = make([]string, 0, 4)
					} else {
						out.Stack = []string{}
					}
				} else {
					out.Stack = (out.Stack)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Stack = append(out.Stack, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC1cedd36EncodeGithubComAvalonbitsRound1fightEndpointsApi(out *jwriter.Writer, in personJSON) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"apelido\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"nome\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"nascimento\":"
		out.RawString(prefix)
		if in.Birthday == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Birthday).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"stack\":"
		out.RawString(prefix)
		if in.Stack == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Stack {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v personJSON) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComAvalonbitsRound1fightEndpointsApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v personJSON) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComAvalonbitsRound1fightEndpointsApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *personJSON) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComAvalonbitsRound1fightEndpointsApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *personJSON) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComAvalonbitsRound1fightEndpointsApi(l, v)
}
