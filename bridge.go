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
	ctx    context.Context
	cancel context.CancelFunc
	xray   *core.Instance
}

func NewXrayInstance() *XrayInstance {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	return &XrayInstance{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (x *XrayInstance) Run(data []byte) error {
	conf, err := core.LoadConfig("json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if x.xray, err = core.New(conf); err != nil {
		return err
	}

	go func() {
		if err = x.xray.Start(); err != nil {
			return
		}
		<-x.ctx.Done()
	}()
	return nil
}

func (x *XrayInstance) Stop() {
	_ = x.xray.Close()
	x.cancel()
}
