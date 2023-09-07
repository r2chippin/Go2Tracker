package controller

import (
	"Go2Tracker/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strconv"
)

func errorSensor(p model.Peer) (string, bool) {
	resp := ""
	res := false

	ip := net.ParseIP(p.IP)
	if ip == nil {
		resp = "IP is nil"
	}
	if ip.To4() == nil {
		resp = "IP is not a IPv4 address"
	}
	if !ip.IsLoopback() && !ip.IsMulticast() && !ip.IsPrivate() {
		resp = "IP is a unusable IPv4 address"
	}

	if resp != "" {
		res = true
	}
	return resp, res
}

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
		model.AddPeer(p)
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
