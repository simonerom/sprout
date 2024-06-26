package datasource

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gorm.io/datatypes"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/machinefi/sprout/types"
)

type message struct {
	gorm.Model
	MessageID      string `gorm:"index:message_id,not null"`
	ClientID       string `gorm:"index:message_fetch,not null,default:''"`
	ProjectID      uint64 `gorm:"index:message_fetch,not null"`
	ProjectVersion string `gorm:"index:message_fetch,not null,default:'0.0'"`
	Data           []byte `gorm:"size:4096"`
	InternalTaskID string `gorm:"index:internal_task_id,not null,default:''"`
}

type task struct {
	gorm.Model
	ProjectID      uint64         `gorm:"index:task_fetch,not null"`
	InternalTaskID string         `gorm:"index:internal_task_id,not null"`
	MessageIDs     datatypes.JSON `gorm:"not null"`
	Signature      string         `gorm:"not null,default:''"`
}

type postgres struct {
	db *gorm.DB
}

func (p *postgres) Retrieve(projectID, nextTaskID uint64) (*types.Task, error) {
	t := task{}
	if err := p.db.Where("id >= ? AND project_id = ?", nextTaskID, projectID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "failed to query task, next_task_id %v", nextTaskID)
	}

	messageIDs := []string{}
	if err := json.Unmarshal(t.MessageIDs, &messageIDs); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal task message ids, task_id %v", t.ID)
	}

	ms := []*message{}
	if err := p.db.Where("message_id IN ?", messageIDs).Find(&ms).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to query task messages, task_id %v", t.ID)
	}
	if len(ms) == 0 {
		return nil, errors.Errorf("invalid task, task_id %v", t.ID)
	}

	ds := [][]byte{}
	for _, m := range ms {
		ds = append(ds, m.Data)
	}

	return &types.Task{
		ID:             uint64(t.ID),
		ProjectID:      ms[0].ProjectID,
		ProjectVersion: ms[0].ProjectVersion,
		Data:           ds,
		ClientID:       ms[0].ClientID,
		Signature:      t.Signature,
	}, nil
}

func NewPostgres(dsn string) (Datasource, error) {
	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	return &postgres{db}, nil
}
