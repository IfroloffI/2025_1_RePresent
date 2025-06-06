// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonC80ae7adDecodeRetargetInternalPayServiceEasyjsonModels(in *jlexer.Lexer, out *TransactionResponse) {
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
		case "transactionId":
			out.TransactionID = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "nextAction":
			out.NextAction = string(in.String())
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
func easyjsonC80ae7adEncodeRetargetInternalPayServiceEasyjsonModels(out *jwriter.Writer, in TransactionResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"transactionId\":"
		out.RawString(prefix[1:])
		out.String(string(in.TransactionID))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"nextAction\":"
		out.RawString(prefix)
		out.String(string(in.NextAction))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TransactionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeRetargetInternalPayServiceEasyjsonModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TransactionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeRetargetInternalPayServiceEasyjsonModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TransactionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeRetargetInternalPayServiceEasyjsonModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TransactionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeRetargetInternalPayServiceEasyjsonModels(l, v)
}
func easyjsonC80ae7adDecodeRetargetInternalPayServiceEasyjsonModels1(in *jlexer.Lexer, out *TopUpRequest) {
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
		case "amount":
			out.Amount = float64(in.Float64())
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
func easyjsonC80ae7adEncodeRetargetInternalPayServiceEasyjsonModels1(out *jwriter.Writer, in TopUpRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Amount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TopUpRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeRetargetInternalPayServiceEasyjsonModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopUpRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeRetargetInternalPayServiceEasyjsonModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopUpRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeRetargetInternalPayServiceEasyjsonModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopUpRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeRetargetInternalPayServiceEasyjsonModels1(l, v)
}
