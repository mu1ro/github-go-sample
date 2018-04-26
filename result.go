package go_ffprobe

import "time"

type Result struct {
	Format struct {
		BitRate        string `json:"bit_rate"`
		Duration       string `json:"duration"`
		Filename       string `json:"filename"`
		FormatLongName string `json:"format_long_name"`
		FormatName     string `json:"format_name"`
		NbPrograms     int    `json:"nb_programs"`
		NbStreams      int    `json:"nb_streams"`
		ProbeScore     int    `json:"probe_score"`
		Size           string `json:"size"`
		StartTime      string `json:"start_time"`
		Tags           struct {
			CompatibleBrands string    `json:"compatible_brands"`
			CreationTime     time.Time `json:"creation_time"`
			Encoder          string    `json:"encoder"`
			MajorBrand       string    `json:"major_brand"`
			MinorVersion     string    `json:"minor_version"`
		} `json:"tags"`
	} `json:"format"`
	Streams []struct {
		AvgFrameRate       string `json:"avg_frame_rate"`
		BitRate            string `json:"bit_rate"`
		BitsPerRawSample   string `json:"bits_per_raw_sample,omitempty"`
		ChromaLocation     string `json:"chroma_location,omitempty"`
		CodecLongName      string `json:"codec_long_name"`
		CodecName          string `json:"codec_name"`
		CodecTag           string `json:"codec_tag"`
		CodecTagString     string `json:"codec_tag_string"`
		CodecTimeBase      string `json:"codec_time_base"`
		CodecType          string `json:"codec_type"`
		CodedHeight        int    `json:"coded_height,omitempty"`
		CodedWidth         int    `json:"coded_width,omitempty"`
		DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
		Disposition        struct {
			AttachedPic     int `json:"attached_pic"`
			CleanEffects    int `json:"clean_effects"`
			Comment         int `json:"comment"`
			Default         int `json:"default"`
			Dub             int `json:"dub"`
			Forced          int `json:"forced"`
			HearingImpaired int `json:"hearing_impaired"`
			Karaoke         int `json:"karaoke"`
			Lyrics          int `json:"lyrics"`
			Original        int `json:"original"`
			TimedThumbnails int `json:"timed_thumbnails"`
			VisualImpaired  int `json:"visual_impaired"`
		} `json:"disposition"`
		Duration          string `json:"duration"`
		DurationTs        int    `json:"duration_ts"`
		HasBFrames        int    `json:"has_b_frames,omitempty"`
		Height            int    `json:"height,omitempty"`
		Index             int    `json:"index"`
		IsAvc             string `json:"is_avc,omitempty"`
		Level             int    `json:"level,omitempty"`
		NalLengthSize     string `json:"nal_length_size,omitempty"`
		NbFrames          string `json:"nb_frames"`
		PixFmt            string `json:"pix_fmt,omitempty"`
		Profile           string `json:"profile"`
		RFrameRate        string `json:"r_frame_rate"`
		Refs              int    `json:"refs,omitempty"`
		SampleAspectRatio string `json:"sample_aspect_ratio,omitempty"`
		StartPts          int    `json:"start_pts"`
		StartTime         string `json:"start_time"`
		Tags              struct {
			CreationTime time.Time `json:"creation_time"`
			HandlerName  string    `json:"handler_name"`
			Language     string    `json:"language"`
		} `json:"tags"`
		TimeBase      string `json:"time_base"`
		Width         int    `json:"width,omitempty"`
		BitsPerSample int    `json:"bits_per_sample,omitempty"`
		ChannelLayout string `json:"channel_layout,omitempty"`
		Channels      int    `json:"channels,omitempty"`
		MaxBitRate    string `json:"max_bit_rate,omitempty"`
		SampleFmt     string `json:"sample_fmt,omitempty"`
		SampleRate    string `json:"sample_rate,omitempty"`
	} `json:"streams"`
}
