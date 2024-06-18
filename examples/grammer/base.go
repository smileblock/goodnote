package grammer

type Base interface {
	Field() string
}

type Sub1 struct {
	field string
}

func (s *Sub1) Field() string {
	return s.field
}

type Sub2 struct {
	field string
}

func (s *Sub2) Field() string {
	return s.field
}

func Demo(b Base) {
	b.Field()
}

func Demo2() {
	Demo(&Sub1{})
}
