package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

// GetDatabase retrieves a database entry by ID.
func (s *Service) GetDatabase(
	ctx context.Context, id uuid.UUID,
) (dbgen.Database, error) {
	return s.dbgen.DatabasesServiceGetDatabase(ctx, id)
}
