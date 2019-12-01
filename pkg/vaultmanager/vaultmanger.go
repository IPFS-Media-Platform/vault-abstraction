package vaultmanager

import (
	"math/big"

	"github.com/IPFS-Media-Platform/vault-abstraction/pkg/contracts/VaultManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager contains the login for the Manager Contract
type Manager struct {
	ec *ethclient.Client

	auth *bind.TransactOpts

	manager *VaultManager.VaultManager

	address string
}

// NewManager Creates a game manager contract with a goven address and private key
func NewManager(ec *ethclient.Client, vmAddress string, auth *bind.TransactOpts) (*Manager, error) {
	vaultManager, err := VaultManager.NewVaultManager(common.HexToAddress(vmAddress), ec)

	return &Manager{
		auth:    auth,
		ec:      ec,
		manager: vaultManager,
		address: vmAddress,
	}, err
}

// GetAllVaultAddresses gets all deployed game jam addresses
func (m *Manager) GetAllVaultAddresses() ([]common.Address, error) {
	return m.manager.GetVaults(&bind.CallOpts{Pending: true})
}

// Adds A New vault to
func (m *Manager) AddNewVault(name string, ipfsHash string) (*types.Transaction, error) {
	txnOpts := &bind.TransactOpts{
		From:     m.auth.From,
		Signer:   m.auth.Signer,
		GasLimit: 2381623,
		Value:    big.NewInt(0),
	}

	return m.manager.AddNewVault(txnOpts, name, ipfsHash)
}
