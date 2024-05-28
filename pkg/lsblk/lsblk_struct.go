package lsblk

import "reflect"

type Device struct {
	Name       string   `json:"name"`
	Path       string   `json:"path"`
	Fsavail    uint64   `json:"fsavail"`
	Fssize     uint64   `json:"fssize"`
	Fsused     uint64   `json:"fsused"`
	Fsusage    uint     `json:"fsusage"` // percent that was used
	Fstype     string   `json:"fstype"`
	Pttype     string   `json:"pttype"`
	Mountpoint string   `json:"mountpoint"`
	UUID       string   `json:"uuid"`
	Rm         bool     `json:"rm"`
	Hotplug    bool     `json:"hotplug"`
	State      string   `json:"state"`
	Group      string   `json:"group"`
	Type       string   `json:"type"`
	Alignment  int      `json:"alignment"`
	Tran       string   `json:"tran"`
	Subsystems string   `json:"subsystems"`
	Model      string   `json:"model"`
	Children   []Device `json:"children"`
}

type _Device struct {
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Fsavail    string    `json:"fsavail"`
	Fssize     string    `json:"fssize"`
	Fstype     string    `json:"fstype"`
	Pttype     string    `json:"pttype"`
	Fsused     string    `json:"fsused"`
	Fsuse      string    `json:"fsuse%"`
	Mountpoint string    `json:"mountpoint"`
	UUID       string    `json:"uuid"`
	Rm         bool      `json:"rm"`
	Hotplug    bool      `json:"hotplug"`
	State      string    `json:"state"`
	Group      string    `json:"group"`
	Type       string    `json:"type"`
	Alignment  int       `json:"alignment"`
	Tran       string    `json:"tran"`
	Subsystems string    `json:"subsystems"`
	Model      string    `json:"model"`
	Children   []_Device `json:"children"`
}

type TypeConverter struct {
	SrcType interface{}
	DstType interface{}
	Fn      func(src interface{}) (dst interface{}, err error)
}

type converterPair struct {
	SrcType reflect.Type
	DstType reflect.Type
}

// Tag Flags
type flags struct {
	BitFlags  map[string]uint8
	SrcNames  tagNameMapping
	DestNames tagNameMapping
}

// Field Tag name mapping
type tagNameMapping struct {
	FieldNameToTag map[string]string
	TagToFieldName map[string]string
}

type Option struct {
	IgnoreEmpty bool
	DeepCopy    bool
	Converters  []TypeConverter
}