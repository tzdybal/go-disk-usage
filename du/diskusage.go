// +build !windows

package du

import "golang.org/x/sys/unix"

// DiskUsage contains usage data and provides user-friendly access methods
type DiskUsage struct {
	stat *unix.Statfs_t
}

// NewDiskUsage returns an object holding the disk usage of volumePath
// or nil in case of error (invalid path, etc)
func NewDiskUsage(volumePath string) *DiskUsage {
	var stat unix.Statfs_t
	if unix.Statfs(volumePath, &stat) == nil {
		return &DiskUsage{&stat}
	}
	return nil
}

// Free returns total free bytes on file system
func (du *DiskUsage) Free() uint64 {
	return du.stat.Bfree * uint64(du.stat.Bsize)
}

// Available returns total available bytes on file system to an unpriveleged user
func (du *DiskUsage) Available() uint64 {
	return uint64(du.stat.Bavail) * uint64(du.stat.Bsize)
}

// Size returns total size of the file system
func (du *DiskUsage) Size() uint64 {
	return du.stat.Blocks * uint64(du.stat.Bsize)
}

// Used returns total bytes used in file system
func (du *DiskUsage) Used() uint64 {
	return du.Size() - du.Free()
}

// Usage returns percentage of use on the file system
func (du *DiskUsage) Usage() float32 {
	return float32(du.Used()) / float32(du.Size())
}
