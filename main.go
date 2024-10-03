package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	store "main/mainc"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

var syncTime time.Time;
var numOfTransactions int;

type Params struct {
	Client     *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
	Address    common.Address
	Instance   *store.Store
}


type GetTransactionDetailsRequest struct {
	TransactionId uint `json:"transactionId"`
}


func approveTransaction(instance *store.Store, auth *bind.TransactOpts, id *big.Int, client *ethclient.Client) error {
	if err := updateNonce(client, auth); err != nil {
		return fmt.Errorf("failed to update nonce: %v", err)
	}

	tx, err := instance.ApproveTransaction(auth, id)
	if err != nil {
		return fmt.Errorf("failed to approve transaction %s: %v", id.String(), err)
	}
	fmt.Printf("Transaction %s approved. TX Hash: %s\n", id.String(), tx.Hash().Hex())
	return nil
}


func updateNonce(client *ethclient.Client, auth *bind.TransactOpts) error {
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	return nil
}

func getParams() Params {

	private_key := os.Getenv("SEPOLIA_PRIVATE_KEY")
	contract_address := os.Getenv("CONTRACT_ADDRESS")
	sepolia_url := os.Getenv("URL")

	var returnStruct Params

	client, err := ethclient.Dial(sepolia_url)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum: %v", err)
	}
	fmt.Println("Successfully connected to Ethereum network")

	privateKey, err := crypto.HexToECDSA(private_key)
	if err != nil {
		log.Fatalf("Error loading private key: %v", err)
		return Params{}
	}

	publicKey := privateKey.Public()]
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
		return Params{}
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting gas price: %v", err)
		return Params{}
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)     
	auth.GasLimit = uint64(300000) 
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contract_address)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}    
	px,erro:= instance.GetAllTransactions(nil)
	if erro!=nil{
		log.Fatal("error setting number of transactions", err)
	}
	numOfTransactions= len(px)
	syncTime = time.Now();
	returnStruct.Instance = instance
	returnStruct.Address = address
	returnStruct.Client = client
	returnStruct.Auth = auth
	returnStruct.PrivateKey = privateKey
	return returnStruct
}

func getContractBalance(c *fiber.Ctx, instance *store.Store) error {
	bal, err := instance.Balance(nil)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"balance": bal,
	})
}

func CreateTransaction(c *fiber.Ctx, client *ethclient.Client, instance *store.Store, auth *bind.TransactOpts) error {
	

	if err := updateNonce(client, auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update nonce: " + err.Error(),
		})
	}

	type TransactionInput struct {
		Address string `json:"address"`
		Amount  uint   `json:"amount"`
	}

	var input TransactionInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	address := common.HexToAddress(input.Address)
	amount := big.NewInt(int64(input.Amount))

	tx, err := instance.CreateTransaction(auth, address, amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction: " + err.Error(),
		})
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transaction receipt: " + err.Error(),
		})
	}

	txx, err := instance.GetAllTransactions(nil)
	if err != nil {
		log.Fatal("There seems to be some error in fetching the ID of the latest transaction ", err)
	}

	transactionId := big.NewInt(int64(len(txx) - 1))
	
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Status":        "Created",
		"TransactionID": transactionId,
		"TxHash":        receipt.TxHash.Hex(),
	})
}

func AddEthToContract(c *fiber.Ctx, client *ethclient.Client, auth *bind.TransactOpts, contractAddress common.Address) error {
	if err := updateNonce(client, auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update nonce: " + err.Error(),
		})
	}

	tenEth := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18))
	
	tx := types.NewTransaction(auth.Nonce.Uint64(), contractAddress, tenEth, auth.GasLimit, auth.GasPrice, nil)
	
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to sign transaction: " + err.Error(),
		})
	}
	
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send transaction: " + err.Error(),
		})
	}
	
	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transaction receipt: " + err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully added 10 ETH to the contract",
		"transactionHash": receipt.TxHash.Hex(),
	})
}

