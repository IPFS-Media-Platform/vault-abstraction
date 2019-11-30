package vaultmanager

import (
	"github.com/IPFS-Media-Platform/vault-abstraction/pkg/contracts/VaultManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
