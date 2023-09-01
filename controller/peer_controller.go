package controller

import (
	"Go2Tracker/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Handle with new peer
func handlePeerAnnounce(c *gin.Context) {
	p := model.Peer{}

	p.InfoHash = c.Query("info_hash")
	p.PeerID = c.Query("peer_id")
	p.IP = c.Query("ip")
	p.Port, _ = strconv.Atoi(c.Query("port"))
	p.Uploaded, _ = strconv.ParseInt(c.Query("uploaded"), 10, 64)
	p.Downloaded, _ = strconv.ParseInt(c.Query("downloaded"), 10, 64)
	p.Event = c.Query("event")

	fmt.Println(p)

	interval := 720
	model.LockPeerLists()

	pl := model.SearchPeerList(p)
	if pl.InfoHash == "noMatch" {
		model.AddPeerList(p)
	} else {

	}
	model.UnlockPeerLists()
	pld := model.ConvertPeersToDictList(pl)

	response := gin.H{}
	response["interval"] = interval
	response["peers"] = pld

	c.JSON(http.StatusOK, response)
}

// RouteAnnounce Handle
func RouteAnnounce(e *gin.Engine) {
	model.InitPeerLists()
	e.GET("/announce", handlePeerAnnounce)
}
