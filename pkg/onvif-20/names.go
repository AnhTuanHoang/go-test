package onvif20

// Onvif WebService
const (
	AnalyticsWebService = "Analytics"
	DeviceWebService    = "Device"
	EventWebService     = "Event"
	ImagingWebService   = "Imaging"
	MediaWebService     = "Media"
	Media2WebService    = "Media2"
	PTZWebService       = "PTZ"
)

// WebService - Analytics
const (
	CreateAnalyticsModules       = "CreateAnalyticsModules"
	CreateRules                  = "CreateRules"
	DeleteAnalyticsModules       = "DeleteAnalyticsModules"
	DeleteRules                  = "DeleteRules"
	GetAnalyticsModuleOptions    = "GetAnalyticsModuleOptions"
	GetAnalyticsModules          = "GetAnalyticsModules"
	GetRuleOptions               = "GetRuleOptions"
	GetRules                     = "GetRules"
	GetSupportedAnalyticsModules = "GetSupportedAnalyticsModules"
	GetSupportedRules            = "GetSupportedRules"
	ModifyAnalyticsModules       = "ModifyAnalyticsModules"
	ModifyRules                  = "ModifyRules"
)

// WebService - Device
const (
	AddIPAddressFilter            = "AddIPAddressFilter"
	AddScopes                     = "AddScopes"
	CreateCertificate             = "CreateCertificate"
	CreateDot1XConfiguration      = "CreateDot1XConfiguration"
	CreateStorageConfiguration    = "CreateStorageConfiguration"
	CreateUsers                   = "CreateUsers"
	DeleteCertificates            = "DeleteCertificates"
	DeleteDot1XConfiguration      = "DeleteDot1XConfiguration"
	DeleteGeoLocation             = "DeleteGeoLocation"
	DeleteStorageConfiguration    = "DeleteStorageConfiguration"
	DeleteUsers                   = "DeleteUsers"
	GetAccessPolicy               = "GetAccessPolicy"
	GetCACertificates             = "GetCACertificates"
	GetCapabilities               = "GetCapabilities"
	GetCertificateInformation     = "GetCertificateInformation"
	GetCertificates               = "GetCertificates"
	GetCertificatesStatus         = "GetCertificatesStatus"
	GetClientCertificateMode      = "GetClientCertificateMode"
	GetDNS                        = "GetDNS"
	GetDPAddresses                = "GetDPAddresses"
	GetDeviceInformation          = "GetDeviceInformation"
	GetDiscoveryMode              = "GetDiscoveryMode"
	GetDot11Capabilities          = "GetDot11Capabilities"
	GetDot11Status                = "GetDot11Status"
	GetDot1XConfiguration         = "GetDot1XConfiguration"
	GetDot1XConfigurations        = "GetDot1XConfigurations"
	GetDynamicDNS                 = "GetDynamicDNS"
	GetEndpointReference          = "GetEndpointReference"
	GetGeoLocation                = "GetGeoLocation"
	GetHostname                   = "GetHostname"
	GetIPAddressFilter            = "GetIPAddressFilter"
	GetNTP                        = "GetNTP"
	GetNetworkDefaultGateway      = "GetNetworkDefaultGateway"
	GetNetworkInterfaces          = "GetNetworkInterfaces"
	GetNetworkProtocols           = "GetNetworkProtocols"
	GetPkcs10Request              = "GetPkcs10Request"
	GetRelayOutputs               = "GetRelayOutputs"
	GetRemoteDiscoveryMode        = "GetRemoteDiscoveryMode"
	GetRemoteUser                 = "GetRemoteUser"
	GetScopes                     = "GetScopes"
	GetServiceCapabilities        = "GetServiceCapabilities"
	GetServices                   = "GetServices"
	GetStorageConfiguration       = "GetStorageConfiguration"
	GetStorageConfigurations      = "GetStorageConfigurations"
	GetSystemBackup               = "GetSystemBackup"
	GetSystemDateAndTime          = "GetSystemDateAndTime"
	GetSystemLog                  = "GetSystemLog"
	GetSystemSupportInformation   = "GetSystemSupportInformation"
	GetSystemUris                 = "GetSystemUris"
	GetUsers                      = "GetUsers"
	GetWsdlUrl                    = "GetWsdlUrl"
	GetZeroConfiguration          = "GetZeroConfiguration"
	LoadCACertificates            = "LoadCACertificates"
	LoadCertificateWithPrivateKey = "LoadCertificateWithPrivateKey"
	LoadCertificates              = "LoadCertificates"
	RemoveIPAddressFilter         = "RemoveIPAddressFilter"
	RemoveScopes                  = "RemoveScopes"
	RestoreSystem                 = "RestoreSystem"
	ScanAvailableDot11Networks    = "ScanAvailableDot11Networks"
	SendAuxiliaryCommand          = "SendAuxiliaryCommand"
	SetAccessPolicy               = "SetAccessPolicy"
	SetCertificatesStatus         = "SetCertificatesStatus"
	SetClientCertificateMode      = "SetClientCertificateMode"
	SetDNS                        = "SetDNS"
	SetDPAddresses                = "SetDPAddresses"
	SetDiscoveryMode              = "SetDiscoveryMode"
	SetDot1XConfiguration         = "SetDot1XConfiguration"
	SetDynamicDNS                 = "SetDynamicDNS"
	SetGeoLocation                = "SetGeoLocation"
	SetHostname                   = "SetHostname"
	SetHostnameFromDHCP           = "SetHostnameFromDHCP"
	SetIPAddressFilter            = "SetIPAddressFilter"
	SetNTP                        = "SetNTP"
	SetNetworkDefaultGateway      = "SetNetworkDefaultGateway"
	SetNetworkInterfaces          = "SetNetworkInterfaces"
	SetNetworkProtocols           = "SetNetworkProtocols"
	SetRelayOutputSettings        = "SetRelayOutputSettings"
	SetRelayOutputState           = "SetRelayOutputState"
	SetRemoteDiscoveryMode        = "SetRemoteDiscoveryMode"
	SetRemoteUser                 = "SetRemoteUser"
	SetScopes                     = "SetScopes"
	SetStorageConfiguration       = "SetStorageConfiguration"
	SetSystemDateAndTime          = "SetSystemDateAndTime"
	SetSystemFactoryDefault       = "SetSystemFactoryDefault"
	SetUser                       = "SetUser"
	SetZeroConfiguration          = "SetZeroConfiguration"
	StartFirmwareUpgrade          = "StartFirmwareUpgrade"
	StartSystemRestore            = "StartSystemRestore"
	SystemReboot                  = "SystemReboot"
	UpgradeSystemFirmware         = "UpgradeSystemFirmware"
)

