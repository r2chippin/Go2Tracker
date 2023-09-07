package controller

import (
	"Go2Tracker/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strconv"
)

func errorSensor(p model.Peer) (bool, string) {
	resp := ""
	res := false

	// detect IP error
	ip := net.ParseIP(p.IP)
	if ip == nil {
		resp += "IP is nil"
	}
	if ip.To4() == nil {
		resp += "IP is not a IPv4 address"
	}
	if ip.IsLoopback() || ip.IsMulticast() || ip.IsPrivate() {
		resp += "IP is a unusable IPv4 address"
	}

	// detect PORT error
	port := p.Port
	if port <= 0 || port >= 65535 {
		resp += "Port out of range"
	}

	if resp != "" {
		res = true
	}
	return res, resp
}

// Handle with new peer
func handlePeerAnnounce(c *gin.Context) {
	response := gin.H{}

	p := model.Peer{}

	p.InfoHash = c.Query("info_hash")
	p.PeerID = c.Query("peer_id")
	p.IP = c.Query("ip")
	p.Port, _ = strconv.Atoi(c.Query("port"))
	p.Uploaded, _ = strconv.ParseInt(c.Query("uploaded"), 10, 64)
	p.Downloaded, _ = strconv.ParseInt(c.Query("downloaded"), 10, 64)
	p.Event = c.Query("event")

	fmt.Println(p)

	errT, errP := errorSensor(p)
	if errT {
		println(errP)
		response["error"] = errP
		c.JSON(http.StatusBadRequest, response)
		return
	}

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

	response["interval"] = interval
	response["peers"] = pld

	c.JSON(http.StatusOK, response)
}

// RouteAnnounce Handle
func RouteAnnounce(e *gin.Engine) {
	model.InitPeerLists()
	e.GET("/announce", handlePeerAnnounce)
}