func getAllTransactions(c *fiber.Ctx, instance *store.Store) error {
	tx, err := instance.GetAllTransactions(nil)
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Status":      "Success",
		"Transactions": tx,
	})
}

func getTransactionDetails(c *fiber.Ctx, instance *store.Store) error {
	var req GetTransactionDetailsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	transactionId := req.TransactionId
	tx, err := instance.GetTransactionDetails(nil, big.NewInt(int64(transactionId)))
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Status":      "Success",
		"Transaction": tx,
	})
}

func ApproveTransaction(c *fiber.Ctx, client *ethclient.Client, instance *store.Store, auth *bind.TransactOpts) error {
	if err := updateNonce(client, auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update nonce: " + err.Error(),
		})
	}

	var req GetTransactionDetailsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	transactionId := req.TransactionId
	tx, err := instance.ApproveTransaction(auth, big.NewInt(int64(transactionId)))
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Status":     "Approved",
		"Transaction": tx.Hash().Hex(),
	})
}

func AddApprovers(c *fiber.Ctx, client *ethclient.Client, instance *store.Store, auth *bind.TransactOpts) error {
	if err := updateNonce(client, auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update nonce: " + err.Error(),
		})
	}

	type AddApproversInput struct {
		Addresses []string `json:"addresses"`
	}

	var input AddApproversInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	var addresses []common.Address
	for _, addr := range input.Addresses {
		addresses = append(addresses, common.HexToAddress(addr))
	}

	tx, err := instance.AddApprovers(auth, addresses)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add approvers: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Status": "Approvers Added",
		"TxHash": tx.Hash().Hex(),
	})
}


func temp(instance *store.Store, auth *bind.TransactOpts, client *ethclient.Client){
	for x:= range time.Tick(10*time.Second) {
		// fmt.Println("syncing at time" , x);
		tx,err :=instance.GetAllTransactions(nil);
		if err!=nil {
			log.Fatal("There seems to be an error fetching the transactions: ", err)
		}
		if len(tx)-numOfTransactions ==0 {
			fmt.Println(len(tx)-numOfTransactions ,"added in ",x.Hour(),":",x.Minute(),":",x.Second()," --- ",syncTime.Hour(),":",syncTime.Minute(),":",syncTime.Second());
		}else{
			for idx:=numOfTransactions;idx<len(tx);idx++ {
				er:= approveTransaction(instance,auth, big.NewInt(int64(idx)),client);
				if er != nil {
					log.Fatal("failed to approve transaction ",er)
				}
				fmt.Println("Transaction Approved with Id ", idx)
			}
		} 
		numOfTransactions=len(tx)
		syncTime=x
	}
	fmt.Println("");
}
func main() {
	app := fiber.New()

	params := getParams()
	fmt.Printf("This is the params from the function: %+v\n", params)
	go temp(params.Instance,params.Auth,params.Client);

	app.Get("/balance", func(c *fiber.Ctx) error {
		return getContractBalance(c, params.Instance)
	})

	app.Post("/transaction", func(c *fiber.Ctx) error {
		return CreateTransaction(c, params.Client, params.Instance, params.Auth)
	})

	app.Post("/add-eth", func(c *fiber.Ctx) error {
		return AddEthToContract(c, params.Client, params.Auth, params.Address)
	})

	app.Get("/get-all-transactions", func(c *fiber.Ctx) error {
		return getAllTransactions(c, params.Instance)
	})

	app.Post("/get-transaction-details", func(c *fiber.Ctx) error {
		return getTransactionDetails(c, params.Instance)
	})
	
	app.Post("/approve-transaction", func(c *fiber.Ctx) error {
		return ApproveTransaction(c, params.Client, params.Instance, params.Auth)
	})

	app.Post("/add-approvers", func(c *fiber.Ctx) error {
		return AddApprovers(c, params.Client, params.Instance, params.Auth)
	})

	// app.Get("/failed-transactions", func(c *fiber.Ctx) error {
		
	// })
	log.Fatal(app.Listen(":8080"))
	

}