// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//go:build nootel

package observability

import (
	"context"

	"code.superseriousbusiness.org/gotosocial/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func InitializeMetrics(ctx context.Context, db db.DB) error {
	return nil
}

func InitializeTracing(ctx context.Context) error {
	return nil
}

func MetricsMiddleware() gin.HandlerFunc {
	return nil
}

func TracingMiddleware() gin.HandlerFunc {
	return nil
}

func InstrumentBun(traces bool, metrics bool) bun.QueryHook {
	return nil
}
