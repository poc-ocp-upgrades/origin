package generator

type Generator interface {
	GenerateValue(expression string) (interface{}, error)
}
