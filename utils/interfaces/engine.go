package interfaces

type IEngine interface {
	LoadPlugins(paths []string) error
	// LoadDb()
	Init() error
	StartServer()
	// LoadPlugins(paths []string)
}
