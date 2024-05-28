package utils

import "syscall"

type DiskStatus struct {
	stat *syscall.Statfs_t
}

type PartitionStat struct {
	Device     string   `json:"device"`
	Mountpoint string   `json:"mountpoint"`
	Fstype     string   `json:"fstype"`
	Opts       []string `json:"opts"`
	IsCrypt    bool     `json:"IsCrypt"`
	DiskType   uint8    `json:"DiskType"`
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (du *DiskStatus) Free() uint64 {
	return du.stat.Bfree * uint64(du.stat.Bsize)
}

func (du *DiskStatus) Available() uint64 {
	return du.stat.Bavail * uint64(du.stat.Bsize)
}

func (du *DiskStatus) Size() uint64 {
	return uint64(du.stat.Blocks) * uint64(du.stat.Bsize)
}

func (du *DiskStatus) Used() uint64 {
	return du.Size() - du.Free()
}

func (du *DiskStatus) Usage() uint8 {
	return uint8((du.Used() * 100) / du.Size())
}
