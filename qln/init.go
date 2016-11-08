package qln

import "github.com/boltdb/bolt"

func (nd *LnNode) Init(dbfilename string, basewal UWallet) error {

	err := nd.OpenDB(dbfilename)
	if err != nil {
		return err
	}
	nd.OmniChan = make(chan []byte, 10)
	go nd.OmniHandler()

	// connect to base wallet
	nd.BaseWallet = basewal
	// ask basewallet for outpoint event messages
	go nd.OPEventHandler()

	return nil
}

// Opens the DB file for the LnNode
func (nd *LnNode) OpenDB(filename string) error {
	var err error

	nd.LnDB, err = bolt.Open(filename, 0644, nil)
	if err != nil {
		return err
	}
	// create buckets if they're not already there
	err = nd.LnDB.Update(func(btx *bolt.Tx) error {
		_, err := btx.CreateBucketIfNotExists(BKTPeers)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
