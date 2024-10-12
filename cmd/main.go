package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/lib/lotuslog"
	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/node/repo"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"go.opencensus.io/stats/view"
	"golang.org/x/xerrors"
	"io/ioutil"
	"math/big"
	"net/http"
	"nhooyr.io/websocket"
	"os"
	"strconv"
	"strings"
	"time"
)

var log = logging.Logger("main")

const FlagWalletRepo = "wallet-repo"
const DcPayChannelAddress = "0xc0C05B9eD79e5CAf2A5432B08B689bf499ECa8EF"
const RpcUrl = "https://api.node.glif.io"
const ChainId = 314

func main() {
	lotuslog.SetupLogLevels()

	local := []*cli.Command{
		walletNew,
		walletSetDefault,
		getWalletDefault,
		walletList,
		walletImport,
		walletExport,
		stake,
		unlock,
		getAllStakeInfo,
	}

	app := &cli.App{
		Name:    "lotus-wallet",
		Usage:   "Basic external wallet",
		Version: "v1.28.1",
		Description: `
lotus-wallet provides a remote wallet service for lotus.

To configure your lotus node to use a remote wallet:
* Run 'lotus-wallet get-api-key' to generate API key
* Start lotus-wallet using 'lotus-wallet run' (see --help for additional flags)
* Edit lotus config (~/.lotus/config.toml)
  * Find the '[Wallet]' section
  * Set 'RemoteBackend' to '[api key]:http://[wallet ip]:[wallet port]'
    (the default port is 1777)
* Start (or restart) the lotus daemon`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    FlagWalletRepo,
				EnvVars: []string{"WALLET_PATH"},
				Value:   "~/.cht-lending-cli", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus",
			},
		},

		Commands: local,
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		//return
	}
}

var walletList = &cli.Command{
	Name:  "list",
	Usage: "List wallet address",
	Action: func(cctx *cli.Context) error {

		log.Info("Starting lotus list ")

		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		list, err := api.WalletList(ctx)

		if err != nil {
			return err
		}

		for _, a := range list {
			fmt.Println(a)
		}

		return nil
	},
}

var walletSetDefault = &cli.Command{
	Name:      "set-default",
	Usage:     "Set default wallet address",
	ArgsUsage: "[address]",
	Action: func(cctx *cli.Context) error {

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return err
		}

		fmt.Println("Default address set to:", addr)
		return api.SetDefault(addr)
	},
}

var walletNew = &cli.Command{
	Name:      "new",
	Usage:     "Generate a new key of the given type",
	ArgsUsage: "[bls|secp256k1 (default secp256k1)]",
	Action: func(cctx *cli.Context) error {

		log.Info("Starting lotus new ")

		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		t := cctx.Args().First()
		if t == "" {
			t = "secp256k1"
		}

		nk, err := api.WalletNew(ctx, types.KeyType(t))
		if err != nil {
			return err
		}

		fmt.Println(nk.String())

		return nil
	},
}

var walletImport = &cli.Command{
	Name:      "import",
	Usage:     "import keys",
	ArgsUsage: "[<path> (optional, will read from stdin if omitted)]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "specify input format for key",
			Value: "hex-lotus",
		},
		&cli.BoolFlag{
			Name:  "as-default",
			Usage: "import the given key as your new default key",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		//ctx := ReqContext(cctx)

		var inpdata []byte
		if !cctx.Args().Present() || cctx.Args().First() == "-" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter private key: ")
			indata, err := reader.ReadBytes('\n')
			if err != nil {
				return err
			}
			inpdata = indata

		} else {
			fdata, err := ioutil.ReadFile(cctx.Args().First())
			if err != nil {
				return err
			}
			inpdata = fdata
		}

		var ki types.KeyInfo
		switch cctx.String("format") {
		case "hex-lotus":
			data, err := hex.DecodeString(strings.TrimSpace(string(inpdata)))
			if err != nil {
				return err
			}

			if err := json.Unmarshal(data, &ki); err != nil {
				return err
			}
		case "json-lotus":
			if err := json.Unmarshal(inpdata, &ki); err != nil {
				return err
			}
		case "gfc-json":
			var f struct {
				KeyInfo []struct {
					PrivateKey []byte
					SigType    int
				}
			}
			if err := json.Unmarshal(inpdata, &f); err != nil {
				return xerrors.Errorf("failed to parse go-filecoin key: %s", err)
			}

			gk := f.KeyInfo[0]
			ki.PrivateKey = gk.PrivateKey
			switch gk.SigType {
			case 1:
				ki.Type = types.KTSecp256k1
			case 2:
				ki.Type = types.KTBLS
			default:
				return fmt.Errorf("unrecognized key type: %d", gk.SigType)
			}
		default:
			return fmt.Errorf("unrecognized format: %s", cctx.String("format"))
		}

		addr, err := api.WalletImport(ctx, &ki)
		if err != nil {
			return err
		}

		//if cctx.Bool("as-default") {
		//	if err := api.WalletSetDefault(ctx, addr); err != nil {
		//		return fmt.Errorf("failed to set default key: %w", err)
		//	}
		//}

		fmt.Printf("imported key %s successfully!\n", addr)
		return nil
	},
}

