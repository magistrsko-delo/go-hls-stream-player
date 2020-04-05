package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type PlaylistController struct {
	
}

func (playlistController *PlaylistController) GetPlaylistMedia(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	fmt.Println(params)

	stream := "#EXTM3U\n" +
		"#EXT-X-VERSION:3\n" +
		"#EXT-X-MEDIA-SEQUENCE:0\n" +
		"#EXT-X-TARGETDURATION:11\n" +
		"#EXT-X-KEY:METHOD=AES-128,URI=\"https://vod.api.24ur.si/vodKey/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJwb3AiLCJjb250ZXh0Ijp7ImJhY2tlbmQiOjAsImRldmljZV9mYW1pbHkiOiJkZXNrdG9wIiwiZGV2aWNlX2lkIjoiYTBkN2RlMDMtNzNhZC00MTQ0LTgzNzYtMmQzNDM5ZTU3NTdiIiwiZHJtX3Byb3RlY3RlZCI6MCwiZW5kX2NodW5rIjowLCJleHBpcmVzIjowLCJpc19sb2NhbF9jaXRpemVuIjp0cnVlLCJtZWRpYV9maWxlbmFtZSI6ImE5NGRmM2JiZGJfNjI0MDU1MDUiLCJtZWRpYV9nZW9sb2NrIjoiIiwibWVkaWFfcHVibGlzaGVkX2Zyb20iOjE1ODYwODE5NDAsIm1lZGlhX3B1Ymxpc2hlZF90byI6MCwic2VjdGlvbl9pZCI6Miwic2l0ZV9pZCI6MSwic2tpcF9nZW9sb2NrIjowLCJzdGFydF9jaHVuayI6MCwidmlzaXRvcl9pZCI6NjU4MTIxLCJ2aXNpdG9yX2lwIjoiMTkzLjc3LjgzLjIxNSJ9LCJleHAiOjE1ODYxNjkzMjksImlhdCI6MTU4NjA4MjkyOSwiaXNzIjoicG9wIn0.GBF6mHjuGd2e662_hrDdOkxDAWLq4s4Z1mYMn5USbDo/a94df3bbdb_62405505/high.m3u8/a94df3bbdb_62405505_1.key\",IV=0x1feac82aa0bde21fe419916c5fa96287\n" +
		"#EXTINF:10.080,\n" +
		"https://redirect.api.24ur.si/v2/58da63acb52ef93e4616a233d5b8bd28dc856b2e/a0d7de03-73ad-4144-8376-2d3439e5757b/658121/62405505/0/0/1/2/10/desktop/https%3A%2F%2Fhlsnvod.24ur.com%2F2020%2F04%2F05%2Fa94df3bbdb_62405505%2Fclip_18.ts\n" +
		"#EXTINF:10.000,\n" +
		"https://redirect.api.24ur.si/v2/58da63acb52ef93e4616a233d5b8bd28dc856b2e/a0d7de03-73ad-4144-8376-2d3439e5757b/658121/62405505/5/0/1/2/20/desktop/https%3A%2F%2Fhlsnvod.24ur.com%2F2020%2F04%2F05%2Fa94df3bbdb_62405505%2Fclip_19.ts\n" +
		"#EXT-X-ENDLIST\n"

	w.Header().Set("Content-Type", "application/x-mpegURL")
	w.Write([]byte(stream))
}
