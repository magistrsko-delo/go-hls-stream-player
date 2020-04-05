package router

import (
	"github.com/gorilla/mux"
	"go-hls-stream-player/controllers"
)

type PlaylistRouter struct {
	Router *mux.Router
}

func (playlistRouter *PlaylistRouter) RegisterHandlers()  {
	controller :=  &controllers.PlaylistController{}
	(*playlistRouter).Router.HandleFunc("/vod/{mediaId}/index.m3u8", controller.GetPlaylistMedia).Methods("GET")
}
