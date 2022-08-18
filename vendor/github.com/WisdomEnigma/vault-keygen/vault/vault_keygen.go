package vault

import "github.com/hashicorp/vault/api"

type Vault_Server struct{ Client *api.Client }

type Keygen struct {
	Vault_path   string
	Vault_record map[string]interface{}
}

type Vault_Services interface {
	SaveKeygen(collection Keygen) (*api.Secret, error)
	GetKeygen(key Keygen) (*api.Secret, error)
}

func NewClient(client *api.Client) Vault_Services { return &Vault_Server{Client: client} }

func (c *Vault_Server) SaveKeygen(collection Keygen) (*api.Secret, error) {

	return c.Client.Logical().Write(collection.Vault_path, collection.Vault_record)
}

func (c *Vault_Server) GetKeygen(key Keygen) (*api.Secret, error) {

	return c.Client.Logical().Read(key.Vault_path)
}
