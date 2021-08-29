package structs

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	clientInstance *ethclient.Client = nil
	HeaderBlock    *big.Int          = big.NewInt(6677972)
)

// Data Block
type (
	DataBlock struct {
		BlockHeaderID        *types.Block
		BlockSenderID        *big.Int
		LengthTransaction    int
		BlockHex             string
		BlockHeader          uint64
		ListTransaction      []types.Transactions
		From                 common.Address
		To                   *common.Address
		Status               uint64
		Addresss             common.Hash
		CounterTransactions  uint
		TransactionsBlock    *types.Transaction
		TypeTransaction      common.Hash
		TypeDataTransactions *types.Transaction
		PendingStatus        bool
	}
	Block struct {
		Balance        *big.Int
		SenderBatchID  string
		RecieveBatchID string
		Amount         *big.Int
		Nonce          uint64
		GasPrice       *big.Int
		GasLimit       uint64
		Checkout       uint64
		DataBlock
	}

	BlockTransactionGateway struct{}

	// operations bit transactions
	BlockInterface interface {
		// Current transaction status, senderbatchid provided
		GetLastTransaction(node Block) (*big.Int, error)
		// get block header status
		GetBlockID() (*big.Int, error)
		// get blocks transactions , or say transactions in blocks
		GetBlockTransactions(b *types.Block) int
		// get hash of current block
		GetBlockHash(b *types.Block) string
		// get header number of a block
		GetBlockHeaderNumber(b *types.Block) uint64
		// get transactions list [0]
		GetBlockTransactionsList(b *types.Block) []types.Transactions
		// get array of transactions [...]Transactions{}
		GetBlockDataTransactions(b *types.Block) types.Transactions
		// get network id that define first genesis blockchain
		GetNetworkChainID() (*big.Int, error)
		// get whole block hash [f(Bx)] => HBx
		BlockHex(b *types.Block) common.Hash
		// get unit transactions against specfic hash block (HBx)
		CountBlockTransactions(hash common.Hash) (uint, error)
		// get all transactions in hash block (HBx)
		GetTransactionBlockRecord(hash common.Hash, iterate uint) (*types.Transaction, error)
		// get HBx transaction type [f(HBx)] => ([...]THBx{})Type
		GetTransactionType(t *types.Transaction) common.Hash
		// get pending Transaction [f([...]THBx{})Type] => THBx + 0
		GetHashTransactionType(hash common.Hash) (*types.Transaction, bool, error)
	}

	PaymentByCrypto struct {
		SenderBatchID            string
		SenderPrivateKey         *ecdsa.PrivateKey
		EthBlockHeader           string
		EthNewPublicKeyGenerator crypto.PublicKey
		EthNewPublic             *ecdsa.PublicKey
		EthAddress               common.Address
		EthNonceAtStatus         uint64
		EthGasUnits              *big.Int
		EthReciptAddress         common.Address
		EthTransaction           *types.Transaction
		FingerPrint              *types.Transaction
	}

	EthToken   struct{}
	BitsBlocks PaymentByCrypto

	EthereumBitsInterface interface {
		BTCECDSAHEX(node Block) (*ecdsa.PrivateKey, error)
		BTCHeaderBlockerID(node Block) string
		BTCECDSAPublic(key *ecdsa.PrivateKey) crypto.PublicKey
		BTCCryptoToKey(key crypto.PublicKey) *ecdsa.PublicKey
		BTCKeyToAddress(key *ecdsa.PublicKey) common.Address
		BTCNoncePendingStatus(address common.Address) (uint64, error)
		BTCGasConsumerPrice() (*big.Int, error)
		BTCNewTransactions(node Block, bit BitsBlocks) *types.Transaction
		BTCTransactionSignature(netID *big.Int, bit BitsBlocks) (*types.Transaction, error)
		TransferBTC(print *types.Transaction) error
	}
)

func (*BlockTransactionGateway) GetLastTransaction(node Block) (*big.Int, error) {

	err := errors.New("Empty Address")
	// require SenderID
	if node.SenderBatchID != "" {
		wallet := common.HexToAddress(node.SenderBatchID)
		pending := new(big.Int)
		node.Balance, err = clientInstance.BalanceAt(context.Background(), wallet, pending)
		if err != nil {
			log.Fatalln("[Request] Fail", err)
			return nil, err
		}
		return node.Balance, nil
	}
	return nil, err
}

