terraform {
  backend "pg" {
    workspaces {
      name = "ipfs-platform"
    }
  }
}

provider "kubernetes" {
}

module "go_api_deployment_vault_abstraction" {
  source          = "github.com/blockchain-abstraction-middleware/deployment/modules/deployment"
  namespace       = "ipfs-platform"
  deployment_name = "vault-abstraction"
  docker_image    = "bamdockerhub/vault-abstraction"
  config_file     = "review.yml"
  config_path     = "/config/"
}