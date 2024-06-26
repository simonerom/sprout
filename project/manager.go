package project

import (
	"bytes"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project/contracts"
	utilscontract "github.com/machinefi/sprout/utils/contract"
)

type Manager struct {
	ipfsEndpoint string
	instance     *contracts.Contracts
	projects     sync.Map // projectID(uint64) -> *Project
	cache        *cache   // optional
}

func (m *Manager) Get(projectID uint64) (*Project, error) {
	var err error
	p, ok := m.projects.Load(projectID)
	if !ok {
		p, err = m.load(projectID)
		if err != nil {
			return nil, err
		}
	}
	return p.(*Project), nil
}

func (m *Manager) load(projectID uint64) (*Project, error) {
	emptyHash := [32]byte{}
	c, err := m.instance.Projects(nil, projectID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
	}
	if c.Uri == "" || bytes.Equal(c.Hash[:], emptyHash[:]) {
		return nil, errors.Errorf("the project not exist, project_id %v", projectID)
	}

	pm := &Meta{
		ProjectID: projectID,
		Uri:       c.Uri,
		Hash:      c.Hash,
	}

	var data []byte
	cached := true
	if m.cache != nil {
		data = m.cache.get(projectID, c.Hash[:])
	}
	if len(data) == 0 {
		cached = false
		data, err = pm.GetProjectRawData(m.ipfsEndpoint)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project raw data, project_id %v", projectID)
		}
	}
	if !cached && m.cache != nil {
		m.cache.set(projectID, data)
	}

	p, err := convertProject(data)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert project, project_id %v", projectID)
	}
	m.projects.Store(projectID, p)
	return p, nil
}

func (m *Manager) watchProjectContract(chainEndpoint, contractAddress string) error {
	projectCh, err := utilscontract.ListAndWatchProject(chainEndpoint, contractAddress)
	if err != nil {
		return err
	}

	go func() {
		for p := range projectCh {
			m.projects.Delete(p.ID)
		}
	}()
	return nil
}

// TODO support local project config
func NewManager(chainEndpoint, contractAddress, projectCacheDir, ipfsEndpoint string) (*Manager, error) {
	var c *cache
	var err error
	if projectCacheDir != "" {
		c, err = newCache(projectCacheDir)
		if err != nil {
			return nil, errors.Wrap(err, "failed to new cache")
		}
	}

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain, endpoint %s", chainEndpoint)
	}
	instance, err := contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to new contract instance, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}

	m := &Manager{
		ipfsEndpoint: ipfsEndpoint,
		instance:     instance,
		cache:        c,
	}
	if err := m.watchProjectContract(chainEndpoint, contractAddress); err != nil {
		return nil, err
	}
	return m, nil
}
