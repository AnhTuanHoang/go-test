package smart

import "bytes"

type Uint128 struct {
	Val [2]uint64
}
type NvmeIdentController struct {
	VendorID        uint16
	Ssvid           uint16
	SerialNumberRaw [20]byte
	ModelNumberRaw  [40]byte
	FirmwareRevRaw  [8]byte
	Rab             uint8
	IEEE            [3]byte
	Cmic            uint8
	Mdts            uint8
	Cntlid          uint16
	Ver             uint32
	Rtd3r           uint32
	Rtd3e           uint32
	Oaes            uint32
	Ctratt          uint32
	Rrls            uint16
	_               [9]byte
	CntrlType       uint8
	Fguid           [16]byte
	Crdt1           uint16
	Crdt2           uint16
	Crdt3           uint16
	_               [119]byte // ...
	Nvmsr           uint8
	Vwci            uint8
	Mec             uint8
	Oacs            uint16
	Acl             uint8
	Aerl            uint8
	Frmw            uint8
	Lpa             uint8
	Elpe            uint8
	Npss            uint8
	Avscc           uint8
	Apsta           uint8
	Wctemp          uint16
	Cctemp          uint16
	Mtfa            uint16
	Hmpre           uint32
	Hmmin           uint32
	Tnvmcap         Uint128
	Unvmcap         Uint128
	Rpmbs           uint32
	Edstt           uint16
	Dsto            uint8
	Fwug            uint8
	Kas             uint16
	Hctma           uint16
	Mntmt           uint16
	Mxtmt           uint16
	Sanicap         uint32
	_               [180]byte               // ...
	Sqes            uint8                   // Submission Queue Entry Size
	Cqes            uint8                   // Completion Queue Entry Size
	_               [2]byte                 // (defined in NVMe 1.3 spec)
	Nn              uint32                  // Number of Namespaces
	Oncs            uint16                  // Optional NVM Command Support
	Fuses           uint16                  // Fused Operation Support
	Fna             uint8                   // Format NVM Attributes
	Vwc             uint8                   // Volatile Write Cache
	Awun            uint16                  // Atomic Write Unit Normal
	Awupf           uint16                  // Atomic Write Unit Power Fail
	Nvscc           uint8                   // NVM Vendor Specific Command Configuration
	_               uint8                   // ...
	Acwu            uint16                  // Atomic Compare & Write Unit
	_               [2]byte                 // ...
	Sgls            uint32                  // SGL Support
	_               [1508]byte              // ...
	Psd             [32]NvmeIdentPowerState // Power State Descriptors
	Vs              [1024]byte              // Vendor Specific
} // 4096 bytes

func (c *NvmeIdentController) ModelNumber() string {
	return string(bytes.TrimSpace(c.ModelNumberRaw[:]))
}

func (c *NvmeIdentController) SerialNumber() string {
	return string(bytes.TrimSpace(c.SerialNumberRaw[:]))
}

func (c *NvmeIdentController) FirmwareRev() string {
	return string(bytes.TrimSpace(c.FirmwareRevRaw[:]))
}

type NvmeIdentPowerState struct {
	MaxPower        uint16 // Maximum Power (specified in MaxPowerScale units)
	_               uint8
	Flags           uint8  // bit 0 - MaxPowerScale, bit 1 - Non-Operational State
	EntryLat        uint32 // Entry Latency
	ExitLat         uint32 // Exit Latency
	ReadThroughput  uint8
	ReadLatency     uint8
	WriteThroughput uint8
	WriteLatency    uint8
	IdlePower       uint16
	IdleScale       uint8
	_               uint8
	ActivePower     uint16
	ActiveWorkScale uint8 // Active Power Workload + Active Power Scale
	_               [9]byte
}

type NvmeLBAF struct {
	Ms uint16 // Metadata Size
	// LBA Data Size (LBADS): This field indicates the LBA data size supported. The value is reported
	// in terms of a power of two (2^n). A value smaller than 9 (i.e., 512 bytes) is not supported. If the
	// value reported is 0h, then the LBA format is not supported / used or is not currently available
	Ds uint8 // LBA Data Size
	Rp uint8 // Relative Performance
}

type NvmeIdentNamespace struct {
	Nsze     uint64  // Namespace Size
	Ncap     uint64  // Namespace Capacity
	Nuse     uint64  // Namespace Utilization
	Nsfeat   uint8   // Namespace Features
	Nlbaf    uint8   // Number of LBA Formats
	Flbas    uint8   // Formatted LBA Size
	Mc       uint8   // Metadata Capabilities
	Dpc      uint8   // End-to-end Data Protection Capabilities
	Dps      uint8   // End-to-end Data Protection Type Settings
	Nmic     uint8   // Namespace Multi-path I/O and Namespace Sharing Capabilities
	Rescap   uint8   // Reservation Capabilities
	Fpi      uint8   // Format Progress Indicator
	Dlfeat   uint8   // Deallocate Logical Block Features
	Nawun    uint16  // Namespace Atomic Write Unit Normal
	Nawupf   uint16  // Namespace Atomic Write Unit Power Fail
	Nacwu    uint16  // Namespace Atomic Compare & Write Unit
	Nabsn    uint16  // Namespace Atomic Boundary Size Normal
	Nabo     uint16  // Namespace Atomic Boundary Offset
	Nabspf   uint16  // Namespace Atomic Boundary Size Power Fail
	Noiob    uint16  // Namespace Optimal I/O Boundary
	Nvmcap   Uint128 // NVM Capacity
	Npwg     uint16  // Namespace Preferred Write Granularity
	Npwa     uint16  // Namespace Preferred Write Alignment
	Npdg     uint16  // Namespace Preferred Deallocate Granularity
	Npda     uint16  // Namespace Preferred Deallocate Alignment
	Nows     uint16  // Namespace Optimal Write Size
	Mssrl    uint16  // Maximum Single Source Range Length
	Mcl      uint32  // Maximum Copy Length
	Msrc     uint8   // Maximum Source Range Count
	_        [11]byte
	Anagrpid uint32 // ANA Group Identifier
	_        [3]byte
	Nsattr   uint8        // Namespace Attributes
	Nvmseid  uint16       // NVM Set Identifier
	Endgid   uint16       // Endurance Group Identifier
	Nguid    [16]byte     // Namespace Globally Unique Identifier
	Eui64    [8]byte      // IEEE Extended Unique Identifier
	Lbaf     [64]NvmeLBAF // LBA Format Support
	Vs       [3712]byte
} // 4096 bytes

