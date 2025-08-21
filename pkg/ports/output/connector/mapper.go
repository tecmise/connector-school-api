package connector

type (
	Call[T any] interface {
		Find(parameter Parameter, response *T) error
		List(parameter Parameter, response *[]T) error
		Ids(parameter Parameter, response *[]int64) error
		Strings(parameter Parameter, response *[]string) error
	}
)
