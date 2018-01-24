package Legacy

import (
	"context"

	ct "github.com/autopilothq/banks/contract/types"
	"github.com/autopilothq/banks/types"
	"github.com/autopilothq/lg"
)

// SetupDatascope sets up an underlying database prior to use by the
// given instance
func SetupDatascope(
	ctx context.Context, dsID types.DatascopeID, log lg.Log, aux ct.Auxiliary,
) error {
	return nil
}
