package sources

import (
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type ESClient struct {
	logger      log.Logger
	nodeAddr    string
	batcherAddr common.Address
}

func NewESClient(nodeAddr string, batcherAddr common.Address, logger log.Logger) (*ESClient, error) {
	return &ESClient{logger: logger, nodeAddr: nodeAddr}, nil
}

func (s *ESClient) GetBlobs(blobHashes []eth.IndexedBlobHash) ([]*eth.BlobSidecar, error) {
	return nil, nil
}
