package routes

import (
	"encoding/json"
	"net/http"

	"github.com/IPFS-Media-Platform/vault-abstraction/pkg/config"
	"github.com/IPFS-Media-Platform/vault-abstraction/pkg/contracts/Vault"
	"github.com/IPFS-Media-Platform/vault-abstraction/pkg/ethereum"
	"github.com/IPFS-Media-Platform/vault-abstraction/pkg/vaultmanager"
	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// ManagerResource implements server.Resource interface
type ManagerResource struct {
	path   string
	config *config.Config
}

// NewResource constructor func
func (r *ManagerResource) NewResource(path string, config *config.Config) *ManagerResource {
	return &ManagerResource{
		path:   path,
		config: config,
	}
}

// Path returns the Resource base path
func (r *ManagerResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *ManagerResource) Routes() http.Handler {
	router := chi.NewRouter()
	// router.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router.Use(cors.Handler)

	privKey, err := crypto.HexToECDSA(r.config.Keys.Admin)
	if err != nil {
		log.WithError(err).Error("Failed to load key")
	}

	ec, err := ethereum.CreateEthClient(r.config.Infura.URL + r.config.Infura.Key)
	if err != nil {
		log.WithError(err).Error("Failed to initialize ethereum client")
	}

	auth := bind.NewKeyedTransactor(privKey)

	vaultManagerContract, err := vaultmanager.NewManager(ec, r.config.Contracts.VaultManagerAddress, auth)
	if err != nil {
		log.WithError(err).Error("Failed to create vault manager")
	}

	router.Get("/get-all-vaults", r.getAllAddresses(vaultManagerContract, ec))
	router.Post("/add-vault", r.addNewVault(vaultManagerContract))

	log.WithFields(log.Fields{"Contract": "Vault Manager", "Address": r.config.Contracts.VaultManagerAddress}).Info("Created manager abstraction")

	return router
}

type VaultResponse struct {
	IpfsHash [50]string
}

func (r *ManagerResource) getAllAddresses(vm *vaultmanager.Manager, ec *ethclient.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		vaultAddressesList, err := vm.GetAllVaultAddresses()
		if err != nil {
			log.WithError(err).Error("GetAllVaultAddresses function on Game Manager contract failed")
		}

		var ipfsHashes [50]string
		for i, address := range vaultAddressesList {
			vault, err := Vault.NewVault(common.HexToAddress(address.String()), ec)
			if err != nil {
				log.Fatal("Wrong", err)
			}

			vaultIpfsHash, err := vault.Get(&bind.CallOpts{Pending: true})
			ipfsHashes[i] = vaultIpfsHash
		}

		payload := &VaultResponse{
			IpfsHash: ipfsHashes,
		}

		json, err := json.Marshal(payload)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}

type AddNewVaultPayload struct {
	IpfsHash string `json:"ipfsHash"`
	Name     string `json:"name"`
}

func (r *ManagerResource) addNewVault(vm *vaultmanager.Manager) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		addNewVaultPayload := new(AddNewVaultPayload)
		json.NewDecoder(req.Body).Decode(&addNewVaultPayload)

		_, err := vm.AddNewVault(addNewVaultPayload.Name, addNewVaultPayload.IpfsHash)
		if err != nil {
			log.WithError(err).Error("GetAllVaultAddresses function on Game Manager contract failed")
		}

		payload := struct {
			VaultStatus string `json:"vaultStatus"`
			StatusCode  int    `json:"statusCode"`
		}{
			addNewVaultPayload.Name,
			http.StatusCreated,
		}

		json, _ := json.Marshal(payload)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}