func (gate *BlockTransactionGateway) GetBlockID(node Block) (*DataBlock, bool, error) {

	var err error

	node.BlockSenderID = HeaderBlock
	node.BlockHeaderID, err = clientInstance.BlockByNumber(context.Background(), node.BlockSenderID)

	if err != nil {
		return &node.DataBlock, false, err
	}

	// Block Data
	node.LengthTransaction = (*gate).GetBlockTransactionsLength(node.BlockHeaderID)
	node.BlockHex = (*gate).GetBlockHash(node.BlockHeaderID)
	node.BlockHeader = (*gate).GetBlockHeaderNumber(node.BlockHeaderID)
	node.ListTransaction[0] = (*gate).GetBlockDataTransactions(node.BlockHeaderID)

	// Get network chain id
	chainID, err := (*gate).GetNetworkChainID()
	if err != nil {
		return &node.DataBlock, false, err
	}

	for _, t := range node.ListTransaction[0] {

		message, err := t.AsMessage(types.NewEIP155Signer(chainID), new(big.Int))

		if err != nil {
			return &node.DataBlock, false, err
		}

		node.From = message.From()
		node.To = message.To()

		recipt, err := clientInstance.TransactionReceipt(context.Background(), t.Hash())
		if err != nil {
			return &node.DataBlock, false, err
		}

		node.Status = recipt.Status
	}

	node.Addresss = (*gate).BlockHex(node.BlockHeaderID)
	node.CounterTransactions, err = (*gate).CountBlockTransactions(node.Addresss)
	if err != nil {
		return &node.DataBlock, false, err
	}

	for i := uint(0); i < (node.CounterTransactions); i++ {

		node.TransactionsBlock, err = (*gate).GetTransactionBlockRecord(node.Addresss, i)
		if err != nil {
			return &node.DataBlock, false, err
		}

		node.TypeTransaction = (*gate).GetTransactionType(node.TransactionsBlock)
		if err != nil {
			return &node.DataBlock, false, err
		}

		node.TypeDataTransactions, node.PendingStatus, err = (*gate).GetHashTransactionType(node.TypeTransaction)
		if err != nil {
			return &node.DataBlock, false, err
		}
	}

	chainBlock := DataBlock{
		BlockHeaderID:        node.BlockHeaderID,
		BlockSenderID:        node.BlockSenderID,
		LengthTransaction:    node.LengthTransaction,
		BlockHex:             node.BlockHex,
		BlockHeader:          node.BlockHeader,
		ListTransaction:      node.ListTransaction,
		From:                 node.From,
		To:                   node.To,
		Status:               node.Status,
		Addresss:             node.Addresss,
		CounterTransactions:  node.CounterTransactions,
		TransactionsBlock:    node.TransactionsBlock,
		TypeTransaction:      node.TypeTransaction,
		TypeDataTransactions: node.TypeDataTransactions,
		PendingStatus:        node.PendingStatus,
	}
	return &chainBlock, true, nil
}

func (*BlockTransactionGateway) GetBlockTransactionsLength(b *types.Block) int {
	return len(b.Transactions())
}

func (*BlockTransactionGateway) GetBlockHash(b *types.Block) string {
	return b.Hash().Hex()
}

func (*BlockTransactionGateway) GetBlockHeaderNumber(b *types.Block) uint64 {
	return b.Number().Uint64()
}

func (B *BlockTransactionGateway) GetBlockTransactionsList(b *types.Block) []types.Transactions {
	return make([]types.Transactions, (*B).GetBlockTransactionsLength(b))
}

func (*BlockTransactionGateway) GetBlockDataTransactions(b *types.Block) types.Transactions {
	return b.Transactions()
}

func (*BlockTransactionGateway) GetNetworkChainID() (*big.Int, error) {
	return clientInstance.NetworkID(context.Background())
}

func (*BlockTransactionGateway) BlockHex(b *types.Block) common.Hash {
	return common.HexToHash(b.Hash().Hex())
}
func (*BlockTransactionGateway) CountBlockTransactions(hash common.Hash) (uint, error) {
	return clientInstance.TransactionCount(context.Background(), hash)
}

func (*BlockTransactionGateway) GetTransactionBlockRecord(hash common.Hash, iterate uint) (*types.Transaction, error) {
	return clientInstance.TransactionInBlock(context.Background(), hash, iterate)
}

func (*BlockTransactionGateway) GetTransactionType(t *types.Transaction) common.Hash {
	return common.HexToHash(t.Hash().Hex())
}

func (*BlockTransactionGateway) GetHashTransactionType(hash common.Hash) (*types.Transaction, bool, error) {
	return clientInstance.TransactionByHash(context.Background(), hash)
}

func (*EthToken) BTCECDSAHEX(node Block) (*ecdsa.PrivateKey, error) {
	return ethCrypto.HexToECDSA(node.SenderBatchID)
}
func (*EthToken) BTCHeaderBlockerID(node Block) string {
	return node.BlockHeaderID.Number().String()
}
func (*EthToken) BTCECDSAPublic(key *ecdsa.PrivateKey) crypto.PublicKey {
	return key.Public()
}

func (*EthToken) BTCCryptoToKey(key crypto.PublicKey) *ecdsa.PublicKey {
	return key.(*ecdsa.PublicKey)
}

func (*EthToken) BTCKeyToAddress(key *ecdsa.PublicKey) common.Address {
	return ethCrypto.PubkeyToAddress(*key)
}

func (*EthToken) BTCNoncePendingStatus(address common.Address) (uint64, error) {
	return clientInstance.PendingNonceAt(context.Background(), address)
}

func (*EthToken) BTCGasConsumerPrice() (*big.Int, error) {
	return clientInstance.SuggestGasPrice(context.Background())
}

func (*EthToken) BTCAddressHex(address string) common.Address {
	return common.HexToAddress(address)
}

func (*EthToken) BTCNetworkID() (*big.Int, error) {
	return clientInstance.NetworkID(context.Background())
}

func (*EthToken) BTCNewTransactions(node Block, bit BitsBlocks) *types.Transaction {
	var nofield []byte
	return types.NewTransaction(node.Nonce, bit.EthReciptAddress, node.Amount, node.GasLimit, node.GasPrice, nofield)
}

func (*EthToken) BTCTransactionSignature(netID *big.Int, bit BitsBlocks) (*types.Transaction, error) {
	return types.SignTx(bit.EthTransaction, types.NewEIP155Signer(netID), bit.SenderPrivateKey)
}

func (*EthToken) TransferBTC(print *types.Transaction) error {
	return clientInstance.SendTransaction(context.Background(), print)
}
