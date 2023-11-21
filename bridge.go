package xraybridge

import (
	"bytes"
	"context"

	_ "github.com/xtls/xray-core/app/proxyman/inbound"
	_ "github.com/xtls/xray-core/app/proxyman/outbound"
	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/main/confloader/external"
	_ "github.com/xtls/xray-core/main/json"
)

type XrayInstance struct {
	ctx context.Context
}

func NewXrayInstance() *XrayInstance {
	return &XrayInstance{
		ctx: context.Background(),
	}
}

func (x *XrayInstance) Run(data []byte) error {
	conf, err := core.LoadConfig("json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	xray, err := core.New(conf)
	if err != nil {
		return err
	}

	go func() {

	}()
	if err = xray.Start(); err != nil {
		return err
	}
	<-x.ctx.Done()
	return nil
}

func (x *XrayInstance) Stop() {
	x.ctx.Done()
}
