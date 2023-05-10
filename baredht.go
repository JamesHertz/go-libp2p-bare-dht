package baredht

// the fun starts now :)
import (
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	kb "github.com/libp2p/go-libp2p-kbucket"
)

type BareDHT struct {
	host      host.Host // the network services we need
	self      peer.ID   // Local peer (yourself)
	selfKey   kb.ID
	// rtRefreshManager *rtrefresh.RtRefreshManager

	bucketSize int
	alpha      int // The concurrency parameter per path
	beta       int // The number of peers closest to a target that must have responded for a query path to terminate

	autoRefresh bool

	// ...
	bootstrapPeers func() []peer.AddrInfo

	maxRecordAge time.Duration

	// ...
	// configuration variables for tests
	testAddressUpdateProcessing bool
} 

func NewBareDHT() *BareDHT {
	return nil
}
