package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	processing "video_processing/grpc/server"
)

func (a *Application) ProcessVideo(ctx context.Context, vd *processing.VideoData) (*processing.Response, error) {
	go a.HLS(vd)

	return &processing.Response{Message: "Processing"}, nil
}

func (a *Application) HLS(vd *processing.VideoData) {
	inputFile := vd.FilePath
	outputDir := "/app/data/videos/"

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Error creating output directory:", err)
	}

	cmd := exec.Command("ffmpeg", "-y", "-i", inputFile,
		"-filter_complex", "[0:v]split=3[720_in][480_in][240_in];"+
			"[720_in]scale=-2:720[720_out];"+
			"[480_in]scale=-2:480[480_out];"+
			"[240_in]scale=-2:240[240_out]",
		"-map", "[720_out]", "-map", "[480_out]", "-map", "[240_out]",
		"-map", "0:a", "-map", "0:a", "-map", "0:a",
		"-b:v:0", "3500k", "-maxrate:v:0", "3500k", "-bufsize:v:0", "3500k",
		"-b:v:1", "1690k", "-maxrate:v:1", "1690k", "-bufsize:v:1", "1690k",
		"-b:v:2", "326k", "-maxrate:v:2", "326k", "-bufsize:v:2", "326k",
		"-b:a:0", "128k", "-b:a:1", "128k", "-b:a:2", "128k",
		"-x264-params", "keyint=60:min-keyint=60:scenecut=0",
		"-var_stream_map", "v:0,a:0,name:720p-4M v:1,a:1,name:480p-2M v:2,a:2,name:240p-500k",
		"-hls_list_size", "0", "-hls_time", "2",
		"-hls_segment_filename", outputDir+"adaptive-%v-%03d.ts",
		"-master_pl_name", "adaptive.m3u8", outputDir+"adaptive-%v.m3u8",
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing FFmpeg:", err)
	}

	fmt.Println("Adaptive HLS files generated successfully in", outputDir)
}
