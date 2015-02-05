package layout

import (
	"bytes"
)

func Base(body string, title string, side string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n\n<html>\n\n<head>\n\n<meta charset=\"utf-8\" />")
	_buffer.WriteString((title))
	_buffer.WriteString("\n\n</head>\n\n<body>\n\n<div>")
	_buffer.WriteString((body))
	_buffer.WriteString("</div>\n\n<div>")
	_buffer.WriteString((side))
	_buffer.WriteString("</div>\n\n</body>\n\n</html>")

	return _buffer.String()
}
