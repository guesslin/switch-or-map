package plugs

const (
	FactoryName = "Factory"
)

type Getter interface {
	Get(string) (string, bool)
}

type Factory func() Getter