var walletExport = &cli.Command{
	Name:      "export",
	Usage:     "export keys",
	ArgsUsage: "[address]",
	Action: func(cctx *cli.Context) error {

		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return err
		}

		ki, err := api.WalletExport(ctx, addr)
		if err != nil {
			return err
		}

		b, err := json.Marshal(ki)
		if err != nil {
			return err
		}

		fmt.Println(hex.EncodeToString(b))
		return nil
	},
}

var getWalletDefault = &cli.Command{
	Name:      "get-default",
	Usage:     "sign a message",
	ArgsUsage: "<signing address> <hexMessage>",
	Action: func(cctx *cli.Context) error {

		// Register all metric views
		if err := view.Register(
			metrics.DefaultViews...,
		); err != nil {
			log.Fatalf("Cannot register the view: %v", err)
		}

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		lw, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		getDefault, err := lw.GetDefault()

		if err != nil {
			return err
		}

		fmt.Println(getDefault)

		return nil
	},
}

func openRepo(cctx *cli.Context) (repo.LockedRepo, types.KeyStore, error) {
	repoPath := cctx.String(FlagWalletRepo)
	r, err := repo.NewFS(repoPath)
	if err != nil {
		return nil, nil, err
	}

	ok, err := r.Exists()
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		if err := r.Init(repo.Worker); err != nil {
			return nil, nil, err
		}
	}

	lr, err := r.Lock(repo.Wallet)
	if err != nil {
		return nil, nil, err
	}

	ks, err := lr.KeyStore()
	if err != nil {
		return nil, nil, err
	}

	return lr, ks, nil
}

type KeyEth struct {
	Type       string `json:"Type"`
	PrivateKey string `json:"PrivateKey"`
}

var stake = &cli.Command{
	Name:      "stake",
	Usage:     "The client pledges and fills coins according to the application limit. 10T DC needs to pledge 1FIL. The amount of pledge is determined according to the DC limit applied by the client",
	ArgsUsage: "[data-set-type dataCap]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "data-set-type",
			Usage: "0-Common data set 1-Private data set",
		},
		&cli.StringFlag{
			Name:  "dataCap",
			Usage: "In Tib",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "Your application dataCap wallet address",
		},
	},
	Action: func(cctx *cli.Context) error {

		dt := cctx.String("data-set-type")

		if dt == "" {
			return fmt.Errorf("data-set-type cannot be empty:  ")
		}

		dc := cctx.String("dataCap")

		if dc == "" {
			return fmt.Errorf("dataCap cannot be empty:  ")
		}

		fm := cctx.String("from")

		if fm == "" {
			return fmt.Errorf("walletAddress cannot be empty:  ")
		}

		nodeAPI, closer, err := getFullNodeAPIV1()

		if err != nil {
			return err
		}
		defer closer()

		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		fa, err := address.NewFromString(fm)
		if err != nil {
			return err
		}

		ki, err := api.WalletExport(ctx, fa)
		if err != nil {
			return err
		}

		var keyEth KeyEth

		b, err := json.Marshal(ki)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &keyEth)
		if err != nil {
			return fmt.Errorf("Error decoding JSON: %w ", err)
		}

		privateKeyBytes, err := base64.StdEncoding.DecodeString(keyEth.PrivateKey)

		if err != nil {
			return err
		}

		sprintfETH := fmt.Sprintf("%x", privateKeyBytes)

		num, err := strconv.Atoi(dc)
		// 计算所需FIL质押数量
		stakeFil := (num / 10) * 1
		sf := big.NewInt(int64(stakeFil))

		multiplier := big.NewInt(1e18)

		// Multiply sf by 10^18
		sf.Mul(sf, multiplier)

		userBalance, err := nodeAPI.WalletBalance(ctx, fa)

		userBalanceInt := userBalance.Int

		if sf.Cmp(userBalanceInt) > 0 {
			return fmt.Errorf("Insufficient balance\n")
		}

		if err != nil {
			return err
		}

		client, err := ethclient.Dial(RpcUrl)
		if err != nil {
			return fmt.Errorf("Failed to connect to the Ethereum client: %v ", err)
		}

		contractAddress := common.HexToAddress(DcPayChannelAddress)
		DcPayChannelContract, err := NewDcPayChannel(contractAddress, client)
		if err != nil {
			return fmt.Errorf("Failed to create contract instance: %v ", err)
		}

		privateKey, err := crypto.HexToECDSA(sprintfETH)
		if err != nil {
			return fmt.Errorf("Failed to convert private key: %v ", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("Failed to cast public key to ECDSA ")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return fmt.Errorf("Failed to retrieve account nonce: %v ", err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return fmt.Errorf("Failed to retrieve suggested gas price: %v ", err)
		}
		chainId := big.NewInt(int64(ChainId))

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = sf                   // in wei (0 if not transferring ether)
		auth.GasLimit = uint64(188028833) // in units
		auth.GasPrice = gasPrice

		dataCap, _ := new(big.Int).SetString(dc, 10)
		dataSetType, _ := new(big.Int).SetString(dt, 10)

		transaction, err := DcPayChannelContract.Stake(auth, dataCap, dataSetType)

		if err != nil {
			fmt.Errorf("Failed signedTx: %v ", err)
		}

		fmt.Printf("Transaction sent: %s\n", transaction.Hash().Hex())

		// Wait for the transaction to be mined
		fmt.Println("Waiting for transaction to be verify...")
		receipt, err := bind.WaitMined(context.Background(), client, transaction)
		if err != nil {
			log.Fatalf("Failed to wait for transaction to be mined: %v", err)
		}

		// Check the status of the transaction
		if receipt.Status == 1 {
			fmt.Printf("Transaction successfully in block %d\n", receipt.BlockNumber.Uint64())
		} else {
			return fmt.Errorf("Transaction failed ")
		}

		return nil
	},
}

