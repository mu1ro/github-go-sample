package go_ffprobe_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	ffprobe "github.com/mu1ro/github-go"
)

// jsonStr is sample exec ffprobe result.
var jsonStr = `
{
    "streams": [
        {
            "index": 0,
            "codec_name": "h264",
            "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
            "profile": "Main",
            "codec_type": "video",
            "codec_time_base": "1/30",
            "codec_tag_string": "avc1",
            "codec_tag": "0x31637661",
            "width": 320,
            "height": 240,
            "coded_width": 320,
            "coded_height": 240,
            "has_b_frames": 0,
            "sample_aspect_ratio": "1:1",
            "display_aspect_ratio": "4:3",
            "pix_fmt": "yuv420p",
            "level": 13,
            "chroma_location": "left",
            "refs": 1,
            "is_avc": "true",
            "nal_length_size": "4",
            "r_frame_rate": "15/1",
            "avg_frame_rate": "15/1",
            "time_base": "1/15360",
            "start_pts": 0,
            "start_time": "0.000000",
            "duration_ts": 209925,
            "duration": "13.666992",
            "bit_rate": "229387",
            "bits_per_raw_sample": "8",
            "nb_frames": "205",
            "disposition": {
                "default": 1,
                "dub": 0,
                "original": 0,
                "comment": 0,
                "lyrics": 0,
                "karaoke": 0,
                "forced": 0,
                "hearing_impaired": 0,
                "visual_impaired": 0,
                "clean_effects": 0,
                "attached_pic": 0,
                "timed_thumbnails": 0
            },
            "tags": {
                "creation_time": "1970-01-01T00:00:00.000000Z",
                "language": "und",
                "handler_name": "VideoHandler"
            }
        },
        {
            "index": 1,
            "codec_name": "aac",
            "codec_long_name": "AAC (Advanced Audio Coding)",
            "profile": "LC",
            "codec_type": "audio",
            "codec_time_base": "1/48000",
            "codec_tag_string": "mp4a",
            "codec_tag": "0x6134706d",
            "sample_fmt": "fltp",
            "sample_rate": "48000",
            "channels": 6,
            "channel_layout": "5.1",
            "bits_per_sample": 0,
            "r_frame_rate": "0/0",
            "avg_frame_rate": "0/0",
            "time_base": "1/48000",
            "start_pts": 0,
            "start_time": "0.000000",
            "duration_ts": 657408,
            "duration": "13.696000",
            "bit_rate": "382488",
            "max_bit_rate": "416704",
            "nb_frames": "642",
            "disposition": {
                "default": 1,
                "dub": 0,
                "original": 0,
                "comment": 0,
                "lyrics": 0,
                "karaoke": 0,
                "forced": 0,
                "hearing_impaired": 0,
                "visual_impaired": 0,
                "clean_effects": 0,
                "attached_pic": 0,
                "timed_thumbnails": 0
            },
            "tags": {
                "creation_time": "1970-01-01T00:00:00.000000Z",
                "language": "und",
                "handler_name": "SoundHandler"
            }
        }
    ],
    "format": {
        "filename": "sample.mp4",
        "nb_streams": 2,
        "nb_programs": 0,
        "format_name": "mov,mp4,m4a,3gp,3g2,mj2",
        "format_long_name": "QuickTime / MOV",
        "start_time": "0.000000",
        "duration": "13.696000",
        "size": "1053651",
        "bit_rate": "615450",
        "probe_score": 100,
        "tags": {
            "major_brand": "isom",
            "minor_version": "512",
            "compatible_brands": "isomiso2avc1mp41",
            "creation_time": "1970-01-01T00:00:00.000000Z",
            "encoder": "Lavf53.24.2"
        }
    }
}
`

// setExecFunc set mock func to ExecFunc
func setExecFunc(resultJson string, err error) {
	ffprobe.ExecFunc = func(f string) ([]byte, error) { return []byte(resultJson), err }
}

func TestExecute(t *testing.T) {
	// ready exp object.
	jsonBytes := ([]byte)(jsonStr)
	exp := ffprobe.Result{}

	if err := json.Unmarshal(jsonBytes, &exp); err != nil {
		t.Errorf("JSON Unmarshal error: %v", err)
		return
	}

	var tests = []struct {
		exp        ffprobe.Result
		resultJson string
		expErr     error
	}{
		{
			exp:        exp,
			resultJson: jsonStr,
			expErr:     nil,
		},
		{
			exp:        ffprobe.Result{},
			resultJson: "",
			expErr:     errors.New("failed"),
		},
		{
			exp:        ffprobe.Result{},
			resultJson: "-",
			expErr:     nil,
		},
	}

	for _, test := range tests {
		setExecFunc(test.resultJson, test.expErr)

		act, err := ffprobe.Execute("-")

		if err == nil {
			if !reflect.DeepEqual(act, test.exp) {
				t.Errorf(`
					failed.
					act. %v
					exp. %v
				`, act, test.exp)
			}
			continue
		}
	}
}