// WebService - Event
const (
	CreatePullPointSubscription = "CreatePullPointSubscription"
	GetEventProperties          = "GetEventProperties"
	PullMessages                = "PullMessages"
	Renew                       = "Renew"
	Seek                        = "Seek"
	SetSynchronizationPoint     = "SetSynchronizationPoint"
	Subscribe                   = "Subscribe"
	SubscriptionReference       = "SubscriptionReference"
	Unsubscribe                 = "Unsubscribe"
)

// WebService - Imaging
const (
	GetCurrentPreset   = "GetCurrentPreset"
	GetImagingSettings = "GetImagingSettings"
	GetMoveOptions     = "GetMoveOptions"
	GetOptions         = "GetOptions"
	GetPresets         = "GetPresets"
	GetStatus          = "GetStatus"
	Move               = "Move"
	SetCurrentPreset   = "SetCurrentPreset"
	SetImagingSettings = "SetImagingSettings"
	Stop               = "Stop"
)

// WebService - Media
const (
	AddAudioDecoderConfiguration               = "AddAudioDecoderConfiguration"
	AddAudioEncoderConfiguration               = "AddAudioEncoderConfiguration"
	AddAudioOutputConfiguration                = "AddAudioOutputConfiguration"
	AddAudioSourceConfiguration                = "AddAudioSourceConfiguration"
	AddMetadataConfiguration                   = "AddMetadataConfiguration"
	AddPTZConfiguration                        = "AddPTZConfiguration"
	AddVideoAnalyticsConfiguration             = "AddVideoAnalyticsConfiguration"
	AddVideoEncoderConfiguration               = "AddVideoEncoderConfiguration"
	AddVideoSourceConfiguration                = "AddVideoSourceConfiguration"
	CreateOSD                                  = "CreateOSD"
	CreateProfile                              = "CreateProfile"
	DeleteOSD                                  = "DeleteOSD"
	DeleteProfile                              = "DeleteProfile"
	GetAudioDecoderConfiguration               = "GetAudioDecoderConfiguration"
	GetAudioDecoderConfigurationOptions        = "GetAudioDecoderConfigurationOptions"
	GetAudioDecoderConfigurations              = "GetAudioDecoderConfigurations"
	GetAudioEncoderConfiguration               = "GetAudioEncoderConfiguration"
	GetAudioEncoderConfigurationOptions        = "GetAudioEncoderConfigurationOptions"
	GetAudioEncoderConfigurations              = "GetAudioEncoderConfigurations"
	GetAudioOutputConfiguration                = "GetAudioOutputConfiguration"
	GetAudioOutputConfigurationOptions         = "GetAudioOutputConfigurationOptions"
	GetAudioOutputConfigurations               = "GetAudioOutputConfigurations"
	GetAudioOutputs                            = "GetAudioOutputs"
	GetAudioSourceConfiguration                = "GetAudioSourceConfiguration"
	GetAudioSourceConfigurationOptions         = "GetAudioSourceConfigurationOptions"
	GetAudioSourceConfigurations               = "GetAudioSourceConfigurations"
	GetAudioSources                            = "GetAudioSources"
	GetCompatibleAudioDecoderConfigurations    = "GetCompatibleAudioDecoderConfigurations"
	GetCompatibleAudioEncoderConfigurations    = "GetCompatibleAudioEncoderConfigurations"
	GetCompatibleAudioOutputConfigurations     = "GetCompatibleAudioOutputConfigurations"
	GetCompatibleAudioSourceConfigurations     = "GetCompatibleAudioSourceConfigurations"
	GetCompatibleMetadataConfigurations        = "GetCompatibleMetadataConfigurations"
	GetCompatibleVideoAnalyticsConfigurations  = "GetCompatibleVideoAnalyticsConfigurations"
	GetCompatibleVideoEncoderConfigurations    = "GetCompatibleVideoEncoderConfigurations"
	GetCompatibleVideoSourceConfigurations     = "GetCompatibleVideoSourceConfigurations"
	GetGuaranteedNumberOfVideoEncoderInstances = "GetGuaranteedNumberOfVideoEncoderInstances"
	GetMetadataConfiguration                   = "GetMetadataConfiguration"
	GetMetadataConfigurationOptions            = "GetMetadataConfigurationOptions"
	GetMetadataConfigurations                  = "GetMetadataConfigurations"
	GetOSD                                     = "GetOSD"
	GetOSDOptions                              = "GetOSDOptions"
	GetOSDs                                    = "GetOSDs"
	GetProfile                                 = "GetProfile"
	GetProfiles                                = "GetProfiles"
	GetSnapshotUri                             = "GetSnapshotUri"
	GetStreamUri                               = "GetStreamUri"
	GetVideoAnalyticsConfiguration             = "GetVideoAnalyticsConfiguration"
	GetVideoAnalyticsConfigurations            = "GetVideoAnalyticsConfigurations"
	GetVideoEncoderConfiguration               = "GetVideoEncoderConfiguration"
	GetVideoEncoderConfigurationOptions        = "GetVideoEncoderConfigurationOptions"
	GetVideoEncoderConfigurations              = "GetVideoEncoderConfigurations"
	GetVideoSourceConfiguration                = "GetVideoSourceConfiguration"
	GetVideoSourceConfigurationOptions         = "GetVideoSourceConfigurationOptions"
	GetVideoSourceConfigurations               = "GetVideoSourceConfigurations"
	GetVideoSourceModes                        = "GetVideoSourceModes"
	GetVideoSources                            = "GetVideoSources"
	RemoveAudioDecoderConfiguration            = "RemoveAudioDecoderConfiguration"
	RemoveAudioEncoderConfiguration            = "RemoveAudioEncoderConfiguration"
	RemoveAudioOutputConfiguration             = "RemoveAudioOutputConfiguration"
	RemoveAudioSourceConfiguration             = "RemoveAudioSourceConfiguration"
	RemoveMetadataConfiguration                = "RemoveMetadataConfiguration"
	RemovePTZConfiguration                     = "RemovePTZConfiguration"
	RemoveVideoAnalyticsConfiguration          = "RemoveVideoAnalyticsConfiguration"
	RemoveVideoEncoderConfiguration            = "RemoveVideoEncoderConfiguration"
	RemoveVideoSourceConfiguration             = "RemoveVideoSourceConfiguration"
	SetAudioDecoderConfiguration               = "SetAudioDecoderConfiguration"
	SetAudioEncoderConfiguration               = "SetAudioEncoderConfiguration"
	SetAudioOutputConfiguration                = "SetAudioOutputConfiguration"
	SetAudioSourceConfiguration                = "SetAudioSourceConfiguration"
	SetMetadataConfiguration                   = "SetMetadataConfiguration"
	SetOSD                                     = "SetOSD"
	SetVideoAnalyticsConfiguration             = "SetVideoAnalyticsConfiguration"
	SetVideoEncoderConfiguration               = "SetVideoEncoderConfiguration"
	SetVideoSourceConfiguration                = "SetVideoSourceConfiguration"
	SetVideoSourceMode                         = "SetVideoSourceMode"
	StartMulticastStreaming                    = "StartMulticastStreaming"
	StopMulticastStreaming                     = "StopMulticastStreaming"
)