func (ns *NvmeIdentNamespace) LbaSize() uint64 {
	return uint64(1) << ns.Lbaf[ns.Flbas&0xf].Ds
}

type NvmeSMARTLog struct {
	CritWarning            uint8  // Critical Warning
	Temperature            uint16 // Composite Temperature
	AvailSpare             uint8  // Available Spare
	SpareThresh            uint8  // Available Spare Threshold
	PercentUsed            uint8  // Percentage Used
	EnduranceCritWarning   uint8  // Endurance Group Critical Warning Summary
	_                      [25]byte
	DataUnitsRead          Uint128   // Data Units Read
	DataUnitsWritten       Uint128   // Data Units Written
	HostReads              Uint128   // Host Read Commands
	HostWrites             Uint128   // Host Write Commands
	CtrlBusyTime           Uint128   // Controller Busy Time
	PowerCycles            Uint128   // Power Cycles
	PowerOnHours           Uint128   // Power On Hours
	UnsafeShutdowns        Uint128   // Unsafe Shutdowns
	MediaErrors            Uint128   // Media and Data Integrity Errors
	NumErrLogEntries       Uint128   // Number of Error Information Log Entries
	WarningTempTime        uint32    // Warning Composite Temperature Time
	CritCompTime           uint32    // Critical Composite Temperature Time
	TempSensor             [8]uint16 // Temperature Sensors
	ThermalTransitionCount [2]uint32 // Thermal Management Transition Count
	ThermalManagementTime  [2]uint32 // Total Time For Thermal Management
	_                      [280]byte
} // 512 bytes

const (
	nvmeAdminDeleteSq      = 0x00
	nvmeAdminCreateSq      = 0x01
	nvmeAdminGetLogPage    = 0x02
	nvmeAdminDeleteCq      = 0x04
	nvmeAdminCreateCq      = 0x05
	nvmeAdminIdentify      = 0x06
	nvmeAdminAbortCmd      = 0x08
	nvmeAdminSetFeatures   = 0x09
	nvmeAdminGetFeatures   = 0x0a
	nvmeAdminAsyncEvent    = 0x0c
	nvmeAdminNsMgmt        = 0x0d
	nvmeAdminActivateFw    = 0x10
	nvmeAdminDownloadFw    = 0x11
	nvmeAdminDevSelfTest   = 0x14
	nvmeAdminNsAttach      = 0x15
	nvmeAdminKeepAlive     = 0x18
	nvmeAdminDirectiveSend = 0x19
	nvmeAdminDirectiveRecv = 0x1a
	nvmeAdminVirtualMgmt   = 0x1c
	nvmeAdminNvmeMiSend    = 0x1d
	nvmeAdminNvmeMiRecv    = 0x1e
	nvmeAdminDbbuf         = 0x7C
	nvmeAdminFormatNvm     = 0x80
	nvmeAdminSecuritySend  = 0x81
	nvmeAdminSecurityRecv  = 0x82
	nvmeAdminSanitizeNvm   = 0x84
	nvmeAdminGetLbaStatus  = 0x86
	nvmeAdminVendorStart   = 0xC0
)

const (
	nvmeLogSupportedPages    = 0x0
	nvmeLogErrorInformation  = 0x1
	nvmeLogSmartInformation  = 0x2
	nvmeLogFirmwareInfo      = 0x3
	nvmeLogChangedNamespace  = 0x4
	nvmeLogCommandsSupported = 0x5
	nvmeLogDeviceSelftest    = 0x6
)

func (d *NVMeDevice) Type() string {
	return "nvme"
}

func (d *NVMeDevice) ReadGenericAttributes() (*GenericAttributes, error) {
	log, err := d.ReadSMART()
	if err != nil {
		return nil, err
	}

	a := GenericAttributes{}
	a.Temperature = uint64(log.Temperature - 273) // NVMe reports the temperature in Kelvins, normalize it to Celsius
	a.Read = log.DataUnitsRead.Val[0]
	a.Written = log.DataUnitsWritten.Val[0]
	a.PowerOnHours = log.PowerOnHours.Val[0]
	a.PowerCycles = log.PowerCycles.Val[0]
	return &a, nil
}

func (d *NVMeDevice) Identify() (*NvmeIdentController, []NvmeIdentNamespace, error) {
	controller, err := d.readControllerIdentifyData()
	if err != nil {
		return nil, nil, err
	}

	var ns []NvmeIdentNamespace
	// QEMU has 256 namespaces for some reason, TODO: clarify
	for i := 0; i < int(controller.Nn); i++ {
		namespace, err := d.readNamespaceIdentifyData(i + 1)
		if err != nil {
			return nil, nil, err
		}
		if namespace.Nsze == 0 {
			continue
		}

		ns = append(ns, *namespace)
	}

	return controller, ns, nil
}
