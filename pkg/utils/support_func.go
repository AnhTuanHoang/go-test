package utils

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)


func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func GetAllDrive() []PartitionStat {
	lstDisk, _ := getAllPartition()
	var listDisk []PartitionStat
	for _, val := range lstDisk {
		if strings.HasPrefix(val.Device, DRIVE_PREFIX) || val.IsCrypt || strings.HasPrefix(val.Device, DRIVE_NEWTECH_PREFIX) {
			diskType := ReadDiskType(val)
			if diskType == USB {
				continue
			}
			if diskType == SSD {
				continue
			}
			val.DiskType = diskType
			listDisk = append(listDisk, val)
		}
	}
	return listDisk
}

func ReadDiskType(device PartitionStat) uint8 {
	if device.IsCrypt {
		ChangeOwnerToRecord(device.Mountpoint)
		return HDD
	}
	//re := regexp.MustCompile("[^a-z]")
	s := strings.Split(device.Device, "/")
	//diskName := re.ReplaceAllString(s[len(s)-1], "")
	diskName := s[len(s)-1][0:3]
	rotaPath := filepath.Join("/sys", "block", diskName, "queue", "rotational")
	lines, err := readLines(rotaPath)
	if err != nil {
		return ^uint8(0)
	}
	intVal, _ := strconv.ParseInt(lines[0], 10, 8)
	if intVal == 1 && isUsb(diskName) {
		intVal = USB
	}
	if intVal == HDD {
		ChangeOwnerToRecord(device.Mountpoint)
	}
	return uint8(intVal)
}

func HostProc(combineWith ...string) string {
	return getEnv("HOST_PROC", "/proc", combineWith...)
}

func HostSys(combineWith ...string) string {
	return getEnv("HOST_SYS", "/sys", combineWith...)
}

func HostDev(combineWith ...string) string {
	return getEnv("HOST_DEV", "/dev", combineWith...)
}

func HostSecret() string {
	return getEnv("SECRET", STR_EMPTY)
}

func ChangeOwnerToRecord(mountPoint string) {
	user := os.Getenv("USER")
	cmd := exec.Command("/bin/sh", "-c", "sudo chown "+user+": "+mountPoint)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return
	}
}

func readLines(filename string) ([]string, error) {
	return readLinesOffsetN(filename, 0, -1)
}

func readLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{STR_EMPTY}, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(line) > 0 {
				ret = append(ret, strings.Trim(line, "\n"))
			}
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}

func stringsHas(target []string, src string) bool {
	for _, t := range target {
		if strings.TrimSpace(t) == src {
			return true
		}
	}
	return false
}

func removeElementInArray(list []interface{}, i int) []interface{} {
	list[i] = list[len(list)-1]
	list[len(list)-1] = STR_EMPTY
	list = list[:len(list)-1]
	return list
}

func getAllPartition() ([]PartitionStat, error) {
	root := HostProc(path.Join("1"))

	hpmPath := os.Getenv("HOST_PROC_MOUNTINFO")
	if hpmPath != STR_EMPTY {
		root = filepath.Dir(hpmPath)
	}

	lines, useMounts, filename, err := readMountFile(root)
	if err != nil {
		if hpmPath != STR_EMPTY { // don't fallback with HOST_PROC_MOUNTINFO
			return nil, err
		}
		lines, useMounts, filename, err = readMountFile(HostProc(path.Join("self")))
		if err != nil {
			return nil, err
		}
	}

	fs, err := getFileSystems()
	if err != nil {
		return nil, err
	}

	ret := make([]PartitionStat, 0, len(lines))

	for _, line := range lines {
		var d PartitionStat
		if useMounts {
			fields := strings.Fields(line)

			d = PartitionStat{
				Device:     fields[0],
				Mountpoint: unescapeFstab(fields[1]),
				Fstype:     fields[2],
				Opts:       strings.Fields(fields[3]),
				IsCrypt:    false,
			}
			if d.Device == "none" || !stringsHas(fs, d.Fstype) {
				continue
			}
		} else {
			parts := strings.Split(line, " - ")
			if len(parts) != 2 {
				return nil, fmt.Errorf("found invalid mountinfo line in file %s: %s ", filename, line)
			}

			fields := strings.Fields(parts[0])
			blockDeviceID := fields[2]
			mountPoint := fields[4]
			mountOpts := strings.Split(fields[5], ",")

			if rootDir := fields[3]; rootDir != STR_EMPTY && rootDir != "/" {
				mountOpts = append(mountOpts, "bind")
			}

			fields = strings.Fields(parts[1])
			fstype := fields[0]
			device := fields[1]

			d = PartitionStat{
				Device:     device,
				Mountpoint: unescapeFstab(mountPoint),
				Fstype:     fstype,
				Opts:       mountOpts,
				IsCrypt:    false,
			}

			if d.Device == "none" || !stringsHas(fs, d.Fstype) {
				continue
			}

			if strings.HasPrefix(d.Device, DRIVE_ENCRYPT) {
				//devpath, err := filepath.EvalSymlinks(HostDev(strings.Replace(d.Device, "/dev", "", -1)))
				//if err == nil {
				//	d.Device = devpath
				//}
				d.IsCrypt = true
			}

			if d.Device == "/dev/root" {
				devpath, err := os.Readlink(HostSys("/dev/block/" + blockDeviceID))
				if err == nil {
					d.Device = strings.Replace(d.Device, "root", filepath.Base(devpath), 1)
				}
			}
		}
		ret = append(ret, d)
	}

	return ret, nil
}

func getFileSystems() ([]string, error) {
	filename := HostProc("filesystems")
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, line := range lines {
		if !strings.HasPrefix(line, "nodev") {
			ret = append(ret, strings.TrimSpace(line))
			continue
		}
		t := strings.Split(line, "\t")
		if len(t) != 2 || t[1] != "zfs" {
			continue
		}
		ret = append(ret, strings.TrimSpace(t[1]))
	}

	return ret, nil
}

func unescapeFstab(path string) string {
	escaped, err := strconv.Unquote(`"` + path + `"`)
	if err != nil {
		return path
	}
	return escaped
}

func readMountFile(root string) (lines []string, useMounts bool, filename string, err error) {
	filename = path.Join(root, "mountinfo")
	lines, err = readLines(filename)
	if err != nil {
		var pathErr *os.PathError
		if !errors.As(err, &pathErr) {
			return
		}
		// if kernel does not support 1/mountinfo, fallback to 1/mounts
		useMounts = true
		filename = path.Join(root, "mounts")
		lines, err = readLines(filename)
		if err != nil {
			return
		}
		return
	}
	return
}

func getEnv(key string, dfault string, combineWith ...string) string {
	value := os.Getenv(key)
	if value == STR_EMPTY {
		value = dfault
	}

	switch len(combineWith) {
	case 0:
		return value
	case 1:
		return filepath.Join(value, combineWith[0])
	default:
		all := make([]string, len(combineWith)+1)
		all[0] = value
		copy(all[1:], combineWith)
		return filepath.Join(all...)
	}
}

func getKey(passphrase []byte, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key(passphrase, salt, 1000, 32, sha256.New), salt
}

func isUsb(deviceName string) bool {
	deviceLink := path.Join("/sys/block/", deviceName)
	symlink, _ := os.Readlink(deviceLink)
	return strings.Contains(symlink, "usb")
}