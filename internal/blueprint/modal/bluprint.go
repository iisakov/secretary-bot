package modal

type Bluprint interface {
	Use() Bluprint
}

type Content map[string]string
