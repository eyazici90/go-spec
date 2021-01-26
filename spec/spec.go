package spec

type Spec interface {
	Eval() func(i interface{}) bool
}

type (
	AndSpec struct {
		ss []Spec
	}
	OrSpec struct {
		ss []Spec
	}
	NotSpec struct {
		s Spec
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

func (a AndSpec) And(s Spec) AndSpec {
	return And(a, s)
}

func (a AndSpec) AndNot(s Spec) AndSpec {
	return And(a, Not(s))
}

func (a AndSpec) Or(s Spec) OrSpec {
	return Or(a, s)
}

func (n NotSpec) And(s Spec) AndSpec {
	return And(n, s)
}

func (n NotSpec) AndNot(s Spec) AndSpec {
	return And(n, Not(s))
}

func (o OrSpec) And(s Spec) AndSpec {
	return And(o, s)
}

func (o OrSpec) AndNot(s Spec) AndSpec {
	return And(o, Not(s))
}

func (o OrSpec) Or(s Spec) OrSpec {
	return Or(o, s)
}

func And(ss ...Spec) AndSpec {
	return AndSpec{
		ss: ss,
	}
}

func Or(ss ...Spec) OrSpec {
	return OrSpec{
		ss: ss,
	}
}

func Not(s Spec) NotSpec {
	return NotSpec{
		s: s,
	}
}
