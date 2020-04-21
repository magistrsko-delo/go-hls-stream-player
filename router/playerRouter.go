package router

import (
	"github.com/gorilla/mux"
	"go-hls-stream-player/controllers"
	"go-hls-stream-player/services"
)

type PlaylistRouter struct {
	Router *mux.Router
}

func (playlistRouter *PlaylistRouter) RegisterHandlers()  {
	controller :=  &controllers.PlaylistController{PlayListService: services.InitPlaylistService()}
	(*playlistRouter).Router.HandleFunc("/vod/{mediaId}/master.m3u8", controller.GetMasterPlaylist).Methods("GET")
	(*playlistRouter).Router.HandleFunc("/vod/{mediaId}/1080p.m3u8", controller.Get1080pPlaylistMedia).Methods("GET")
	(*playlistRouter).Router.HandleFunc("/vod/{mediaId}/480p.m3u8", controller.Get480pPlaylistMedia).Methods("GET")

	(*playlistRouter).Router.HandleFunc("/vod/sequence/{sequenceId}/1080p.m3u8", controller.GetSequencePlaylist).Methods("GET")
}
