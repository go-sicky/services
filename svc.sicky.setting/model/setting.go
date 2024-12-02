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
 * @package model
 * @author Dr.NP <np@herewe.tech>
 * @since 11/20/2024
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

type Setting struct {
	bun.BaseModel `bun:"table:sicky.setting"`

	ID     uuid.UUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Key    string          `bun:"key,notnull" json:"key"`
	Value  json.RawMessage `bun:"value,type:jsonb" json:"value"`
	Raw    bool            `bun:"raw" json:"raw"`
	Status int64           `bun:"status" json:"status"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (m *Setting) Set(ctx context.Context) error {
	iq := driver.DB.NewInsert().Model(m).
		On("CONFLICT (key) DO UPDATE").
		Set("value = COALESCE(EXCLUDED.value, setting.value), raw = COALESCE(EXCLUDED.raw, setting.raw), status = COALESCE(EXCLUDED.status, setting.status), updated_at = CURRENT_TIMESTAMP")

	_, err := iq.Returning("*").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "setting upsert failed", "error", err.Error())
	}

	return err
}

func (m *Setting) Get(ctx context.Context) error {
	sq := driver.DB.NewSelect().Model((*Setting)(nil))
	if m.ID != uuid.Nil {
		sq = sq.Where("id = ?", m.ID)
	}

	if m.Key != "" {
		sq = sq.Where("key = ?", m.Key)
	}

	err := sq.Scan(ctx, m)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No rows
			logger.Logger.WarnContext(ctx, "setting get no records")
			m.ID = uuid.Nil
			m.Key = ""

			return nil
		} else {
			logger.Logger.ErrorContext(ctx, "setting get failed", "error", err.Error())
		}
	}

	return err
}

func (m *Setting) Delete(ctx context.Context) error {
	dq := driver.DB.NewDelete().Model((*Setting)(nil))
	if m.ID != uuid.Nil {
		dq = dq.Where("id = ?", m.ID)
	}

	if m.Key != "" {
		dq = dq.Where("key = ?", m.Key)
	}

	ret, err := dq.Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, "setting delete failed", "error", err.Error())
	}

	a, _ := ret.RowsAffected()
	if a > 0 {
		m.Raw = true
	} else {
		m.Raw = false
	}

	return err
}

func InitSetting(ctx context.Context, dropTable bool) error {
	if dropTable {
		// Drop exist table
		_, err := driver.DB.NewDropTable().Model((*Setting)(nil)).Exec(ctx)
		if err != nil {
			logger.Logger.ErrorContext(ctx, err.Error())

			return err
		}
	}

	// Create table
	_, err := driver.DB.NewCreateTable().Model((*Setting)(nil)).Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	// Indexes
	_, err = driver.DB.NewCreateIndex().
		Model((*Setting)(nil)).
		Index("idx_sicky_setting_crated_at").
		Column("created_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	_, err = driver.DB.NewCreateIndex().
		Model((*Setting)(nil)).
		Index("idx_sicky_setting_updated_at").
		Column("updated_at").Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	// Unique
	_, err = driver.DB.NewCreateIndex().
		Model((*Setting)(nil)).
		Unique().
		Index("unq_sicky_setting_key").
		Column("key").
		Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	// Gin
	_, err = driver.DB.NewCreateIndex().
		Model((*Setting)(nil)).
		Using("gin").
		Index("gin_sicky_setting_value").
		Column("value").
		Exec(ctx)
	if err != nil {
		logger.Logger.ErrorContext(ctx, err.Error())

		return err
	}

	logger.Logger.InfoContext(ctx, "table <sicky.setting> initialized")

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
