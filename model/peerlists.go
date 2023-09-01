package model

import (
	"fmt"
	"sync"
)

var (
	pls     []PeerList
	plsLock sync.Mutex
)

// InitPeerLists For test only
func InitPeerLists() {
	pls = append(pls, Pl0)
	pls = append(pls, Pl1)
	//fmt.Print(pls)
}

func SearchPeerList(p Peer) PeerList {
	searchTarget := p.InfoHash
	re := PeerList{InfoHash: "noMatch"}
	for i, pl := range pls {
		//fmt.Print(" [" + pl.InfoHash + " " + searchTarget + "]")
		fmt.Println(pl.Peers)
		if pl.InfoHash == searchTarget {
			re = pl
			flag := true
			for _, pEX := range pls[i].Peers {
				if pEX.PeerID == p.PeerID {
					flag = false
				}
			}
			if flag {
				pls[i].Peers = append(pls[i].Peers, p)
			}
		}
		fmt.Println(pl.Peers)
	}
	return re
}

func AddPeerList(p Peer) {
	pl := PeerList{Peers: []Peer{p}, InfoHash: p.InfoHash}
	pls = append(pls, pl)
}

/*
func AddPeer(p Peer) {

}
// TO DO
*/

// LockPeerLists Lock
func LockPeerLists() {
	plsLock.Lock()
}

// UnlockPeerLists Unlock
func UnlockPeerLists() {
	plsLock.Unlock()
}
