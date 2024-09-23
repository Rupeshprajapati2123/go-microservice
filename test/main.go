// package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"
// 	"math/big"
// 	"os"
// 	"reflect"
// 	"sync"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"

// 	store "main/mainc" // Ensure this import path is correct for your project structure
// )

// type Params struct {
// 	Client     *ethclient.Client
// 	PrivateKey *ecdsa.PrivateKey
// 	Auth       *bind.TransactOpts
// 	Address    common.Address
// 	Instance   *store.Store
// }

// type GetTransactionDetailsRequest struct {
// 	TransactionId uint `json:"transactionId"`
// }

// type EventEmitter struct {
// 	eventChan chan *big.Int
// 	wg        sync.WaitGroup
// }

// func NewEventEmitter() *EventEmitter {
// 	return &EventEmitter{
// 		eventChan: make(chan *big.Int, 100), // Buffered channel to prevent blocking
// 	}
// }

// func (e *EventEmitter) Emit(id *big.Int) {
// 	e.eventChan <- id
// }

// func (e *EventEmitter) Listen(instance *store.Store, auth *bind.TransactOpts) {
// 	e.wg.Add(1)
// 	go func() {
// 		defer e.wg.Done()
// 		for id := range e.eventChan {
// 			fmt.Printf("Received Transaction ID: %s\n", id.String())
// 			tx, err := instance.ApproveTransaction(auth, id)
// 			if err != nil {
// 				fmt.Println("Failed to Approve Transaction ", err)
// 			}
			
// 			fmt.Println("Transaction Approved ", id)
// 			fmt.Println("hash", tx)
// 		}	
// 	}()
// }



// func (e *EventEmitter) Stop() {
// 	close(e.eventChan)
// 	e.wg.Wait()
// }

// func getParams() Params {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Panic("Error Loading env file")
// 		return Params{}
// 	}
// 	private_key := os.Getenv("SEPOLIA_PRIVATE_KEY")
// 	contract_address := os.Getenv("CONTRACT_ADDRESS")
// 	sepolia_url := os.Getenv("URL")

// 	client, err := ethclient.Dial(sepolia_url)
// 	if err != nil {
// 		log.Fatalf("Error connecting to Ethereum: %v", err)
// 	}
// 	fmt.Println("Successfully connected to Ethereum network")

// 	privateKey, err := crypto.HexToECDSA(private_key)
// 	if err != nil {
// 		log.Fatalf("Error loading private key: %v", err)
// 		return Params{}
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("Error casting public key to ECDSA")
// 		return Params{}
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatalf("Error getting nonce: %v", err)
// 		return Params{}
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error getting gas price: %v", err)
// 		return Params{}
// 	}

// 	auth := bind.NewKeyedTransactor(privateKey)
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)
// 	auth.GasLimit = uint64(300000)
// 	auth.GasPrice = gasPrice

// 	address := common.HexToAddress(contract_address)
// 	instance, err := store.NewStore(address, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return Params{
// 		Instance:   instance,
// 		Address:    address,
// 		Client:     client,
// 		Auth:       auth,
// 		PrivateKey: privateKey,
// 	}
// }

// func getContractBalance(c *fiber.Ctx, instance *store.Store) error {
// 	bal, err := instance.Balance(nil)
// 	if err != nil {
// 		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"balance": bal,
// 	})
// }

// func CreateTransaction(c *fiber.Ctx, client *ethclient.Client, instance *store.Store, auth *bind.TransactOpts, emitter *EventEmitter) error {
// 	type TransactionInput struct {
// 		Address string `json:"address"`
// 		Amount  uint   `json:"amount"`
// 	}

// 	var input TransactionInput

// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body: " + err.Error(),
// 		})
// 	}

// 	address := common.HexToAddress(input.Address)
// 	amount := big.NewInt(int64(input.Amount))

// 	tx, err := instance.CreateTransaction(auth, address, amount)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to create transaction: " + err.Error(),
// 		})
// 	}

// 	receipt, err := bind.WaitMined(context.Background(), client, tx)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to get transaction receipt: " + err.Error(),
// 		})
// 	}

// 	var transactionId *big.Int
// 	for _, log := range receipt.Logs {
// 		event, err := instance.ParseTransactionCreated(*log)
// 		if err == nil && event != nil {
// 			v := reflect.ValueOf(event).Elem()
// 			for i := 0; i < v.NumField(); i++ {
// 				if v.Field(i).Type() == reflect.TypeOf(big.NewInt(0)) {
// 					transactionId = v.Field(i).Interface().(*big.Int)
// 					break
// 				}
// 			}
// 			break
// 		}
// 	}

