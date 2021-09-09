package iface

import (
	"context"

	blocklist "github.com/cloudflare/go-ipfs-blocklist"
	cid "github.com/ipfs/go-cid"
)

type ResolvedContent struct {
	Cid   cid.Cid  `json:"cid"`
	Links []string `json:"links"`
}

// SafemodeAPI specifies the interface to Safemode
type SafemodeAPI interface {
	Block(ctx context.Context, data blocklist.BlockData) ([]ResolvedContent, error)
	Unblock(ctx context.Context, data blocklist.BlockData) ([]cid.Cid, error)
	Search(ctx context.Context, content string) (*blocklist.BlocklistItem, error)
	Purge(ctx context.Context, content string) (cid.Cid, error)
	GetLogs(ctx context.Context, limit int) ([]*blocklist.Action, error)
	AddLog(ctx context.Context, act *blocklist.Action) error
	Contains(ctx context.Context, id cid.Cid) (bool, error)
	ResolveContent(ctx context.Context, content string) (*ResolvedContent, error)
}
