package model

import (
	"sync"
)

var (
	pls     []Peers
	plsLock sync.Mutex
)

// InitPeerLists For test only
func InitPeerLists() {
	pls = append(pls, Pl0)
	pls = append(pls, Pl1)
	//fmt.Print(pls)
}

func SearchPeerList(p Peer) Peers {
	searchTarget := p.InfoHash
	re := Peers{InfoHash: "noMatch"}
	for _, pl := range pls {
		//fmt.Print(" [" + pl.InfoHash + " " + searchTarget + "]")
		//fmt.Println(pl.Peers)
		if pl.InfoHash == searchTarget {
			re = pl
		}
		//fmt.Println(pl.Peers)
	}
	return re
}

func AddPeerList(p Peer) {
	pl := Peers{Peers: []Peer{p}, InfoHash: p.InfoHash}
	pls = append(pls, pl)
}

func AddPeer(p Peer) {
	for i, pl := range pls {
		if pl.InfoHash == p.InfoHash {
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
	}
}

// LockPeerLists Lock
func LockPeerLists() {
	plsLock.Lock()
}

// UnlockPeerLists Unlock
func UnlockPeerLists() {
	plsLock.Unlock()
}
