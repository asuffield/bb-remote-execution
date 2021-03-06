package cas

import (
	"context"

	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/buildbarn/bb-storage/pkg/filesystem"
)

// FileFetcher is responsible for fetching files from the Content
// Addressable Storage (CAS), storing its contents inside a file on
// disk.
type FileFetcher interface {
	GetFile(ctx context.Context, digest digest.Digest, directory filesystem.Directory, name string, isExecutable bool) error
}
