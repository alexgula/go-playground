package helper

import (
	"bytes"
	"github.com/alexgula/go-playground/text/template02/model"
	"github.com/sipin/gorazor/gorazor"
)

func Msg(u *model.User) string {
	var _buffer bytes.Buffer

	username := u.Name

	if u.Email != "" {

		username += "(" + u.Email + ")"

	}

	_buffer.WriteString("\n\n<div class=\"welcome\">\n\n<h4>Hello ")
	_buffer.WriteString(gorazor.HTMLEscape(username))
	_buffer.WriteString("</h4>\n\n\n\n<div>")
	_buffer.WriteString((u.Intro))
	_buffer.WriteString("</div>\n\n</div>")

	return _buffer.String()
}
