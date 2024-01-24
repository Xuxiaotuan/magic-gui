package util

import ffmpeg "github.com/u2takey/ffmpeg-go"

func CovertMp4ToMp4(source string, target string) {
	err := ffmpeg.Input(source).
		Output(target, ffmpeg.KwArgs{"c:v": "libx265"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return
	}
}

func CovertMp4ToMp3(source string, target string) {
	err := ffmpeg.Input(source).
		Output(target).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return
	}
}
