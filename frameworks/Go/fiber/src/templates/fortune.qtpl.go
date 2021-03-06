// Code generated by qtc from "fortune.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line frameworks/Go/fiber/src/templates/fortune.qtpl:1
package templates

//line frameworks/Go/fiber/src/templates/fortune.qtpl:1
import (
	qtio353123 "io"

	qt343634 "github.com/valyala/quicktemplate"
)

//line frameworks/Go/fiber/src/templates/fortune.qtpl:1
var (
	_ = qtio353123.Copy
	_ = qt343634.AcquireByteBuffer
)

//line frameworks/Go/fiber/src/templates/fortune.qtpl:2
type Fortune struct {
	ID      int32
	Message string
}

//line frameworks/Go/fiber/src/templates/fortune.qtpl:8
func StreamFortunePage(qw271823 *qt343634.Writer, rows []Fortune) {
//line frameworks/Go/fiber/src/templates/fortune.qtpl:8
	qw271823.N().S(`<!DOCTYPE html>
<html>
<head>
<title>Fortunes</title>
</head>
<body>
<table>
<tr><th>id</th><th>message</th></tr>
`)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:16
	for _, r := range rows {
//line frameworks/Go/fiber/src/templates/fortune.qtpl:16
		qw271823.N().S(`
<tr><td>`)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:17
		qw271823.N().D(int(r.ID))
//line frameworks/Go/fiber/src/templates/fortune.qtpl:17
		qw271823.N().S(`</td><td>`)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:17
		qw271823.E().S(r.Message)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:17
		qw271823.N().S(`</td></tr>
`)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:18
	}
//line frameworks/Go/fiber/src/templates/fortune.qtpl:18
	qw271823.N().S(`
</table>
</body>
</html>
`)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
}

//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
func WriteFortunePage(qq182734 qtio353123.Writer, rows []Fortune) {
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	qw271823 := qt343634.AcquireWriter(qq182734)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	StreamFortunePage(qw271823, rows)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	qt343634.ReleaseWriter(qw271823)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
}

//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
func FortunePage(rows []Fortune) string {
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	qb423253 := qt343634.AcquireByteBuffer()
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	WriteFortunePage(qb423253, rows)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	qs435345 := string(qb423253.B)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	qt343634.ReleaseByteBuffer(qb423253)
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
	return qs435345
//line frameworks/Go/fiber/src/templates/fortune.qtpl:22
}
