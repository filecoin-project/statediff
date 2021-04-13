package tstracker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	pgchainbs "github.com/filecoin-project/go-bs-postgres-chainnotated"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("tipset-tracker")

type TrackingChainstore interface {
	blockstore.Blockstore
	blockstore.Viewer
	io.Closer
	DeleteMany([]cid.Cid) error
	View(cid.Cid, func([]byte) error) error
	CurrentDbTipSetKey(context.Context) (*types.TipSetKey, abi.ChainEpoch, error)
	GetCurrentTipset(ctx context.Context) []cid.Cid
	StoreTipSetVist(ctx context.Context, ts *types.TipSet, isHeaChange bool) error
}

type tcs struct {
	*pgchainbs.PgBlockstore
	lastSeenTipSetMu sync.RWMutex
	lastSeenTipSet   *types.TipSet
}

func NewTrackingPgChainstoreFromEnv(ctx context.Context, readOnly bool) (TrackingChainstore, error) {
	pgbsCfg := pgchainbs.PgBlockstoreConfig{
		StoreIsWritable:         !readOnly,
		AutoUpdateSchema:        !readOnly,
		CacheInactiveBeforeRead: true,
		// DisableBlocklinkParsing: true,
	}

	envVarConn := "LOTUS_CHAINSTORE_PG_CONNSTRING"
	envVarCacheGiB := "LOTUS_CHAINSTORE_CACHE_GIB"
	envVarLogAccess := "LOTUS_CHAINSTORE_LOG_ACCESS"
	envVarNamespace := "LOTUS_CHAINSTORE_SCHEMA_NAMESPACE"

	pgbsCfg.PgxConnectString = os.Getenv(envVarConn)
	if pgbsCfg.PgxConnectString == "" {
		return nil, fmt.Errorf(
			"you must set the '%s' environment variable to a valid PostgreSQL connection string, e.g. 'postgres:///{{dbname}}?user={{user}}&password={{pass}}&host=/var/run/postgresql'",
			envVarConn,
		)
	}

	pgbsCfg.InstanceNamespace = os.Getenv(envVarNamespace)
	if pgbsCfg.InstanceNamespace == "" {
		pgbsCfg.InstanceNamespace = "main"
	}

	logDetailedAccess := os.Getenv(envVarLogAccess)
	if logDetailedAccess != "" && logDetailedAccess != "0" && logDetailedAccess != "false" {
		pgbsCfg.LogDetailedAccess = true
	}

	if customCacheSizeStr := os.Getenv(envVarCacheGiB); customCacheSizeStr != "" {
		var parseErr error
		if pgbsCfg.CacheSizeGiB, parseErr = strconv.ParseUint(customCacheSizeStr, 10, 8); parseErr != nil {
			return nil, fmt.Errorf("failed to parse requested cache size '%s': %s", customCacheSizeStr, parseErr)
		}
	}

	pgbs, err := pgchainbs.NewPgBlockstore(ctx, pgbsCfg)
	if err != nil {
		return nil, fmt.Errorf("blockstore instantiation failed: %s", err)
	}

	return &tcs{PgBlockstore: pgbs}, nil
}

func NewTrackingPgChainstore(ctx context.Context, conn string) (TrackingChainstore, error) {
	readOnly := true

	pgbsCfg := pgchainbs.PgBlockstoreConfig{
		StoreIsWritable:         !readOnly,
		AutoUpdateSchema:        !readOnly,
		CacheInactiveBeforeRead: true,
		// DisableBlocklinkParsing: true,
	}

	envVarConn := "LOTUS_CHAINSTORE_PG_CONNSTRING"
	envVarCacheGiB := "LOTUS_CHAINSTORE_CACHE_GIB"
	envVarLogAccess := "LOTUS_CHAINSTORE_LOG_ACCESS"
	envVarNamespace := "LOTUS_CHAINSTORE_SCHEMA_NAMESPACE"

	pgbsCfg.PgxConnectString = conn
	if pgbsCfg.PgxConnectString == "" {
		pgbsCfg.PgxConnectString = os.Getenv(envVarConn)
	}
	if pgbsCfg.PgxConnectString == "" {
		return nil, fmt.Errorf(
			"you must set the '%s' environment variable to a valid PostgreSQL connection string, e.g. 'postgres:///{{dbname}}?user={{user}}&password={{pass}}&host=/var/run/postgresql'",
			envVarConn,
		)
	}

	pgbsCfg.InstanceNamespace = os.Getenv(envVarNamespace)
	if pgbsCfg.InstanceNamespace == "" {
		pgbsCfg.InstanceNamespace = "main"
	}

	logDetailedAccess := os.Getenv(envVarLogAccess)
	if logDetailedAccess != "" && logDetailedAccess != "0" && logDetailedAccess != "false" {
		pgbsCfg.LogDetailedAccess = true
	}

	if customCacheSizeStr := os.Getenv(envVarCacheGiB); customCacheSizeStr != "" {
		var parseErr error
		if pgbsCfg.CacheSizeGiB, parseErr = strconv.ParseUint(customCacheSizeStr, 10, 8); parseErr != nil {
			return nil, fmt.Errorf("failed to parse requested cache size '%s': %s", customCacheSizeStr, parseErr)
		}
	}

	pgbs, err := pgchainbs.NewPgBlockstore(ctx, pgbsCfg)
	if err != nil {
		return nil, fmt.Errorf("blockstore instantiation failed: %s", err)
	}

	return &tcs{PgBlockstore: pgbs}, nil
}

func (*tcs) DeleteMany([]cid.Cid) error {
	return errors.New("not supported")
}

func (*tcs) View(cid.Cid, func([]byte) error) error {
	return errors.New("not supported")
}

func (t *tcs) Store() blockstore.Blockstore {
	return t
}
func (t *tcs) Head(ctx context.Context) cid.Cid {
	tsk, _, err := t.CurrentDbTipSetKey(ctx)
	if err != nil {
		return cid.Undef
	}
	cids := tsk.Cids()
	if len(cids) > 0 {
		return cids[0]
	}
	return cid.Undef
}
func (t *tcs) CidAtHeight(ctx context.Context, h int64) (cid.Cid, error) {
	tsd, err := t.GetFilTipSetHead(ctx)
	if err != nil {
		return cid.Undef, err
	}

	tsd, err = t.FindFilTipSet(ctx, tsd.TipSetCids, tsd.Epoch-abi.ChainEpoch(h))
	if err != nil {
		return cid.Undef, err
	}

	cids := tsd.TipSetCids
	if len(cids) > 0 {
		return cids[0], nil
	}
	return cid.Undef, nil
}

func (t *tcs) Reload() error {
	return nil
}
