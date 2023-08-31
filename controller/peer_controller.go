package controller

import (
	"Go2Tracker/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Return index
func handlePeerAnnounce(c *gin.Context) {
	p := model.Peer{}

	p.InfoHash = c.Query("info_hash")
	p.PeerID = c.Query("peer_id")
	p.IP = c.Query("ip")
	p.Port, _ = strconv.Atoi(c.Query("port"))
	p.Uploaded, _ = strconv.ParseInt(c.Query("uploaded"), 10, 64)
	p.Downloaded, _ = strconv.ParseInt(c.Query("downloaded"), 10, 64)
	p.Event = c.Query("event")

	fmt.Printf("info_hash: %s peer_id: %s; ip: %s;", p.InfoHash, p.PeerID, p.IP)
}

// RouteAnnounce Handle
func RouteAnnounce(r *gin.Engine) {
	r.GET("/announce", handlePeerAnnounce)
}
