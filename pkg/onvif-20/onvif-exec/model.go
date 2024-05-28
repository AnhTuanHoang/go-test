package main


type StreamUri struct {
	MediaURI MediaURI
}

type MediaURI struct {
	URI                 string
	Timeout             string
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
}

type Option struct {
	Options VideoEncoderConfigurationOptions
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange
	H264         H264Options
}

type H264Options struct {
	ResolutionsAvailable  []MediaBounds
	GovLengthRange        IntRange
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
	BitrateRange          IntRange
	H264ProfilesSupported []string // 'Baseline', 'Main', 'Extended', 'High'
}

type IntRange struct {
	Min int
	Max int
}

type ProfileRS struct {
	Profiles []string
}

type Profile struct {
	Profiles []MediaProfile
}

type MediaProfile struct {
	Name               string
	Token              string
	VideoSourceConfiguration  MediaSourceConfig
	VideoEncoderConfiguration VideoEncoderConfig
	AudioSourceConfiguration  MediaSourceConfig
	AudioEncoderConfiguration AudioEncoderConfig
	PTZConfiguration          PTZConfig
}

type MediaSourceConfig struct {
	Name        string
	Token       string
	SourceToken string
	Bounds      MediaBounds
}

type VideoEncoderConfig struct {
	Name                string
	Token               string
	Encoding            string
	Quality             float64
	RateControl         VideoRateControl
	Resolution          MediaBounds
	SessionTimeout      string
	H264                H264Configuration
	Multicast           Multicast
	GuaranteedFrameRate bool
	UseCount            int
}

type AudioEncoderConfig struct {
	Name           string
	Token          string
	Encoding       string
	Bitrate        int
	SampleRate     int
	SessionTimeout string
}

type PTZConfig struct {
	Name      string
	Token     string
	NodeToken string
}

type MediaBounds struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type VideoRateControl struct {
	BitrateLimit     int
	EncodingInterval int
	FrameRateLimit   int
}

type Multicast struct {
	Address   IPAddress
	Port      int
	TTL       int
	AutoStart bool
}

type H264Configuration struct {
	GovLength   int
	H264Profile string //'Baseline', 'Main', 'Extended', 'High'
}

type IPAddress struct {
	Type        string
	IPv4Address string
}

type StreamUriResult struct {
	Uri string
	Profile string
}