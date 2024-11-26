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
 * @file setting.go
 * @package service
 * @author Dr.NP <np@herewe.tech>
 * @since 11/20/2024
 */

package service

import (
	"context"
	"encoding/json"

	"github.com/go-sicky/services/svc.sicky.setting/model"
	"github.com/go-sicky/services/svc.sicky.setting/proto"
	"github.com/google/uuid"
)

type Setting struct{}

type RawValue struct {
	Raw string `json:"raw"`
}

func (s *Setting) InitDB(ctx context.Context) error {
	return model.InitSetting(ctx, false)
}

func (s *Setting) Set(ctx context.Context, item *proto.SettingDef) (*proto.SettingDef, error) {
	m := &model.Setting{
		Key:    item.Key,
		Raw:    item.Raw,
		Status: item.Status,
	}

	if !item.Raw {
		if json.Valid([]byte(item.Value)) {
			m.Value = json.RawMessage(item.Value)
		}
	} else {
		v := RawValue{
			Raw: item.Value,
		}
		m.Value, _ = json.Marshal(v)
	}

	err := m.Set(ctx)
	ret := &proto.SettingDef{
		Id:  m.ID.String(),
		Key: m.Key,
	}

	return ret, err
}

func (s *Setting) GetByID(ctx context.Context, id uuid.UUID) (*proto.SettingDef, error) {
	m := &model.Setting{
		ID: id,
	}

	err := m.Get(ctx)
	if m.ID == uuid.Nil {
		return nil, err
	}

	ret := &proto.SettingDef{
		Id:     m.ID.String(),
		Key:    m.Key,
		Raw:    m.Raw,
		Status: m.Status,
	}
	if m.Raw {
		v := RawValue{}
		json.Unmarshal(m.Value, &v)
		ret.Value = v.Raw
	} else {
		ret.Value = string(m.Value)
	}

	return ret, err
}

func (s *Setting) GetByKey(ctx context.Context, key string) (*proto.SettingDef, error) {
	m := &model.Setting{
		Key: key,
	}

	err := m.Get(ctx)
	if m.Key == "" {
		return nil, err
	}

	ret := &proto.SettingDef{
		Id:     m.ID.String(),
		Key:    m.Key,
		Raw:    m.Raw,
		Status: m.Status,
	}
	if m.Raw {
		v := RawValue{}
		json.Unmarshal(m.Value, &v)
		ret.Value = v.Raw
	} else {
		ret.Value = string(m.Value)
	}

	return ret, err
}

func (s *Setting) DeleteByID(ctx context.Context, id uuid.UUID) (bool, error) {
	m := &model.Setting{
		ID: id,
	}

	err := m.Delete(ctx)

	return m.Raw, err
}

func (s *Setting) DeleteByKey(ctx context.Context, key string) (bool, error) {
	m := &model.Setting{
		Key: key,
	}

	err := m.Delete(ctx)

	return m.Raw, err
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