// WebService - Media2
const (
	AddConfiguration           = "AddConfiguration"
	GetAnalyticsConfigurations = "GetAnalyticsConfigurations"
	RemoveConfiguration        = "RemoveConfiguration"
)

// WebService - PTZ
const (
	AbsoluteMove                = "AbsoluteMove"
	ContinuousMove              = "ContinuousMove"
	CreatePresetTour            = "CreatePresetTour"
	GeoMove                     = "GeoMove"
	GetCompatibleConfigurations = "GetCompatibleConfigurations"
	GetConfiguration            = "GetConfiguration"
	GetConfigurationOptions     = "GetConfigurationOptions"
	GetConfigurations           = "GetConfigurations"
	GetNode                     = "GetNode"
	GetNodes                    = "GetNodes"
	GetPresetTour               = "GetPresetTour"
	GetPresetTourOptions        = "GetPresetTourOptions"
	GetPresetTours              = "GetPresetTours"
	GotoHomePosition            = "GotoHomePosition"
	GotoPreset                  = "GotoPreset"
	ModifyPresetTour            = "ModifyPresetTour"
	OperatePresetTour           = "OperatePresetTour"
	RelativeMove                = "RelativeMove"
	RemovePreset                = "RemovePreset"
	RemovePresetTour            = "RemovePresetTour"
	SetConfiguration            = "SetConfiguration"
	SetHomePosition             = "SetHomePosition"
	SetPreset                   = "SetPreset"
)
