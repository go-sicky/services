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
 * @file list.go
 * @package service
 * @author Dr.NP <np@herewe.tech>
 * @since 11/21/2024
 */

package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-sicky/services/svc.sicky.list/model"
	"github.com/go-sicky/services/svc.sicky.list/proto"
	"github.com/google/uuid"
)

type List struct{}

type RawValue struct {
	Raw string `json:"raw"`
}

func (s *List) InitDB(ctx context.Context) error {
	return model.InitList(ctx, false)
}

func (s *List) Add(ctx context.Context, item *proto.ListDef) (*proto.ListDef, error) {
	m := &model.List{
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

	err := m.Add(ctx)
	ret := &proto.ListDef{
		Id:  m.ID.String(),
		Key: m.Key,
	}

	return ret, err
}

func (s *List) Get(ctx context.Context, id uuid.UUID) (*proto.ListDef, error) {
	m := &model.List{
		ID: id,
	}

	err := m.Get(ctx)
	if m.ID == uuid.Nil {
		return nil, err
	}

	ret := &proto.ListDef{
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

func (s *List) Set(ctx context.Context, item *proto.ListDef) (bool, error) {
	if uuid.Validate(item.Id) != nil {
		return false, errors.New("invalid list item ID")
	}

	m := &model.List{
		ID:     uuid.MustParse(item.Id),
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

	c, err := m.Set(ctx)
	if c > 0 {
		return true, nil
	}

	return false, err
}

func (s *List) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	m := &model.List{
		ID: id,
	}

	c, err := m.Delete(ctx)
	if c > 0 {
		return true, nil
	}

	return false, err
}

func (s *List) List(ctx context.Context, key string, offset, limit int64) ([]*proto.ListDef, error) {
	m := &model.List{
		Key: key,
	}

	items, err := m.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	r := make([]*proto.ListDef, 0)
	for _, item := range items {
		if item.Raw {
			v := RawValue{}
			json.Unmarshal(item.Value, &v)
			item.Value = json.RawMessage(v.Raw)
		}

		r = append(r, &proto.ListDef{
			Id:     item.ID.String(),
			Key:    item.Key,
			Value:  string(item.Value),
			Raw:    item.Raw,
			Status: item.Status,
		})
	}

	return r, err
}

func (s *List) All(ctx context.Context, key string) ([]*proto.ListDef, error) {
	m := &model.List{
		Key: key,
	}

	items, err := m.All(ctx)
	if err != nil {
		return nil, err
	}

	r := make([]*proto.ListDef, 0)
	for _, item := range items {
		if item.Raw {
			v := RawValue{}
			json.Unmarshal(item.Value, &v)
			item.Value = json.RawMessage(v.Raw)
		}

		r = append(r, &proto.ListDef{
			Id:     item.ID.String(),
			Key:    item.Key,
			Value:  string(item.Value),
			Raw:    item.Raw,
			Status: item.Status,
		})
	}

	return r, err
}

func (s *List) Count(ctx context.Context, key string) (int64, error) {
	m := &model.List{
		Key: key,
	}

	return m.Count(ctx)
}

func (s *List) Purge(ctx context.Context, key string) (bool, error) {
	m := &model.List{
		Key: key,
	}

	c, err := m.Purge(ctx)
	if c > 0 {
		return true, nil
	}

	return false, err
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