var unlock = &cli.Command{
	Name:      "unlock",
	Usage:     "Unlock the pledged fil",
	ArgsUsage: "[address stakeId]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "stakeId",
			Usage: "1-Common data set 2-Private data set",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "Your application dataCap wallet address",
		},
	},
	Action: func(cctx *cli.Context) error {

		si := cctx.String("stakeId")

		if si == "" {
			return fmt.Errorf("stakeId cannot be empty:  ")
		}

		fm := cctx.String("from")

		if fm == "" {
			return fmt.Errorf("walletAddress cannot be empty:  ")
		}

		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		fa, err := address.NewFromString(fm)
		if err != nil {
			return err
		}

		ki, err := api.WalletExport(ctx, fa)
		if err != nil {
			return err
		}

		var keyEth KeyEth

		b, err := json.Marshal(ki)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &keyEth)
		if err != nil {
			return fmt.Errorf("Error decoding JSON: %w ", err)
		}

		privateKeyBytes, err := base64.StdEncoding.DecodeString(keyEth.PrivateKey)

		if err != nil {
			return err
		}

		sprintfETH := fmt.Sprintf("%x\n", privateKeyBytes)

		client, err := ethclient.Dial(RpcUrl)
		if err != nil {
			return fmt.Errorf("Failed to connect to the Ethereum client: %v ", err)
		}

		contractAddress := common.HexToAddress(DcPayChannelAddress)
		DcPayChannelContract, err := NewDcPayChannel(contractAddress, client)
		if err != nil {
			return fmt.Errorf("Failed to create contract instance: %v ", err)
		}

		privateKey, err := crypto.HexToECDSA(sprintfETH)
		if err != nil {
			return fmt.Errorf("Failed to convert private key: %v ", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("Failed to cast public key to ECDSA ")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return fmt.Errorf("Failed to retrieve account nonce: %v ", err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return fmt.Errorf("Failed to retrieve suggested gas price: %v ", err)
		}
		chainId := big.NewInt(int64(ChainId))

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)        // in wei (0 if not transferring ether)
		auth.GasLimit = uint64(188028833) // in units
		auth.GasPrice = gasPrice

		stakeId, _ := new(big.Int).SetString(si, 10)

		transaction, err := DcPayChannelContract.Unlock(auth, stakeId)

		if err != nil {
			fmt.Errorf("Failed signedTx: %v ", err)
		}

		fmt.Printf("Transaction sent: %s\n", transaction.Hash().Hex())

		// Wait for the transaction to be mined
		fmt.Println("Waiting for transaction to be verify...")
		receipt, err := bind.WaitMined(context.Background(), client, transaction)
		if err != nil {
			log.Fatalf("Failed to wait for transaction to be mined: %v", err)
		}

		// Check the status of the transaction
		if receipt.Status == 1 {
			fmt.Printf("Transaction successfully in block %d\n", receipt.BlockNumber.Uint64())
		} else {
			return fmt.Errorf("Transaction failed ")
		}

		return nil
	},
}