// 	if transactionId == nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to retrieve transaction ID from logs",
// 		})
// 	}

// 	emitter.Emit(transactionId)

// 	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
// 		"Status":        "Created",
// 		"TransactionID": transactionId.String(),
// 		"TxHash":        tx.Hash().Hex(),
// 	})
// }

// func AddEthToContract(c *fiber.Ctx, client *ethclient.Client, auth *bind.TransactOpts, contractAddress common.Address) error {
// 	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to get latest nonce: " + err.Error(),
// 		})
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))

// 	tenEth := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18))
	
// 	tx := types.NewTransaction(auth.Nonce.Uint64(), contractAddress, tenEth, auth.GasLimit, auth.GasPrice, nil)
	
// 	signedTx, err := auth.Signer(auth.From, tx)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to sign transaction: " + err.Error(),
// 		})
// 	}
	
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to send transaction: " + err.Error(),
// 		})
// 	}
	
// 	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to get transaction receipt: " + err.Error(),
// 		})
// 	}
	
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "Successfully added 10 ETH to the contract",
// 		"transactionHash": receipt.TxHash.Hex(),
// 	})
// }

// func getAllTransactions(c *fiber.Ctx, instance *store.Store) error {
// 	tx, err := instance.GetAllTransactions(nil)
// 	if err != nil {
// 		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
// 		"Status":       "Success",
// 		"Transactions": tx,
// 	})
// }

// func getTransactionDetails(c *fiber.Ctx, instance *store.Store) error {
// 	var req GetTransactionDetailsRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body: " + err.Error(),
// 		})
// 	}

// 	transactionId := req.TransactionId
// 	tx, err := instance.GetTransactionDetails(nil, big.NewInt(int64(transactionId)))
// 	if err != nil {
// 		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
// 		"Status":       "Success",
// 		"Transactions": tx,
// 	})
// }

// func ApproveTransaction(c *fiber.Ctx, instance *store.Store, auth *bind.TransactOpts) error {
// 	var req GetTransactionDetailsRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body: " + err.Error(),
// 		})
// 	}

// 	transactionId := req.TransactionId
// 	tx, err := instance.ApproveTransaction(auth, big.NewInt(int64(transactionId)))
// 	if err != nil {
// 		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
// 		"Status":       "Approved",
// 		"Transactions": tx,
// 	})
// }

// func AddApprovers(c *fiber.Ctx, instance *store.Store, auth *bind.TransactOpts) error {
// 	type AddApproversInput struct {
// 		Addresses []string `json:"addresses"`
// 	}

// 	var input AddApproversInput

// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body: " + err.Error(),
// 		})
// 	}

// 	var addresses []common.Address
// 	for _, addr := range input.Addresses {
// 		addresses = append(addresses, common.HexToAddress(addr))
// 	}

// 	tx, err := instance.AddApprovers(auth, addresses)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to add approvers: " + err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
// 		"Status": "Approvers Added",
// 		"TxHash": tx.Hash().Hex(),
// 	})
// }

// func main() {
// 	app := fiber.New()

// 	params := getParams()
// 	fmt.Printf("This is the params from the function: %+v\n", params)
	
// 	emitter := NewEventEmitter()
	
// 	// Start listening for events
// 	emitter.Listen(params.Instance, params.Auth)
	
// 	// Defer stopping the emitter
// 	defer emitter.Stop()

// 	app.Get("/balance", func(c *fiber.Ctx) error {
// 		return getContractBalance(c, params.Instance)
// 	})

// 	app.Post("/transaction", func(c *fiber.Ctx) error {
// 		return CreateTransaction(c, params.Client, params.Instance, params.Auth, emitter)
// 	})

// 	app.Post("/add-eth", func(c *fiber.Ctx) error {
// 		return AddEthToContract(c, params.Client, params.Auth, params.Address)
// 	})

// 	app.Get("/get-all-transactions", func(c *fiber.Ctx) error {
// 		return getAllTransactions(c, params.Instance)
// 	})

// 	app.Post("/get-transaction-details", func(c *fiber.Ctx) error {
// 		return getTransactionDetails(c, params.Instance)
// 	})
	
// 	app.Post("/approve-transaction", func(c *fiber.Ctx) error {
// 		return ApproveTransaction(c, params.Instance, params.Auth)
// 	})

// 	app.Post("/add-approvers", func(c *fiber.Ctx) error {
// 		return AddApprovers(c, params.Instance, params.Auth)
// 	})

// 	log.Fatal(app.Listen(":8080"))
// }