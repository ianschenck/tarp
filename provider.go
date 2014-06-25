package tarp

type Provider interface {
	Init(env map[string]string) error
	Open(dsn string) (Package, error)
	Store(dsn string, p Package) error
}
