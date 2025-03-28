package compose

import (
	"context"

	"github.com/octohelm/storage/deprecated/pkg/dal"
	"github.com/octohelm/storage/deprecated/pkg/dal/compose/querierpatcher"
	"github.com/octohelm/storage/pkg/sqlbuilder"
)

func DeleteWhereIfFound[M sqlbuilder.Model, T any](ctx context.Context, idCol sqlbuilder.TypedColumn[T], patchers ...querierpatcher.Typed[M]) error {
	return dal.Delete[M]().
		Where(idCol.V(
			querierpatcher.InSelectIfExists(ctx, idCol, patchers...),
		)).
		Save(ctx)
}
