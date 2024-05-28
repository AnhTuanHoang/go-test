package main

import (
	"encoding/json"
	"fmt"
	"os"
	"test-func/pkg/onvif-20"
	"test-func/pkg/onvif-20/device"
	"test-func/pkg/onvif-20/media"
	"test-func/pkg/onvif-20/media2"
	"test-func/pkg/onvif-20/ptz"
	"test-func/pkg/onvif-20/xsd/onvif"
)

const ErrorPrefix = "Error:"

func main() {
	argsWithoutProg := os.Args[1:]
	ip := argsWithoutProg[0]
	port := argsWithoutProg[1]
	user := argsWithoutProg[2]
	password := argsWithoutProg[3]
	function := argsWithoutProg[4]
	var param string
	if len(argsWithoutProg) >= 6 {
		param = argsWithoutProg[5]
	}
	url := "%s:%s"
	xAddr := fmt.Sprintf(url, ip, port)
	dev, err := onvif20.NewDevice(onvif20.DeviceParams{
		Xaddr:    xAddr,
		Username: user,
		Password: password,
	})
	if err != nil {
		fmt.Printf("%s Create Device error", ErrorPrefix)
	}
	switch function {
	case onvif20.GetProfiles:
		getProfiles(dev)
	case onvif20.GetStreamUri:
		getStreamUri(dev)
	case onvif20.GotoPreset:
		gotoPreset(dev, param)
	case onvif20.GetNetworkInterfaces:
		getNetwork(dev)
	default:
		fmt.Printf("%s Function not support yet", ErrorPrefix)
	}
}

func getNetwork(dev *onvif20.Device) {
	input, _ := json.Marshal(device.GetNetworkInterfaces{})
	c, err := dev.CallOnvifFunction(onvif20.DeviceWebService, onvif20.GetNetworkInterfaces, input)
	fmt.Println(err)
	fmt.Println(c)
}

func getListProfile(dev *onvif20.Device) Profile {
	var c string
	input, _ := json.Marshal(media.GetProfiles{})
	c, err := dev.CallOnvifFunction(onvif20.MediaWebService, onvif20.GetProfiles, input)
	if err != nil {
		fmt.Println("onvif 10 error:", err.Error())
		input, _ = json.Marshal(media2.GetProfiles{})
		c, err = dev.CallOnvifFunction(onvif20.Media2WebService, onvif20.GetProfiles, input)
		if err != nil {
			fmt.Println("onvif 20 error:", err.Error())
		}
	}
	profiles := Profile{}
	_ = json.Unmarshal([]byte(c), &profiles)
	return profiles
}

func getProfiles(dev *onvif20.Device) {
	profiles := getListProfile(dev)
	for _, profi := range profiles.Profiles{
		fmt.Println(profi.Token)
	}
}

func getStreamUri(dev *onvif20.Device) {
	profiles := getListProfile(dev)
	for _, tokenInfo := range profiles.Profiles {
		token := tokenInfo.Token
		tokenReference := onvif.ReferenceToken(token)
		stream := onvif.StreamType("RTP-Unicast")
		protocol := onvif.TransportProtocol("RTSP")
		transport := onvif.Transport{
			Protocol: &protocol,
			Tunnel:   nil,
		}
		streamSetup := onvif.StreamSetup{
			Stream:    &stream,
			Transport: &transport,
		}
		getStreamUriConfig := media.GetStreamUri{
			ProfileToken: &tokenReference,
			StreamSetup:  &streamSetup,
		}
		input, _ := json.Marshal(getStreamUriConfig)
		rs, _ := dev.CallOnvifFunction(onvif20.MediaWebService, onvif20.GetStreamUri, input)
		result := StreamUri{}
		_ = json.Unmarshal([]byte(rs), &result)
		output := StreamUriResult{
			Uri:     result.MediaURI.URI,
			Profile: token,
		}
		b,_ := json.Marshal(output)
		fmt.Println(string(b))
	}
}

func gotoPreset(dev *onvif20.Device, ptzToken string) {
	profiles := getListProfile(dev)
	if len(profiles.Profiles) > 0 {
		firstToken := profiles.Profiles[0].Token
		profileToken := onvif.ReferenceToken(firstToken)
		presetToken := onvif.ReferenceToken(ptzToken)
		presetData := ptz.GotoPreset{
			XMLName:      "",
			ProfileToken: &profileToken,
			PresetToken:  &presetToken,
			Speed:        nil,
		}
		input, _ := json.Marshal(presetData)
		_,err := dev.CallOnvifFunction(onvif20.PTZWebService, onvif20.GotoPreset, input)
		fmt.Println(err)
	}
}