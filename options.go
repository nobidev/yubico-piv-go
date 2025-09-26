package ykpiv

type Option func(*YkPIV) error

func WithEncrypted(encrypted bool) Option {
	return func(y *YkPIV) error {
		y.Encrypted = encrypted
		return nil
	}
}
