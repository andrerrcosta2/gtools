// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

type Registry struct {
	engines map[string]*Engine
}

func NewRegistry() *Registry {
	return &Registry{
		engines: make(map[string]*Engine),
	}
}

func (r *Registry) Register(name string, engine *Engine) {
	r.engines[name] = engine
}

func (r *Registry) Get(name string) *Engine {
	return r.engines[name]
}

func PersistenceEngine() *Engine {
	return NewEngine(
		identitySemantics{}, // strong priority
		primitiveSemantics{},
	)
}

func DefaultEngine() *Engine {
	return NewEngine(
		identitySemantics{},
		comparableSemantics{},
		primitiveSemantics{},
		pointerSemantics{},
	)
}

func TestingEngine() *Engine {
	return NewEngine(
		comparableSemantics{},
		primitiveSemantics{},
		pointerSemantics{},
	)
}
