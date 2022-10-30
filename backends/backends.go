package backends

import "context"

type Backend interface {
	WriteMetadata(ctx context.Context, hash string, data map[string]string) (map[string]string, error)
	ReadMetadata(ctx context.Context, hash string, keys []string) (map[string]string, error)

	StoreArtifacts(ctx context.Context, hash string, paths []string) ([]string, error)
}
