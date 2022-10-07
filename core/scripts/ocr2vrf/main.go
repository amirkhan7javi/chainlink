package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"

	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/arbitrum_vrf_beacon"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

type commonSetConfigArgs struct {
	onchainPubKeys         string
	offchainPubKeys        string
	configPubKeys          string
	peerIDs                string
	transmitters           string
	schedule               string
	f                      uint
	deltaProgress          time.Duration
	deltaResend            time.Duration
	deltaRound             time.Duration
	deltaGrace             time.Duration
	deltaStage             time.Duration
	maxRounds              uint8
	maxDurationQuery       time.Duration
	maxDurationObservation time.Duration
	maxDurationReport      time.Duration
	maxDurationAccept      time.Duration
	maxDurationTransmit    time.Duration
}

type dkgSetConfigArgs struct {
	commonSetConfigArgs
	dkgEncryptionPubKeys string
	dkgSigningPubKeys    string
	keyID                string
}

type vrfBeaconSetConfigArgs struct {
	commonSetConfigArgs
	confDelays string
}

func main() {
	e := helpers.SetupEnv(false)

	switch os.Args[1] {

	case "dkg-deploy":
		deployDKG(e)

	case "coordinator-deploy":
		cmd := flag.NewFlagSet("coordinator-deploy", flag.ExitOnError)
		beaconPeriodBlocks := cmd.Int64("beacon-period-blocks", 1, "beacon period in number of blocks")
		linkAddress := cmd.String("link-address", "", "link contract address")
		helpers.ParseArgs(cmd, os.Args[2:], "beacon-period-blocks", "link-address")
		deployVRFCoordinator(e, big.NewInt(*beaconPeriodBlocks), *linkAddress)

	case "beacon-deploy":
		cmd := flag.NewFlagSet("beacon-deploy", flag.ExitOnError)
		coordinatorAddress := cmd.String("coordinator-address", "", "coordinator contract address")
		linkAddress := cmd.String("link-address", "", "link contract address")
		dkgAddress := cmd.String("dkg-address", "", "dkg contract address")
		keyID := cmd.String("key-id", "", "key ID")
		helpers.ParseArgs(cmd, os.Args[2:], "beacon-deploy", "coordinator-address", "link-address", "dkg-address", "key-id")
		deployVRFBeacon(e, *coordinatorAddress, *linkAddress, *dkgAddress, *keyID)

	case "dkg-add-client":
		cmd := flag.NewFlagSet("dkg-add-client", flag.ExitOnError)
		dkgAddress := cmd.String("dkg-address", "", "DKG contract address")
		keyID := cmd.String("key-id", "", "key ID")
		clientAddress := cmd.String("client-address", "", "client address")
		helpers.ParseArgs(cmd, os.Args[2:], "dkg-address", "key-id", "client-address")
		addClientToDKG(e, *dkgAddress, *keyID, *clientAddress)

	case "dkg-remove-client":
		cmd := flag.NewFlagSet("dkg-add-client", flag.ExitOnError)
		dkgAddress := cmd.String("dkg-address", "", "DKG contract address")
		keyID := cmd.String("key-id", "", "key ID")
		clientAddress := cmd.String("client-address", "", "client address")
		helpers.ParseArgs(cmd, os.Args[2:], "dkg-address", "key-id", "client-address")
		removeClientFromDKG(e, *dkgAddress, *keyID, *clientAddress)

	case "dkg-set-config":
		cmd := flag.NewFlagSet("dkg-set-config", flag.ExitOnError)
		dkgAddress := cmd.String("dkg-address", "", "DKG contract address")
		keyID := cmd.String("key-id", "", "key ID")
		onchainPubKeys := cmd.String("onchain-pub-keys", "", "comma-separated list of OCR on-chain pubkeys")
		offchainPubKeys := cmd.String("offchain-pub-keys", "", "comma-separated list of OCR off-chain pubkeys")
		configPubKeys := cmd.String("config-pub-keys", "", "comma-separated list of OCR config pubkeys")
		peerIDs := cmd.String("peer-ids", "", "comma-separated list of peer IDs")
		transmitters := cmd.String("transmitters", "", "comma-separated list transmitters")
		dkgEncryptionPubKeys := cmd.String("dkg-encryption-pub-keys", "", "comma-separated list of DKG encryption pubkeys")
		dkgSigningPubKeys := cmd.String("dkg-signing-pub-keys", "", "comma-separated list of DKG signing pubkeys")
		schedule := cmd.String("schedule", "", "comma-separted list of transmission schedule")
		f := cmd.Uint("f", 1, "number of faulty oracles")
		deltaProgress := cmd.Duration("delta-progress", 30*time.Second, "duration of delta progress")
		deltaResend := cmd.Duration("delta-resend", 10*time.Second, "duration of delta resend")
		deltaRound := cmd.Duration("delta-round", 10*time.Second, "duration of delta round")
		deltaGrace := cmd.Duration("delta-grace", 20*time.Second, "duration of delta grace")
		deltaStage := cmd.Duration("delta-stage", 20*time.Second, "duration of delta stage")
		maxRounds := cmd.Uint("max-rounds", 3, "maximum number of rounds")
		maxDurationQuery := cmd.Duration("max-duration-query", 10*time.Millisecond, "maximum duration of query")
		maxDurationObservation := cmd.Duration("max-duration-observation", 10*time.Second, "maximum duration of observation method")
		maxDurationReport := cmd.Duration("max-duration-report", 10*time.Second, "maximum duration of report method")
		maxDurationAccept := cmd.Duration("max-duration-accept", 10*time.Millisecond, "maximum duration of shouldAcceptFinalizedReport method")
		maxDurationTransmit := cmd.Duration("max-duration-transmit", 1*time.Second, "maximum duration of shouldTransmitAcceptedReport method")

		helpers.ParseArgs(cmd,
			os.Args[2:],
			"dkg-address",
			"key-id",
			"onchain-pub-keys",
			"offchain-pub-keys",
			"config-pub-keys",
			"peer-ids",
			"transmitters",
			"dkg-encryption-pub-keys",
			"dkg-signing-pub-keys",
			"schedule")

		commands := dkgSetConfigArgs{
			commonSetConfigArgs: commonSetConfigArgs{
				onchainPubKeys:         *onchainPubKeys,
				offchainPubKeys:        *offchainPubKeys,
				configPubKeys:          *configPubKeys,
				peerIDs:                *peerIDs,
				transmitters:           *transmitters,
				schedule:               *schedule,
				f:                      *f,
				deltaProgress:          *deltaProgress,
				deltaResend:            *deltaResend,
				deltaRound:             *deltaRound,
				deltaGrace:             *deltaGrace,
				deltaStage:             *deltaStage,
				maxRounds:              uint8(*maxRounds),
				maxDurationQuery:       *maxDurationQuery,
				maxDurationObservation: *maxDurationObservation,
				maxDurationReport:      *maxDurationReport,
				maxDurationAccept:      *maxDurationAccept,
				maxDurationTransmit:    *maxDurationTransmit,
			},
			dkgEncryptionPubKeys: *dkgEncryptionPubKeys,
			dkgSigningPubKeys:    *dkgSigningPubKeys,
			keyID:                *keyID,
		}

		setDKGConfig(e, *dkgAddress, commands)

	case "beacon-set-config":
		cmd := flag.NewFlagSet("beacon-set-config", flag.ExitOnError)
		beaconAddress := cmd.String("beacon-address", "", "VRF beacon contract address")
		confDelays := cmd.String("conf-delays", "1,2,3,4,5,6,7,8", "comma-separted list of 8 confirmation delays")
		onchainPubKeys := cmd.String("onchain-pub-keys", "", "comma-separated list of OCR on-chain pubkeys")
		offchainPubKeys := cmd.String("offchain-pub-keys", "", "comma-separated list of OCR off-chain pubkeys")
		configPubKeys := cmd.String("config-pub-keys", "", "comma-separated list of OCR config pubkeys")
		peerIDs := cmd.String("peer-ids", "", "comma-separated list of peer IDs")
		transmitters := cmd.String("transmitters", "", "comma-separated list transmitters")
		schedule := cmd.String("schedule", "", "comma-separted list of transmission schedule")
		f := cmd.Uint("f", 1, "number of faulty oracles")
		// TODO: Adjust default delta* and maxDuration* values below after benchmarking latency
		deltaProgress := cmd.Duration("delta-progress", 30*time.Second, "duration of delta progress")
		deltaResend := cmd.Duration("delta-resend", 10*time.Second, "duration of delta resend")
		deltaRound := cmd.Duration("delta-round", 10*time.Second, "duration of delta round")
		deltaGrace := cmd.Duration("delta-grace", 20*time.Second, "duration of delta grace")
		deltaStage := cmd.Duration("delta-stage", 20*time.Second, "duration of delta stage")
		maxRounds := cmd.Uint("max-rounds", 3, "maximum number of rounds")
		maxDurationQuery := cmd.Duration("max-duration-query", 10*time.Millisecond, "maximum duration of query")
		maxDurationObservation := cmd.Duration("max-duration-observation", 10*time.Second, "maximum duration of observation method")
		maxDurationReport := cmd.Duration("max-duration-report", 10*time.Second, "maximum duration of report method")
		maxDurationAccept := cmd.Duration("max-duration-accept", 10*time.Millisecond, "maximum duration of shouldAcceptFinalizedReport method")
		maxDurationTransmit := cmd.Duration("max-duration-transmit", 1*time.Second, "maximum duration of shouldTransmitAcceptedReport method")

		helpers.ParseArgs(cmd,
			os.Args[2:],
			"beacon-address",
			"onchain-pub-keys",
			"offchain-pub-keys",
			"config-pub-keys",
			"peer-ids",
			"transmitters",
			"schedule")

		commands := vrfBeaconSetConfigArgs{
			commonSetConfigArgs: commonSetConfigArgs{
				onchainPubKeys:         *onchainPubKeys,
				offchainPubKeys:        *offchainPubKeys,
				configPubKeys:          *configPubKeys,
				peerIDs:                *peerIDs,
				transmitters:           *transmitters,
				schedule:               *schedule,
				f:                      *f,
				deltaProgress:          *deltaProgress,
				deltaResend:            *deltaResend,
				deltaRound:             *deltaRound,
				deltaGrace:             *deltaGrace,
				deltaStage:             *deltaStage,
				maxRounds:              uint8(*maxRounds),
				maxDurationQuery:       *maxDurationQuery,
				maxDurationObservation: *maxDurationObservation,
				maxDurationReport:      *maxDurationReport,
				maxDurationAccept:      *maxDurationAccept,
				maxDurationTransmit:    *maxDurationTransmit,
			},
			confDelays: *confDelays,
		}

		setVRFBeaconConfig(e, *beaconAddress, commands)

	case "coordinator-set-producer":
		cmd := flag.NewFlagSet("coordinator-set-producer", flag.ExitOnError)
		coordinatorAddress := cmd.String("coordinator-address", "", "VRF coordinator contract address")
		beaconAddress := cmd.String("beacon-address", "", "VRF beacon contract address")
		helpers.ParseArgs(cmd, os.Args[2:], "coordinator-address", "beacon-address")
		setProducer(e, *coordinatorAddress, *beaconAddress)

	case "coordinator-request-randomness":
		cmd := flag.NewFlagSet("coordinator-request-randomness", flag.ExitOnError)
		coordinatorAddress := cmd.String("coordinator-address", "", "VRF coordinator contract address")
		numWords := cmd.Uint("num-words", 1, "number of words to request")
		subID := cmd.Uint64("sub-id", 0, "subscription ID")
		confDelay := cmd.Int64("conf-delay", 1, "confirmation delay")
		helpers.ParseArgs(cmd, os.Args[2:], "coordinator-address", "sub-id")
		requestRandomness(e, *coordinatorAddress, uint16(*numWords), *subID, big.NewInt(*confDelay))

	case "coordinator-redeem-randomness":
		cmd := flag.NewFlagSet("coordinator-redeem-randomness", flag.ExitOnError)
		coordinatorAddress := cmd.String("coordinator-address", "", "VRF coordinator contract address")
		requestID := cmd.Int64("request-id", 0, "request ID")
		helpers.ParseArgs(cmd, os.Args[2:], "coordinator-address", "request-id")
		redeemRandomness(e, *coordinatorAddress, big.NewInt(*requestID))

	case "beacon-info":
		cmd := flag.NewFlagSet("coordinator-info", flag.ExitOnError)
		beaconAddress := cmd.String("beacon-address", "", "VRF coordinator contract address")
		helpers.ParseArgs(cmd, os.Args[2:], "beacon-address")
		beacon := newVRFBeacon(e, common.HexToAddress(*beaconAddress))
		keyID, err := beacon.SKeyID(nil)
		helpers.PanicErr(err)
		fmt.Println("coordinator key id:", hexutil.Encode(keyID[:]))
		keyHash, err := beacon.SProvingKeyHash(nil)
		helpers.PanicErr(err)
		fmt.Println("coordinator proving key hash:", hexutil.Encode(keyHash[:]))

	case "consumer-deploy":
		cmd := flag.NewFlagSet("consumer-deploy", flag.ExitOnError)
		coordinatorAddress := cmd.String("coordinator-address", "", "VRF coordinator address")
		shouldFail := cmd.Bool("should-fail", false, "shouldFail flag")
		beaconPeriodBlocks := cmd.Int64("beacon-period-blocks", 1, "beacon period in number of blocks")
		helpers.ParseArgs(cmd, os.Args[2:], "coordinator-address", "beacon-period-blocks")
		deployVRFBeaconCoordinatorConsumer(e, *coordinatorAddress, *shouldFail, big.NewInt(*beaconPeriodBlocks))

	case "consumer-request-randomness":
		cmd := flag.NewFlagSet("consumer-request-randomness", flag.ExitOnError)
		consumerAddress := cmd.String("consumer-address", "", "VRF coordinator consumer address")
		numWords := cmd.Uint("num-words", 1, "number of words to request")
		subID := cmd.Uint64("sub-id", 0, "subscription ID")
		confDelay := cmd.Int64("conf-delay", 1, "confirmation delay")
		helpers.ParseArgs(cmd, os.Args[2:], "consumer-address", "sub-id")
		requestRandomnessFromConsumer(e, *consumerAddress, uint16(*numWords), *subID, big.NewInt(*confDelay))

	case "consumer-redeem-randomness":
		cmd := flag.NewFlagSet("consumer-redeem-randomness", flag.ExitOnError)
		consumerAddress := cmd.String("consumer-address", "", "VRF coordinator consumer address")
		requestID := cmd.Int64("request-id", 0, "request ID")
		numWords := cmd.Int64("num-words", 1, "number of words to print after redeeming")
		helpers.ParseArgs(cmd, os.Args[2:], "consumer-address", "request-id")
		redeemRandomnessFromConsumer(e, *consumerAddress, big.NewInt(*requestID), *numWords)

	case "consumer-request-callback":
		cmd := flag.NewFlagSet("consumer-request-callback", flag.ExitOnError)
		consumerAddress := cmd.String("consumer-address", "", "VRF coordinator consumer address")
		numWords := cmd.Uint("num-words", 1, "number of words to request")
		subID := cmd.Uint64("sub-id", 0, "subscription ID")
		confDelay := cmd.Int64("conf-delay", 1, "confirmation delay")
		callbackGasLimit := cmd.Uint("cb-gas-limit", 50_000, "callback gas limit")
		helpers.ParseArgs(cmd, os.Args[2:], "consumer-address")
		requestRandomnessCallback(
			e,
			*consumerAddress,
			uint16(*numWords),
			*subID,
			big.NewInt(int64(*confDelay)),
			uint32(*callbackGasLimit),
			nil, // test consumer doesn't use any args
		)

	case "consumer-read-randomness":
		cmd := flag.NewFlagSet("consumer-read-randomness", flag.ExitOnError)
		consumerAddress := cmd.String("consumer-address", "", "VRF coordinator consumer address")
		requestID := cmd.String("request-id", "", "VRF request ID")
		numWords := cmd.Int("num-words", 1, "number of words to fetch")
		helpers.ParseArgs(cmd, os.Args[2:], "consumer-address")
		consumer := newBeaconConsumer(e, common.HexToAddress(*consumerAddress))
		printRandomnessFromConsumer(consumer, decimal.RequireFromString(*requestID).BigInt(), int64(*numWords))

	case "consumer-request-callback-batch":
		cmd := flag.NewFlagSet("consumer-request-callback", flag.ExitOnError)
		consumerAddress := cmd.String("consumer-address", "", "VRF beacon consumer address")
		numWords := cmd.Uint("num-words", 1, "number of words to request")
		subID := cmd.Uint64("sub-id", 0, "subscription ID")
		confDelay := cmd.Int64("conf-delay", 1, "confirmation delay")
		batchSize := cmd.Int64("batch-size", 1, "batch size")
		callbackGasLimit := cmd.Uint("cb-gas-limit", 200_000, "callback gas limit")
		helpers.ParseArgs(cmd, os.Args[2:], "consumer-address")
		requestRandomnessCallbackBatch(
			e,
			*consumerAddress,
			uint16(*numWords),
			*subID,
			big.NewInt(int64(*confDelay)),
			uint32(*callbackGasLimit),
			nil, // test consumer doesn't use any args,
			big.NewInt(*batchSize),
		)
	case "dkg-setup":
		setupDKGNodes(e)
	case "ocr2vrf-setup":
		setupOCR2VRFNodes(e)
	case "decode-error":
		cmd := flag.NewFlagSet("decode-error", flag.ExitOnError)
		errorHex := cmd.String("error-hex", "", "custom error in hex")
		helpers.ParseArgs(cmd, os.Args[2:], "error-hex")
		if !strings.HasPrefix(*errorHex, "0x") {
			*errorHex = "0x" + *errorHex
		}
		errorBytes := hexutil.MustDecode(*errorHex)
		tabi := evmtypes.MustGetABI(arbitrum_vrf_beacon.ArbitrumVRFBeaconMetaData.ABI)
		var te *abi.Error
		for _, e := range tabi.Errors {
			if e.ID.Hex()[2:10] == (*errorHex)[2:10] {
				te = &e
				break
			}
		}
		if te == nil {
			fmt.Println("Can't find corresponding error - wrong contract specified?")
			os.Exit(1)
		} else {
			fmt.Println("Error found:", te.Sig)
		}

		data, err := te.Unpack(errorBytes)
		helpers.PanicErr(err)
		fmt.Printf("error data unpacked: %+v\n", data)
	default:
		panic("unrecognized subcommand: " + os.Args[1])
	}
}
