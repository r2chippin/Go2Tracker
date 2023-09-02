package model

var (
	Peer0 = Peer{
		InfoHash:   "hash0",
		PeerID:     "id0",
		IP:         "127.0.0.0",
		Port:       8080,
		Uploaded:   0,
		Downloaded: 0,
		Left:       99,
		Event:      "Downloading",
	}

	Peer1 = Peer{
		InfoHash:   "hash1",
		PeerID:     "id1",
		IP:         "127.0.0.1",
		Port:       8081,
		Uploaded:   0,
		Downloaded: 0,
		Left:       99,
		Event:      "Downloading",
	}
	Peer2 = Peer{
		InfoHash:   "hash1",
		PeerID:     "id2",
		IP:         "127.0.0.2",
		Port:       8082,
		Uploaded:   0,
		Downloaded: 0,
		Left:       99,
		Event:      "Downloading",
	}

	Pl0 = Peers{
		Peers:    []Peer{Peer0},
		InfoHash: "hash0",
	}

	Pl1 = Peers{
		Peers:    []Peer{Peer1, Peer2},
		InfoHash: "hash1",
	}
)

type Peer struct {
	InfoHash   string `json:"info_hash"`
	PeerID     string `json:"peer_id"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	Uploaded   int64  `json:"uploaded"`
	Downloaded int64  `json:"downloaded"`
	Left       int64  `json:"left"`
	Event      string `json:"event"`
}

type Peers struct {
	Peers    []Peer `json:"peers"`
	InfoHash string `json:"info_hash"`
}

func ConvertPeersToDictList(pl Peers) []map[string]interface{} {
	dictList := make([]map[string]interface{}, len(pl.Peers))

	for i, peer := range pl.Peers {
		dict := make(map[string]interface{})
		dict["peer_id"] = peer.PeerID
		dict["ip"] = peer.IP
		dict["port"] = peer.Port
		dictList[i] = dict
	}

	return dictList
}
