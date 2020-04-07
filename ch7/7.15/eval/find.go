package eval

// findVariables takes an expression and a slice and returns the slice
// populated with variables that are found within the expression.
func FindVariables(e Expr, s []Var) []Var {
	switch v := e.(type) {
	case binary:
		s = FindVariables(v.x, s)
		s = FindVariables(v.y, s)
	case Var:
		s = append(s, v)
	case unary:
		s = FindVariables(v.x, s)
	case call:
		for _, arg := range v.args {
			s = FindVariables(arg, s)
		}
	case minimum:
		// Not sure why I can't do a multi case statement above.
		// i.e. case call, minimum:
		// when I try it the code won't compile saying "type Expr
		// has no field or method args"
		// As a result we have some code duplication
		for _, arg := range v.args {
			s = FindVariables(arg, s)
		}
	}
	return s
}
