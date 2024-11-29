//go:build !windows && !(arm64 && linux)

package du

func (du *DiskUsage) BlockSize() int {
	return int(du.stat.Bsize)
}
