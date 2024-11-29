//go:build arm64 && linux

package du

func (du *DiskUsage) BlockSize() int {
	return int(du.stat.Frsize)
}
