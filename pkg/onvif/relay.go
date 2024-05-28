package onvif


func (device Device) GetReplayConfiguration() (interface{}, error) {
	// create soap
	soap := SOAP{
		User:     device.User,
		Password: device.Password,
		Body:     `<GetReplayConfiguration xmlns="http://www.onvif.org/ver10/replay/wsdl"/>`,
	}

	var result interface{}
	// send request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return result, err
	}

	// parse response
	_, err = response.ValueForPath("Envelope.Body.GetReplayConfigurationResponse")
	if err != nil {
		return result, err
	}
	return result, nil
}

func (device Device) GetReplayServiceCapabilities() (interface{}, error) {
	// create soap
	soap := SOAP{
		User:     device.User,
		Password: device.Password,
		Body:     `<GetServiceCapabilities xmlns="http://www.onvif.org/ver10/replay/wsdl"/>`,
	}

	var result interface{}
	// send request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return result, err
	}

	// parse response
	_, err = response.ValueForPath("Envelope.Body.GetServiceCapabilitiesResponse")
	if err != nil {
		return result, err
	}
	return result, nil
}

func (device Device) GetReplayUri(recordingToken string) (string, error) {
	// create soap
	soap := SOAP{
		User:     device.User,
		Password: device.Password,
		Body: `<GetReplayUri xmlns="http://www.onvif.org/ver10/replay/wsdl">
						<RecordingToken>` + recordingToken + `</RecordingToken>		
						<StreamSetup>
							<Stream xmlns="http://www.onvif.org/ver10/schema">RTP-Unicast</Stream>
							<Transport xmlns="http://www.onvif.org/ver10/schema">
								<Protocol xmlns="http://www.onvif.org/ver10/schema">TCP</Protocol>
							</Transport>
						</StreamSetup>
					</GetReplayUri>`,
	}

	var result = ""
	// send request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return result, err
	}

	// parse response
	data, err := response.ValueForPath("Envelope.Body.GetReplayUriResponse.Uri")
	if err != nil {
		return result, err
	}
	return interfaceToString(data), nil
}
