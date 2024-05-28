package ffprobe

import (
	"time"
)

var ffprobeBinPath = "ffprobe"

type StreamType string

const (
	StreamAny        StreamType = ""
	StreamVideo      StreamType = "video"
	StreamAudio      StreamType = "audio"
	StreamSubtitle   StreamType = "subtitle"
	StreamData       StreamType = "data"
	StreamAttachment StreamType = "attachment"
)

type ProbeData struct {
	Streams []*Stream `json:"streams"`
	Format  *Format   `json:"format"`
}

type Format struct {
	Filename         string  `json:"filename"`
	NBStreams        int     `json:"nb_streams"`
	NBPrograms       int     `json:"nb_programs"`
	FormatName       string  `json:"format_name"`
	FormatLongName   string  `json:"format_long_name"`
	StartTimeSeconds float64 `json:"start_time,string"`
	DurationSeconds  float64 `json:"duration,string"`
	Size             string  `json:"size"`
	BitRate          string  `json:"bit_rate"`
	ProbeScore       int     `json:"probe_score"`
	TagList          Tags    `json:"tags"`
}

type Stream struct {
	Index              int               `json:"index"`
	ID                 string            `json:"id"`
	CodecName          string            `json:"codec_name"`
	CodecLongName      string            `json:"codec_long_name"`
	CodecType          string            `json:"codec_type"`
	CodecTimeBase      string            `json:"codec_time_base"`
	CodecTagString     string            `json:"codec_tag_string"`
	CodecTag           string            `json:"codec_tag"`
	RFrameRate         string            `json:"r_frame_rate"`
	AvgFrameRate       string            `json:"avg_frame_rate"`
	TimeBase           string            `json:"time_base"`
	StartPts           int               `json:"start_pts"`
	StartTime          string            `json:"start_time"`
	DurationTs         uint64            `json:"duration_ts"`
	Duration           string            `json:"duration"`
	BitRate            string            `json:"bit_rate"`
	BitsPerRawSample   string            `json:"bits_per_raw_sample"`
	NbFrames           string            `json:"nb_frames"`
	Disposition        StreamDisposition `json:"disposition,omitempty"`
	TagList            Tags              `json:"tags"`
	FieldOrder         string            `json:"field_order,omitempty"`
	Profile            string            `json:"profile,omitempty"`
	Width              int               `json:"width"`
	Height             int               `json:"height"`
	HasBFrames         int               `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string            `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string            `json:"display_aspect_ratio,omitempty"`
	PixFmt             string            `json:"pix_fmt,omitempty"`
	Level              int               `json:"level,omitempty"`
	ColorRange         string            `json:"color_range,omitempty"`
	ColorSpace         string            `json:"color_space,omitempty"`
	SampleFmt          string            `json:"sample_fmt,omitempty"`
	SampleRate         string            `json:"sample_rate,omitempty"`
	Channels           int               `json:"channels,omitempty"`
	ChannelLayout      string            `json:"channel_layout,omitempty"`
	BitsPerSample      int               `json:"bits_per_sample,omitempty"`
}

type StreamDisposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
}

func (f *Format) StartTime() (duration time.Duration) {
	return time.Duration(f.StartTimeSeconds * float64(time.Second))
}

func (f *Format) Duration() (duration time.Duration) {
	return time.Duration(f.DurationSeconds * float64(time.Second))
}

func (p *ProbeData) StreamType(streamType StreamType) (streams []Stream) {
	for _, s := range p.Streams {
		if s == nil {
			continue
		}
		switch streamType {
		case StreamAny:
			streams = append(streams, *s)
		default:
			if s.CodecType == string(streamType) {
				streams = append(streams, *s)
			}
		}
	}
	return streams
}

func (p *ProbeData) FirstVideoStream() *Stream {
	return p.firstStream(StreamVideo)
}

func (p *ProbeData) FirstAudioStream() *Stream {
	return p.firstStream(StreamAudio)
}

func (p *ProbeData) FirstSubtitleStream() *Stream {
	return p.firstStream(StreamSubtitle)
}

func (p *ProbeData) FirstDataStream() *Stream {
	return p.firstStream(StreamData)
}

func (p *ProbeData) FirstAttachmentStream() *Stream {
	return p.firstStream(StreamAttachment)
}

func (p *ProbeData) firstStream(streamType StreamType) *Stream {
	for _, s := range p.Streams {
		if s == nil {
			continue
		}
		if s.CodecType == string(streamType) {
			return s
		}
	}
	return nil
}
