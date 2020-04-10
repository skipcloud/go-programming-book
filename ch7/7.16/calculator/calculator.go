package calculator

import (
	"fmt"
	"html/template"

	"github.com/skipcloud/go-programming-book/ch7/7.16/eval"
)

type Calculator struct {
	input  string
	result float64
	tmpl   string
}

const outputTemplate = `
<html>
<body>
<style>
table {
	font-size: 48px;
}
td {
	text-align: center;
	height: 70px;
	width: 70px;
}
th {
	text-align: left;
}
a {
	text-decoration: none;
	color: black;
}
</style>
<table>
	<tr>
		<th colspan="4">{{ data}}</th>
	</tr>
	<tr>
		<td><a href="/reset">c</a></td>
		<td><a href="/input?data=1">1</a></td>
		<td><a href="/input?data=2">2</a></td>
		<td><a href="/input?data=3">3</a></td>
	</tr>
		<td><a href="/input?data=%2B">+</a></td>
		<td><a href="/input?data=4">4</a></td>
		<td><a href="/input?data=5">5</a></td>
		<td><a href="/input?data=6">6</a></td>
	<tr>
		<td><a href="/input?data=-">-</a></td>
		<td><a href="/input?data=7">7</a></td>
		<td><a href="/input?data=8">8</a></td>
		<td><a href="/input?data=9">9</a></td>
	<tr>
		<td><a href="/input?data=*">*</a></td>
		<td><a href="/input?data=%2F">/</a></td>
		<td><a href="/input?data=0">0</a></td>
		<td><a href="/">=</a></td>
	</tr>
</table>
</body>
</html>
`

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) NewTemplate() *template.Template {
	return template.Must(
		template.New("output").
			Funcs(template.FuncMap{"data": c.data}).
			Parse(outputTemplate),
	)
}
func (c *Calculator) Reset() {
	c.input = ""
	c.result = 0
}

func (c *Calculator) Input(s string) {
	c.input += s
}

func (c *Calculator) Calculate() error {
	if c.input == "" {
		return nil
	}
	e, err := eval.Parse(c.input)
	if err != nil {
		return err
	}
	c.input = ""
	c.result = e.Eval(eval.Env{})

	return nil
}

func (c *Calculator) data() string {
	if c.input != "" {
		return c.input
	}
	return fmt.Sprintf("%g", c.result)
}
