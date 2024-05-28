package wsdiscovery

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	onvif20 "test-func/pkg/onvif-20"
	"time"

	"github.com/beevik/etree"
	"github.com/google/uuid"
	"golang.org/x/net/ipv4"
)

const (
	bufSize = 8192
)

// GetAvailableDevicesAtSpecificEthernetInterface sends a ws-discovery Probe Message via
// UDP multicast to Discover NVT type Devices
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []onvif20.Device {
	types := []string{"dn:NetworkVideoTransmitter"}
	namespaces := map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl", "ds": "http://www.onvif.org/ver10/device/wsdl"}

	probeResponses := SendProbe(interfaceName, nil, types, namespaces)
	for _, val := range probeResponses{
		fmt.Println("Meomeo: ", val)
	}
	nvtDevices := make([]onvif20.Device, 0)
	nvtDevices, err := DevicesFromProbeResponses(probeResponses)
	if err != nil {
		fmt.Printf("Failed to discover Onvif camera: %s\n", err.Error())
	}

	return nvtDevices
}

func DevicesFromProbeResponses(probeResponses []string) ([]onvif20.Device, error) {
	nvtDevices := make([]onvif20.Device, 0)
	xaddrSet := make(map[string]struct{})
	for _, j := range probeResponses {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return nil, err
		}

		probeMatches := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch")
		for _, probeMatch := range probeMatches {
			var xaddr string
			if address := probeMatch.FindElement("./XAddrs"); address != nil {
				u, err := url.Parse(address.Text())
				if err != nil {
					fmt.Printf("Invalid XAddrs: %s\n", address.Text())
					continue
				}
				xaddr = u.Host
			}
			if _, dupe := xaddrSet[xaddr]; dupe {
				fmt.Printf("Skipping duplicate XAddr: %s\n", xaddr)
				continue
			}

			var endpointRefAddress string
			if ref := probeMatch.FindElement("./EndpointReference/Address"); ref != nil {
				uuidElements := strings.Split(ref.Text(), ":")
				endpointRefAddress = uuidElements[len(uuidElements)-1]
			}

			dev, err := onvif20.NewDevice(onvif20.DeviceParams{
				Xaddr:              xaddr,
				EndpointRefAddress: endpointRefAddress,
				HttpClient: &http.Client{
					Timeout: 2 * time.Second,
				},
			})
			if err != nil {
				fmt.Printf("Failed to connect to camera at %s: %s\n", xaddr, err.Error())
				continue
			}
			xaddrSet[xaddr] = struct{}{}
			nvtDevices = append(nvtDevices, *dev)
			fmt.Printf("Onvif WS-Discovery: Find Xaddr: %-25s EndpointRefAddress: %s\n", xaddr, string(endpointRefAddress))
		}
	}

	return nvtDevices, nil
}

// SendProbe to device
func SendProbe(interfaceName string, scopes, types []string, namespaces map[string]string) []string {
	probeSOAP := BuildProbeMessage(uuid.NewString(), scopes, types, namespaces)
	return SendUDPMulticast(probeSOAP.String(), interfaceName)
}

func SendUDPMulticast(msg string, interfaceName string) []string {
	var responses []string
	data := []byte(msg)

	c, err := net.ListenPacket("udp4", "0.0.0.0:0")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	p := ipv4.NewPacketConn(c)

	// 239.255.255.250 port 3702 is the multicast address and port used by ws-discovery
	group := net.IPv4(239, 255, 255, 250)
	dest := &net.UDPAddr{IP: group, Port: 3702}

	var iface *net.Interface
	if interfaceName == "" {
		iface = nil
	} else {
		iface, err = net.InterfaceByName(interfaceName)
		if err != nil {
			fmt.Printf("Error calling InterfaceByName for interface %q: %s\n", interfaceName, err.Error())
		}
	}

	if err = p.JoinGroup(iface, &net.UDPAddr{IP: group}); err != nil {
		fmt.Printf("Error calling JoinGroup for ws-discovery: %s\n", err.Error())
	}
	if iface != nil {
		if err = p.SetMulticastInterface(iface); err != nil {
			fmt.Printf("Error calling SetMulticastInterface for interface %q: %s\n", interfaceName, err.Error())
		}
		if err = p.SetMulticastTTL(2); err != nil {
			fmt.Printf("Error calling SetMulticastTTL: %s\n", err.Error())
		}
	}
	if _, err = p.WriteTo(data, nil, dest); err != nil {
		fmt.Printf("Error writing to ws-discovery multicast address %s: %s\n", dest.String(), err.Error())
	}

	if err = p.SetReadDeadline(time.Now().Add(time.Second * 5)); err != nil {
		fmt.Printf("Error setting read deadline: %s\n", err.Error())
		return nil
	}

	b := make([]byte, bufSize)

	// keep reading from the PacketConn until the read deadline expires or an error occurs
	for {
		n, _, _, err := p.ReadFrom(b)
		if err != nil {
			// ErrDeadlineExceeded is expected once the read timeout is expired
			if !errors.Is(err, os.ErrDeadlineExceeded) {
				fmt.Printf("Unexpected error occurred while reading ws-discovery responses: %s\n", err.Error())
			}
			break
		}
		responses = append(responses, string(b[0:n]))
	}
	return responses
}
