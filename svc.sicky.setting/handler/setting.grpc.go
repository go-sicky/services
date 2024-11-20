/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2024 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file setting.grpc.go
 * @package handler
 * @author Dr.NP <np@herewe.tech>
 * @since 11/20/2024
 */

package handler

import (
	"context"
	"errors"

	"github.com/go-sicky/services/svc.sicky.setting/proto"
	"github.com/go-sicky/services/svc.sicky.setting/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCSetting struct {
	proto.UnimplementedSettingServer

	svcSetting *service.Setting
}

func NewGRPCSetting() *GRPCSetting {
	h := &GRPCSetting{
		svcSetting: new(service.Setting),
	}

	return h
}

func (h *GRPCSetting) Name() string {
	return "grpc.setting"
}

func (h *GRPCSetting) Type() string {
	return "grpc"
}

func (h *GRPCSetting) Register(app *grpc.Server) {
	proto.RegisterSettingServer(app, h)
}

/* {{{ [Method] */
func (h *GRPCSetting) InitDB(ctx context.Context, e *emptypb.Empty) (*proto.InitDBResp, error) {
	var errs, err error
	err = h.svcSetting.InitDB(ctx)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	resp := &proto.InitDBResp{
		Result: true,
	}

	if errs != nil {
		resp.Result = false
	}

	return resp, errs
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
