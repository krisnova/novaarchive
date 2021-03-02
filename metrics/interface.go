package metrics

type Backend interface {
	Type() string
	Get(key string) (string, error)
	Set(key string, value string) error
}
