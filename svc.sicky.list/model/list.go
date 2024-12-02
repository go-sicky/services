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
 * @package model
 * @author Dr.NP <np@herewe.tech>
 * @since 11/21/2024
 */

package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-sicky/sicky/driver"
	"github.com/go-sicky/sicky/logger"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type List struct {
	bun.BaseModel `bun:"table:sicky.list"`

	ID     uuid.UUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Key    string          `bun:"key,notnull" json:"key"`
	Value  json.RawMessage `bun:"value,type:jsonb" json:"value"`
	Raw    bool            `bun:"raw" json:"raw"`
	Status int64           `bun:"status" json:"status"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero" json:"-"`
}

func (m *List) Add(ctx context.Context) error {
	m.ID = uuid.Nil
	iq := driver.DB.NewInsert().Model(m)
	_, err := iq.Returning("*").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list insert failed", "error", err.Error())
	}

	return err
}

func (m *List) Get(ctx context.Context) error {
	sq := driver.DB.NewSelect().Model((*List)(nil))
	if m.ID != uuid.Nil {
		sq = sq.Where("id = ?", m.ID)
	}

	if m.Key != "" {
		sq = sq.Where("key = ?", m.Key)
	}

	err := sq.Limit(1).Scan(ctx, m)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No rows
			logger.Logger.WarnContext(ctx, "list get no records")
			m.ID = uuid.Nil
			m.Key = ""

			return nil
		} else {
			logger.Logger.ErrorContext(ctx, "list get failed", "error", err.Error())
		}
	}

	return err
}

func (m *List) Set(ctx context.Context) (int64, error) {
	uq := driver.DB.NewUpdate().Model((*List)(nil))
	if m.ID != uuid.Nil {
		uq = uq.Where("id = ?", m.ID)
	}

	if m.Key != "" {
		uq = uq.Where("key = ?", m.Key)
	}

	if m.Value != nil {
		uq = uq.Set("value = ?", m.Value).Set("raw = ?", m.Raw)
	}

	if m.Status != 0 {
		uq = uq.Set("status = ?", m.Status)
	}

	r, err := uq.Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list set failed", "error", err.Error())

		return 0, err
	}

	ra, err := r.RowsAffected()
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list set RowsAffected() failed", "error", err.Error())

		return 0, err
	}

	return ra, nil
}

func (m *List) Delete(ctx context.Context) (int64, error) {
	dq := driver.DB.NewDelete().Model((*List)(nil))
	if m.ID != uuid.Nil {
		dq = dq.Where("id = ?", m.ID)
	}

	if m.Key != "" {
		dq = dq.Where("key = ?", m.Key)
	}

	r, err := dq.Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list delete failed", "error", err.Error())

		return 0, err
	}

	ra, err := r.RowsAffected()
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list delete RowsAffected() failed", "error", err.Error())

		return 0, err
	}

	return ra, nil
}

func (m *List) List(ctx context.Context, offset, limit int64) ([]*List, error) {
	var (
		items []*List
	)

	sq := driver.DB.NewSelect().Model((*List)(nil)).Where("key = ?", m.Key)
	err := sq.Offset(int(offset)).Limit(int(limit)).Scan(ctx, &items)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list list failed", "error", err.Error())

		return nil, err
	}

	return items, nil
}

func (m *List) All(ctx context.Context) ([]*List, error) {
	var (
		items []*List
	)

	sq := driver.DB.NewSelect().Model((*List)(nil)).Where("key = ?", m.Key)
	err := sq.Scan(ctx, &items)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list all failed", "error", err.Error())

		return nil, err
	}

	return items, nil
}

func (m *List) Count(ctx context.Context) (int64, error) {
	sq := driver.DB.NewSelect().Model((*List)(nil)).Where("key = ?", m.Key)
	c, err := sq.Count(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list count failed", "error", err.Error())

		return 0, err
	}

	return int64(c), nil
}

func (m *List) Purge(ctx context.Context) (int64, error) {
	dq := driver.DB.NewDelete().Model((*List)(nil)).Where("key = ?", m.Key)
	r, err := dq.Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list purge failed", err.Error())

		return 0, err
	}

	c, err := r.RowsAffected()
	if err != nil {
		logger.Logger.ErrorContext(ctx, "list purge RowsAffected() failed", err.Error())

		return 0, err
	}

	return c, nil
}

func InitList(ctx context.Context, dropTable bool) error {
	if dropTable {
		// Drop exist table
		_, err := driver.DB.NewDropTable().Model((*List)(nil)).Exec(ctx)
		if err != nil {
			logger.Logger.ErrorContext(ctx, err.Error())

			return err
		}
	}

	// Create table
	_, err := driver.DB.NewCreateTable().Model((*List)(nil)).Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	// Indexes
	_, err = driver.DB.NewCreateIndex().
		Model((*List)(nil)).
		Index("idx_list_crated_at").
		Column("created_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	_, err = driver.DB.NewCreateIndex().
		Model((*List)(nil)).
		Index("idx_list_updated_at").
		Column("updated_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	_, err = driver.DB.NewCreateIndex().
		Model((*List)(nil)).
		Index("idx_list_deleted_at").
		Column("deleted_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	_, err = driver.DB.NewCreateIndex().
		Model((*List)(nil)).
		Index("idx_list_key").
		Column("key").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	logger.Logger.InfoContext(ctx, "table <sicky.list> initialized")

	return err
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
