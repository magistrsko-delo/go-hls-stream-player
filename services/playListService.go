package services

import (
	"go-hls-stream-player/grpc_client"
	pbTimeshift "go-hls-stream-player/proto/timeshift_service"
	"strconv"
	"strings"
)

type PlaylistService struct {
	timeShiftClient *grpc_client.TimeShiftClient
}

func (playlistService *PlaylistService) GenerateMediaPlaylist(mediaId int32) (string, error)  {
	stream := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:3",
		"#EXT-X-MEDIA-SEQUENCE:0",
		"#EXT-X-TARGETDURATION:6",
	}

	vods, err := playlistService.get1080pMediaStream(mediaId)

	if err != nil {
		return "", err
	}

	stream = append(stream, vods...)
	stream = append(stream, "#EXT-X-ENDLIST")

	return strings.Join(stream, "\n"), nil
}

func (playlistService *PlaylistService) get1080pMediaStream(mediaId int32) ([]string, error)  {

	rsp, err := playlistService.timeShiftClient.GetMediaChunkInfo(mediaId)

	if err != nil {
		return nil, err
	}

	mediaData := rsp.GetData()
	chunksData := [] *pbTimeshift.ChunkResponse{}
	for i := 0; i < len(mediaData); i++ {
		if mediaData[i].GetResolution() == "1920x1080" {
			chunksData = mediaData[i].GetChunks()
			break
		}
	}

	vod1800p := [] string{}

	for i := 0; i < len(chunksData); i++ {
		vod1800p = append(vod1800p, "#EXTINF:" + strconv.FormatFloat(chunksData[i].GetLength(), 'f', 6, 64)  + ",")
		vod1800p = append(vod1800p, chunksData[i].GetChunksUrl())
		vod1800p = append(vod1800p, "EXT-X-DISCONTINUITY")
	}

	return vod1800p, nil
}

func InitPlaylistService()  *PlaylistService  {
	return &PlaylistService{timeShiftClient:grpc_client.InitTimeShiftClient()}
}
