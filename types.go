package ykpiv

const (
	DefaultReader    = "Yubikey"
	DefaultEncrypted = false
)

type YkPIV struct {
	Reader    string
	Encrypted bool

	state State
}

func New(opts ...Option) (*YkPIV, error) {
	y := &YkPIV{
		Reader:    DefaultReader,
		Encrypted: DefaultEncrypted,
	}
	for _, f := range opts {
		if err := f(y); err != nil {
			return nil, err
		}
	}
	return y, nil
}

func MustNew(opts ...Option) *YkPIV {
	y, err := New(opts...)
	if err != nil {
		panic(err)
	}
	return y
}
