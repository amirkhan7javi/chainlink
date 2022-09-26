package directrequestocr

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type RequestState int8

const (
	REQ_IN_PROGRESS  = 1
	REQ_RESULT_READY = 2
	REQ_TRANSMITTED  = 3
	REQ_CONFIRMED    = 4
)

type ErrType int8

const (
	ERR_NONE            = 1
	ERR_NODE_EXCEPTION  = 2
	ERR_SANDBOX_TIMEOUT = 3
	ERR_USER_EXCEPTION  = 4
)

type Request struct {
	ID                int64
	ContractRequestID [32]byte
	RunID             int64
	ReceivedAt        time.Time
	RequestTxHash     *common.Hash
	State             RequestState
	ResultReadyAt     time.Time
	Result            []byte
	ErrorType         ErrType
	Error             string
	// True if this node submitted an observation for this request in any OCR rounds.
	IsOCRParticipant  bool
	TransmittedResult []byte
	TransmittedError  string
}
