package ykpiv

import (
	"fmt"
)

func (y *YkPIV) Init() error {
	if rc := cInit(&y.state, 0); !rc.isOK() {
		return fmt.Errorf("failed initializing library: %w", rc)
	}
	return nil
}

func (y *YkPIV) Done() {
	if y.state != nil {
		_ = cDone(y.state)
	}
}

func (y *YkPIV) Connect() error {
	if rc := cConnect(y.state, y.Reader, y.Encrypted); !rc.isOK() {
		return fmt.Errorf("failed to connect to yubikey: %w", rc)
	}
	return nil
}

func (y *YkPIV) InitAndConnect() error {
	if err := y.Init(); err != nil {
		return err
	}
	return y.Connect()
}

func (y *YkPIV) Run(f func() error) error {
	if err := y.InitAndConnect(); err != nil {
		return err
	}
	defer y.Done()
	return f()
}

func (y *YkPIV) MustRun(f func() error) {
	if err := y.Run(f); err != nil {
		panic(err)
	}
}
