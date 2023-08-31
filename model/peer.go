package model

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

type PeerList struct {
	Peers []Peer `json:"peers"`
}

/*
type Peers struct {
}
*/

func ConvertPeersToDictList(pl PeerList) []map[string]interface{} {
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
