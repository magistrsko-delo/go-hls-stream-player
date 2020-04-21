package services

import (
	"go-hls-stream-player/Models"
	"go-hls-stream-player/grpc_client"
	pbTimeshift "go-hls-stream-player/proto/timeshift_service"
	"strconv"
	"strings"
)

type PlaylistService struct {
	timeShiftClient *grpc_client.TimeShiftClient
	streamPlayListInit []string

}

func (playlistService *PlaylistService) GenerateMasterMediaPlaylist(mediaId int) (string, error) {
	stream := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:3",
	}

	stream = append(stream, "#EXT-X-STREAM-INF:BANDWIDTH=1400000,RESOLUTION=842x480")
	stream = append(stream, Models.GetEnvStruct().Url + "v1/vod/" + strconv.Itoa(mediaId) + "/480p.m3u8")

	stream = append(stream, "#EXT-X-STREAM-INF:BANDWIDTH=5000000,RESOLUTION=1920x1080")
	stream = append(stream, Models.GetEnvStruct().Url + "v1/vod/" + strconv.Itoa(mediaId) + "/1080p.m3u8")

	return strings.Join(stream, "\n"), nil
}


func (playlistService *PlaylistService) GenerateMediaPlaylist1080p(mediaId int32) (string, error)  {
	stream := playlistService.streamPlayListInit

	rsp, err := playlistService.timeShiftClient.GetMediaChunkInfo(mediaId)

	if err != nil {
		return "", err
	}
	chunksData := playlistService.getChunksResolutionData("1920x1080", rsp.GetData())
	vods := playlistService.generateChunksPlaylistDataFromMetadata(chunksData)

	stream = append(stream, vods...)
	stream = append(stream, "#EXT-X-ENDLIST")

	return strings.Join(stream, "\n"), nil
}

func (playlistService *PlaylistService) GenerateMediaPlaylist1480p(mediaId int32) (string, error)  {
	stream := playlistService.streamPlayListInit

	rsp, err := playlistService.timeShiftClient.GetMediaChunkInfo(mediaId)

	if err != nil {
		return "", err
	}
	chunksData := playlistService.getChunksResolutionData("842x480", rsp.GetData())
	vods := playlistService.generateChunksPlaylistDataFromMetadata(chunksData)

	stream = append(stream, vods...)
	stream = append(stream, "#EXT-X-ENDLIST")

	return strings.Join(stream, "\n"), nil
}

/// SEQUENCE PLAYLIST... ONLY ONE RESOLUTION...
func (playlistService *PlaylistService) GenerateSequencePlaylist(sequenceId int32) (string, error)  {
	stream := playlistService.streamPlayListInit

	rsp, err := playlistService.timeShiftClient.GetSequenceChunkInfo(sequenceId)
	if err != nil {
		return "", err
	}

	chunkData := playlistService.getChunksResolutionData("1920x1080", rsp.GetData())
	vods := playlistService.generateChunksPlaylistDataFromMetadata(chunkData)

	stream = append(stream, vods...)
	stream = append(stream, "#EXT-X-ENDLIST")

	return strings.Join(stream, "\n"), nil
}



// UTILITY FUNCTIONS

func (playlistService *PlaylistService) getChunksResolutionData(resolution string, data [] *pbTimeshift.ChunkResolutionResponse) [] *pbTimeshift.ChunkResponse  {
	mediaData := data
	chunksData := [] *pbTimeshift.ChunkResponse{}
	for i := 0; i < len(mediaData); i++ {
		if mediaData[i].GetResolution() == resolution {
			chunksData = mediaData[i].GetChunks()
			break
		}
	}

	return chunksData
}


func (playlistService *PlaylistService) generateChunksPlaylistDataFromMetadata(chunksData [] *pbTimeshift.ChunkResponse) []string {
	vod1800p := [] string{}

	for i := 0; i < len(chunksData); i++ {
		vod1800p = append(vod1800p, "#EXT-X-DISCONTINUITY")
		vod1800p = append(vod1800p, "#EXTINF:" + strconv.FormatFloat(chunksData[i].GetLength(), 'f', 6, 64)  + ",")
		vod1800p = append(vod1800p, chunksData[i].GetChunksUrl())
	}
	return vod1800p
}


func InitPlaylistService()  *PlaylistService  {
	return &PlaylistService{
		timeShiftClient:grpc_client.InitTimeShiftClient(),
		streamPlayListInit: []string{
			"#EXTM3U",
			"#EXT-X-VERSION:3",
			"#EXT-X-MEDIA-SEQUENCE:0",
			"#EXT-X-TARGETDURATION:5",
			"#EXT-X-PLAYLIST-TYPE:VOD",
		},
	}
}
