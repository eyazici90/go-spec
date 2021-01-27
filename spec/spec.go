package spec

type Specifier interface {
	Eval() func(i interface{}) bool
}

type (
	AndSpec struct {
		ss []Specifier
	}
	OrSpec struct {
		ss []Specifier
	}
	NotSpec struct {
		s Specifier
	}
)

func (a AndSpec) Eval() func(interface{}) bool {
	return func(i interface{}) bool {
		for _, v := range a.ss {
			if !v.Eval()(i) {
				return false
			}
		}
		return true
	}
}

func (o OrSpec) Eval() func(i interface{}) bool {
	return func(i interface{}) bool {
		for _, v := range o.ss {
			if v.Eval()(i) {
				return true
			}
		}
		return false
	}
}

func (n NotSpec) Eval() func(i interface{}) bool {
	return func(i interface{}) bool {
		return !n.s.Eval()(i)
	}
}

func (a AndSpec) And(s Specifier) AndSpec {
	return And(a, s)
}

func (a AndSpec) AndNot(s Specifier) AndSpec {
	return And(a, Not(s))
}

func (a AndSpec) Or(s Specifier) OrSpec {
	return Or(a, s)
}

func (n NotSpec) And(s Specifier) AndSpec {
	return And(n, s)
}

func (n NotSpec) AndNot(s Specifier) AndSpec {
	return And(n, Not(s))
}

func (o OrSpec) And(s Specifier) AndSpec {
	return And(o, s)
}

func (o OrSpec) AndNot(s Specifier) AndSpec {
	return And(o, Not(s))
}

func (o OrSpec) Or(s Specifier) OrSpec {
	return Or(o, s)
}

func And(ss ...Specifier) AndSpec {
	return AndSpec{
		ss: ss,
	}
}

func Or(ss ...Specifier) OrSpec {
	return OrSpec{
		ss: ss,
	}
}

func Not(s Specifier) NotSpec {
	return NotSpec{
		s: s,
	}
}
