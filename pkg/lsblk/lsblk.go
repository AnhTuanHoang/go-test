package lsblk

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"math"
	"os/exec"
	"strconv"
	"strings"
)


func runCmd(command string) (output []byte, err error) {
	if len(command) == 0 {
		return nil, errors.New("invalid command")
	}
	commands := strings.Fields(command)
	output, err = exec.Command(commands[0], commands[1:]...).Output()
	return output, err
}


func ListDevices(driveName string) (devices map[string]Device, err error) {
	output, err := runCmd("lsblk -e7 -b -J -o name,path,fsavail,fssize,fstype,pttype,fsused,fsuse%,mountpoint,uuid,rm,hotplug,state,group,type,alignment,tran,subsystems,model " + driveName)
	if err != nil {
		return nil, err
	}

	lsblkRsp := make(map[string][]_Device)
	err = json.Unmarshal(output, &lsblkRsp)
	if err != nil {
		return nil, err
	}

	devices = make(map[string]Device)
	for _, _device := range lsblkRsp["blockdevices"] {
		var device Device
		copier.Copy(&device, &_device)

		device.Fsavail, _ = strconv.ParseUint(_device.Fsavail, 10, 64)
		device.Fsused, _ = strconv.ParseUint(_device.Fsused, 10, 64)
		device.Fssize, _ = strconv.ParseUint(_device.Fssize, 10, 64)
		if device.Fssize > 0 {
			device.Fsusage = uint(math.Round(float64(device.Fsused*100) / float64(device.Fssize)))
		}

		for i, child := range _device.Children {
			device.Children[i].Fsavail, _ = strconv.ParseUint(child.Fsavail, 10, 64)
			device.Children[i].Fsused, _ = strconv.ParseUint(child.Fsused, 10, 64)
			device.Children[i].Fssize, _ = strconv.ParseUint(child.Fssize, 10, 64)
			if device.Children[i].Fssize > 0 {
				device.Children[i].Fsusage = uint(math.Round(float64(device.Children[i].Fsused*100) / float64(device.Children[i].Fssize)))
			}
		}
		devices[device.Name] = device
	}

	return devices, nil
}

