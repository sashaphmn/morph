package flags

import (
	"time"

	"github.com/urfave/cli"
)

const envVarPrefix = "TX_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	BuildEnvFlag = cli.StringFlag{
		Name: "build-env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   prefixEnvVar("BUILD_ENV"),
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar("L1_ETH_RPC"),
	}
	L2EthRpcsFlag = cli.StringSliceFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URLs for L2",
		Required: true,
		EnvVar:   prefixEnvVar("L2_ETH_RPCS"),
	}

	// l1 contract
	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the rollup contract",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}
	L1SequencerAddressFlag = cli.StringFlag{
		Name:     "l1-sequencer",
		Usage:    "Address of the sequencer contract",
		Required: true,
		EnvVar:   prefixEnvVar("L1_SEQUENCER_ADDRESS"),
	}

	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("POLL_INTERVAL"),
	}

	PrivateKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Usage:    "The private key to use for sending to the rollup contract",
		EnvVar:   prefixEnvVar("L1_PRIVATE_KEY"),
		Required: true,
	}

	TxTimeoutFlag = cli.DurationFlag{
		Name:     "tx-timeout",
		Usage:    "Timeout for transaction submission",
		Value:    10,
		EnvVar:   prefixEnvVar("TX_TIMEOUT"),
		Required: true,
	}

	// finalize flags
	FinalizeFlag = cli.BoolFlag{
		Name:     "finalize",
		Usage:    "Enable finalize",
		EnvVar:   prefixEnvVar("FINALIZE"),
		Required: true,
	}

	MaxFinalizeNumFlag = cli.Uint64Flag{
		Name:     "max-finalize-number",
		Usage:    "Maximum number of finalize",
		EnvVar:   prefixEnvVar("MAX_FINALIZE_NUM"),
		Required: true,
	}

	// decentralize config
	PriorityRollupFlag = cli.BoolFlag{
		Name:     "priority-rollup",
		Usage:    "Enable priority rollup",
		EnvVar:   prefixEnvVar("PRIORITY_ROLLUP"),
		Required: true,
	}

	// L2 contract
	SubmitterAddressFlag = cli.StringFlag{
		Name:     "submitter-address",
		Usage:    "Address of the submitter contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SUBMITTER_ADDRESS"),
	}
	L2SequencerAddressFlag = cli.StringFlag{
		Name:     "l2-sequencer",
		Usage:    "Address of the sequencer contract",
		Required: true,
		EnvVar:   prefixEnvVar("L2_SEQUENCER_ADDRESS"),
	}

	/* Optional Flags */
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	// metrics
	MetricsServerEnable = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Enable metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostname = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "Hostname at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPort = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "Port at which the metrics server is running",
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}

	// tx fee limit
	TxFeeLimitFlag = cli.Uint64Flag{
		Name:     "tx-fee-limit",
		Usage:    "The maximum fee for a transaction",
		Value:    5e17, //0.5eth
		EnvVar:   prefixEnvVar("TX_FEE_LIMIT"),
		Required: true,
	}

	// log to file
	LogFilename = cli.StringFlag{
		Name:   "log.filename",
		Usage:  "The target file for writing logs, backup log files will be retained in the same directory.",
		EnvVar: prefixEnvVar("LOG_FILENAME"),
	}
	LogFileMaxSize = cli.IntFlag{
		Name:   "log.maxsize",
		Usage:  "The maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes. It is used only when log.filename is provided.",
		Value:  100,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_SIZE"),
	}
	LogFileMaxAge = cli.IntFlag{
		Name:   "log.maxage",
		Usage:  "The maximum number of days to retain old log files based on the timestamp encoded in their filename. It defaults to 30 days. It is used only when log.filename is provided.",
		Value:  30,
		EnvVar: prefixEnvVar("LOG_FILE_MAX_AGE"),
	}
	LogCompress = cli.BoolFlag{
		Name:   "log.compress",
		Usage:  "Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression. It is used only when log.filename is provided.",
		EnvVar: prefixEnvVar("LOG_COMPRESS"),
	}

	// rollup interval
	RollupInterval = cli.DurationFlag{
		Name:   "ROLLUP_INTERVAL",
		Usage:  "Interval for rollup",
		Value:  500 * time.Millisecond,
		EnvVar: prefixEnvVar("ROLLUP_INTERVAL"),
	}
	// finalize interval
	FinalizeInterval = cli.DurationFlag{
		Name:   "FINALIZE_INTERVAL",
		Usage:  "Interval for finalize",
		Value:  2 * time.Second,
		EnvVar: prefixEnvVar("FINALIZE_INTERVAL"),
	}
	// tx process interval
	TxProcessInterval = cli.DurationFlag{
		Name:   "TX_PROCESS_INTERVAL",
		Usage:  "Interval for tx process",
		Value:  2 * time.Second,
		EnvVar: prefixEnvVar("TX_PROCESS_INTERVAL"),
	}

	// rollup tx gas base
	RollupTxGasBase = cli.Uint64Flag{
		Name:   "ROLLUP_TX_GAS_BASE",
		Usage:  "The base fee for a rollup transaction",
		Value:  530000,
		EnvVar: prefixEnvVar("ROLLUP_TX_GAS_BASE"),
	}
	// rollup tx gas per l1msg
	RollupTxGasPerL1Msg = cli.Uint64Flag{
		Name:   "ROLLUP_TX_GAS_PER_L1_MSG",
		Usage:  "The gas cost for each L1 message included in a rollup transaction",
		Value:  4200,
		EnvVar: prefixEnvVar("ROLLUP_TX_GAS_PER_L1_MSG"),
	}

	GasLimitBuffer = cli.Uint64Flag{
		Name:   "GAS_LIMIT_BUFFER",
		Usage:  "The gas limit buffer for a transaction",
		Value:  20, // add 20%
		EnvVar: prefixEnvVar("GAS_LIMIT_BUFFER"),
	}

	// journal path
	JournalFlag = cli.StringFlag{
		Name:   "JOURNAL_FILE_PATH",
		Usage:  "The path of the journal file",
		EnvVar: prefixEnvVar("JOURNAL_FILE_PATH"),
		Value:  "journal.rlp",
	}

	CalldataFeeBumpFlag = cli.Uint64Flag{
		Name:   "CALL_DATA_FEE_BUMP",
		Usage:  "The fee bump for call data",
		Value:  100, //fee = x * origin_fee/100
		EnvVar: prefixEnvVar("CALL_DATA_FEE_BUMP"),
	}

	MaxTxsInPendingPoolFlag = cli.Uint64Flag{
		Name:   "MAX_TXS_IN_PENDING_POOL",
		Usage:  "The maximum number of transactions in the pending pool",
		Value:  12,
		EnvVar: prefixEnvVar("MAX_TXS_IN_PENDING_POOL"),
	}

	RoughEstimateGasFlag = cli.BoolFlag{
		Name:   "ROUGH_ESTIMATE_GAS",
		Usage:  "Whether to use rough estimate gas",
		EnvVar: prefixEnvVar("ROUGH_ESTIMATE_GAS"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	L1EthRpcFlag,
	L2EthRpcsFlag,
	RollupAddressFlag,
	L1SequencerAddressFlag,
	PollIntervalFlag,
	TxTimeoutFlag,
	FinalizeFlag,
	MaxFinalizeNumFlag,
	PriorityRollupFlag,
	SubmitterAddressFlag,
	L2SequencerAddressFlag,
	PrivateKeyFlag,
	TxFeeLimitFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	MetricsServerEnable,
	MetricsHostname,
	MetricsPort,

	LogFilename,
	LogFileMaxSize,
	LogFileMaxAge,
	LogCompress,

	RollupInterval,
	FinalizeInterval,
	TxProcessInterval,

	RollupTxGasBase,
	RollupTxGasPerL1Msg,
	GasLimitBuffer,

	JournalFlag,

	CalldataFeeBumpFlag,
	MaxTxsInPendingPoolFlag,
	RoughEstimateGasFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
