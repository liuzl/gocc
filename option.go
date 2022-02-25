package gocc

var defaultOptions = options{
	dir:    DefaultDir(),
	loader: OSLoader{},
}

type options struct {
	dir    string
	loader Loader
}

type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}
func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}
func WithDir(dir string) Option {
	return newFuncOption(func(o *options) {
		o.dir = dir
	})
}
func WithLoader(loader Loader) Option {
	return newFuncOption(func(o *options) {
		if loader == nil {
			o.loader = OSLoader{}
		} else {
			o.loader = loader
		}
	})
}
