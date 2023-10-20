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

func SearchPeerList(p Peer) (int, int, Peers) {
	targetInfoHash := p.InfoHash
	targetPeerID := p.PeerID

	var x, y int
	var re Peers

	for i, pl := range pls {
		if pl.InfoHash == targetInfoHash {
			x = i
			re = pl
			// do not return request itself
			for j, pEX := range re.Peers {
				if pEX.PeerID == targetPeerID {
					y = j
					copy(re.Peers[j:], re.Peers[j+1:])
					re.Peers = re.Peers[:len(re.Peers)-1]
					break
				}
			}
			break
		}
	}

	return x, y, re
}

func AddPeerList(p Peer) {
	pl := Peers{Peers: []Peer{p}, InfoHash: p.InfoHash}
	pls = append(pls, pl)
}

func UpdatePeers(p Peer) {
	for i, pl := range pls {
		if pl.InfoHash == p.InfoHash {
			flag := true
			for j, pEX := range pls[i].Peers {
				if p.PeerID == pEX.PeerID {
					pls[i].Peers[j] = p
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
