type Room struct{
	Peers *Peers
	Hub *chat.Hub
}

type Peers struct{
	ListLock sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type peerConnectionState struct{
	PeerConnection *webrtc.PeerConnection
	websocket *ThreadSafeWriter
}

type ThreadSafeWriter struct{
	coon *websocket.Conn
	Mutex sync.Mutex
}

func (t *ThreadSafeWriter) WriteJSON(V interface()) error{
	t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

func (p *Peers) AddTrack(t *webrts.TrackRemote) *webrtc.TrackLocalStaticRTP{
	p.ListLock.Lock()
	p.SignalPeerConnections()
}()

trackLocal, errr := webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodecCapability, t.ID().StreamID())

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	p.TrackLocals[t.ID()] = trackLocal
	return trackLocal

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP){
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.SignalPeerConnections()
	}()
	delete(p.TrackLocals, t.ID())
}

func  (p *Peers)SignalPeerConnections(){
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.dispatchKeyFrames()
	}()

	attemptSync := func() (tryAgain bool){
		for i := range p.Connections{
			if p.Connections[i].PeerConnection.ConnectionState() == webrtc.PeerConnectionStateClosed{
				p.Connections = append(p.Connections[:i],np.Connections[i+1:]...)
				logPrintln("a".p.Connections)
				return true
			}

			existingSenders : map[string]bool{}
			for _,sender := range p.Connections[i].PeerConnection.GetSenders(){
				if sender.Track() == nil {
					continue
				}
				
				existingSenders[senders.Track().ID()] = true

				if _, ok := p.TrackLocals[sender.Track().ID()]; !ok {
					if err := p.Connections[i].PeerConnection.RemoveTrack(sender); err != nil{
						return true
					}
				}
			}

			for _, receiver := range p.Connections[i].PeerConnection.GetReceivers(){
				if receiver.Track() == nil{
					continue
				}
				existingSenders[receiver.Track().ID()] = true
			}

			for trackID := range p.TrackLocals{
				if _,ok := existingSenders[traclID]; !ok{
					if _, err := p.Connetcions[i].PeerConnection.AddTrack(p.TrackLocals[trackID]; err != nil){
						return true
					}
				}
			}
		}
	}
}

func(p *Peers) dispatchKeyFrames(){

}

type WesocketMessage struct{
	Event string `json:"Event`
	Data string `json:"data"`
}