var getAllStakeInfo = &cli.Command{
	Name:      "getAllStakeInfo",
	Usage:     "Get full pledge information",
	ArgsUsage: "[address stakeId]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "Your application dataCap wallet address",
		},
	},
	Action: func(cctx *cli.Context) error {

		fm := cctx.String("from")

		if fm == "" {
			return fmt.Errorf("walletAddress cannot be empty:  ")
		}

		ctx := context.Background()

		lr, ks, err := openRepo(cctx)
		if err != nil {
			return err
		}
		defer lr.Close() // nolint

		api, err := wallet.NewWallet(ks)
		if err != nil {
			return err
		}

		fa, err := address.NewFromString(fm)
		if err != nil {
			return err
		}

		ki, err := api.WalletExport(ctx, fa)
		if err != nil {
			return err
		}

		var keyEth KeyEth

		b, err := json.Marshal(ki)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &keyEth)
		if err != nil {
			return fmt.Errorf("Error decoding JSON: %w ", err)
		}

		privateKeyBytes, err := base64.StdEncoding.DecodeString(keyEth.PrivateKey)

		if err != nil {
			return err
		}

		sprintfETH := fmt.Sprintf("%x", privateKeyBytes)

		client, err := ethclient.Dial(RpcUrl)
		if err != nil {
			return fmt.Errorf("Failed to connect to the Ethereum client: %v ", err)
		}

		contractAddress := common.HexToAddress(DcPayChannelAddress)
		DcPayChannelContract, err := NewDcPayChannel(contractAddress, client)
		if err != nil {
			return fmt.Errorf("Failed to create contract instance: %v ", err)
		}

		privateKey, err := crypto.HexToECDSA(sprintfETH)
		if err != nil {
			return fmt.Errorf("Failed to convert private key: %v ", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("Failed to cast public key to ECDSA ")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		callOpts := &bind.CallOpts{
			Pending: false,
			From:    fromAddress,
			Context: context.Background(),
		}

		transaction, err := DcPayChannelContract.GetAllStakeInfo(callOpts, fromAddress)

		if err != nil {
			fmt.Errorf("Failed signedTx: %v ", err)
		}
		// 打印表头
		printStakeInfoArray(transaction)

		return nil
	},
}

// 格式化输出函数
// 格式化输出函数
func printStakeInfoArray(stakes []DcPayChannelStakeInfo) {
	// 打印表头
	fmt.Printf("%-10s %-25s %-20s %-20s %-10s\n", "StakeId", "Amount (FIL)", "Start Time", "Lock Time (days)", "Active")

	// 遍历数组并打印每个对象
	for _, stake := range stakes {
		// 将 Unix 时间戳转换为人类可读的时间格式
		startTime := time.Unix(stake.StartTime.Int64(), 0).Format("2006-01-02 15:04:05")

		// 计算 Amount 往前进 18 个零 (即从 attoFIL 转换为 FIL)
		amountInFIL := new(big.Float).Quo(new(big.Float).SetInt(stake.Amount), big.NewFloat(1e18))

		// 打印每个字段，使用适当的宽度格式化输出
		fmt.Printf("%-10s %-25s %-20s %-20d %-10t\n",
			stake.StakeId.String(),
			amountInFIL.Text('f', 18), // Amount 显示为 FIL，保留 18 位小数
			startTime,
			stake.LockTime.Int64()/(24*3600), // 将秒数转为天数
			stake.Active,
		)
	}
}

// Connect to the WebSocket endpoint
// Function to establish WebSocket connection
func connectToWebSocket(endpoint string, headers http.Header) (*websocket.Conn, error) {
	conn, _, err := websocket.Dial(context.Background(), endpoint, &websocket.DialOptions{
		HTTPHeader: headers,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	return conn, nil
}

// Function to return FullNode and jsonrpc.ClientCloser
func getFullNodeAPIV1() (api.FullNode, jsonrpc.ClientCloser, error) {
	// WebSocket endpoint for Lotus API
	endpoint := "wss://wss.node.glif.io/apigw/lotus/rpc/v1"

	// Add necessary headers (e.g., Authorization)
	headers := http.Header{}
	//headers.Add("Authorization", "Bearer YOUR_API_TOKEN_HERE") // Replace with your token

	// Establish WebSocket connection
	conn, err := connectToWebSocket(endpoint, headers)
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close(websocket.StatusInternalError, "closing connection")

	// Construct the FullNode API client over WebSocket
	var fullNode v1api.FullNodeStruct
	fd, closer, err := NewFullNodeRPCV1(context.TODO(), endpoint, headers, &fullNode)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to create Lotus RPC client: %w", err)
	}

	return fd, closer, nil
}

// NewFullNodeRPCV1 sets up the Lotus FullNode RPC client
func NewFullNodeRPCV1(ctx context.Context, addr string, requestHeader http.Header, target *v1api.FullNodeStruct, opts ...jsonrpc.Option) (api.FullNode, jsonrpc.ClientCloser, error) {
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin",
		api.GetInternalStructs(target), requestHeader, append([]jsonrpc.Option{jsonrpc.WithErrors(api.RPCErrors)}, opts...)...)

	return target, closer, err
}
