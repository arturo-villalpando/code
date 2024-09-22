package classes

/**
 * Polymorphism
 */
type ScreamInterface interface {
	Scream() string
}

func Scream(g ScreamInterface) string {
	return g.Scream()
}
