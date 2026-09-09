package appconstant

const (
	EqualitySign         string = "eq"
	GreaterThanSign      string = "gr"
	LessThanSign         string = "ls"
	LessEqualThanSing    string = "lse"
	GreaterEqualThanSign string = "gre"
	NotEqualitySign      string = "neq"
)

var ComparatorOperatorSigns = [6]string{
	EqualitySign,
	NotEqualitySign,
	GreaterThanSign,
	GreaterEqualThanSign,
	LessThanSign,
	LessEqualThanSing,
}
