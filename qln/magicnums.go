package qln

import "github.com/btcsuite/btcutil/hdkeychain"

const (
	// high 3 bytes are in sequence, low 3 bytes are in time
	seqMask  = 0xff000000 // assert high byte
	timeMask = 0x21000000 // 1987 to 1988

	MSGID_POINTREQ  = 0x30
	MSGID_POINTRESP = 0x31
	MSGID_CHANDESC  = 0x32
	MSGID_CHANACK   = 0x33
	MSGID_SIGPROOF  = 0x34

	MSGID_CLOSEREQ  = 0x40 // close channel
	MSGID_CLOSERESP = 0x41

	MSGID_TEXTCHAT = 0x70 // send a text message

	MSGID_RTS    = 0x80 // pushing funds in channel; request to send
	MSGID_ACKSIG = 0x81 // pulling funds in channel; acknowledge update and sign
	MSGID_SIGREV = 0x82 // pushing funds; signing new state and revoking old
	MSGID_REVOKE = 0x83 // pulling funds; revoking previous channel state

	MSGID_FWDMSG     = 0x20
	MSGID_FWDAUTHREQ = 0x21

	MSGID_
)

const (
	UseWallet          = 0 | hdkeychain.HardenedKeyStart
	UseChannelFund     = 20 | hdkeychain.HardenedKeyStart
	UseChannelRefund   = 30 | hdkeychain.HardenedKeyStart
	UseChannelHAKDBase = 40 | hdkeychain.HardenedKeyStart
	UseChannelElkrem   = 8888 | hdkeychain.HardenedKeyStart
	// links Id and channel. replaces UseChannelFund
	UseChannelNonce = 100 | hdkeychain.HardenedKeyStart

	UseIdKey = 111 + hdkeychain.HardenedKeyStart
)

var (
	BKTPeers   = []byte("pir") // all peer data is in this bucket.
	KEYIdx     = []byte("idx") // index for key derivation
	KEYutxo    = []byte("utx") // serialized utxo for the channel
	KEYUnsig   = []byte("usg") // unsigned fund tx
	KEYCladr   = []byte("cdr") // coop close address (Don't make fun of my lisp)
	KEYState   = []byte("ima") // channel state
	KEYElkRecv = []byte("elk") // elkrem receiver
	KEYqclose  = []byte("qcl") // channel close outpoint & height
)
