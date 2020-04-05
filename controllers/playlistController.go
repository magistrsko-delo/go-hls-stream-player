package controllers

import (
	"github.com/gorilla/mux"
	"go-hls-stream-player/services"
	"net/http"
	"strconv"
)

type PlaylistController struct {
	PlayListService *services.PlaylistService
}

func (playlistController *PlaylistController) GetMasterPlaylist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	mediaId, err := strconv.Atoi(params["mediaId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	masterPlaylist, err := playlistController.PlayListService.GenerateMasterMediaPlaylist(mediaId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/x-mpegURL")
	w.Write([]byte(masterPlaylist))
}

func (playlistController *PlaylistController) Get1080pPlaylistMedia(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	mediaId, err := strconv.Atoi(params["mediaId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist, err := playlistController.PlayListService.GenerateMediaPlaylist(int32(mediaId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-mpegURL")
	w.Write([]byte(playlist))
}
