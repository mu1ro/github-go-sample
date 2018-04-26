package go_ffprobe

import (
	"encoding/json"
	"os/exec"
)

var (
	// ExecFunc is command func.
	ExecFunc = ExecCmd
)

// Execute exec command and bind result to struct.
func Execute(fileName string) (r Result, err error) {
	out, err := ExecFunc(fileName)

	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(out, &r); err != nil {
		return r, err
	}

	return r, nil
}

// ExecCmd exec ffprobe command and return result of json.
func ExecCmd(fileName string) ([]byte, error) {
	return exec.Command("ffprobe",
		"-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", fileName).Output()
}
