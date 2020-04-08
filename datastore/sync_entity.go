package datastore

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/brave-experiments/sync-server/sync_pb"
	"github.com/brave-experiments/sync-server/utils"
	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
)

// SyncEntity represents the underline DB schema of sync entities.
type SyncEntity struct {
	ID                     string         `json:"id_string" db:"id"`
	ParentID               sql.NullString `json:"parent_id_string" db:"parent_id"`
	OldParentID            sql.NullString `json:"old_parent_id" db:"old_parent_id"`
	Version                int64          `json:"version" db:"version"`
	Mtime                  int64          `json:"mtime" db:"mtime"`
	Ctime                  int64          `json:"ctime" db:"ctime"`
	Name                   sql.NullString `json:"name" db:"name"`
	NonUniqueName          sql.NullString `json:"non_unique_name" db:"non_unique_name"`
	ServerDefinedUniqueTag sql.NullString `json:"server_defined_unique_tag" db:"server_defined_unique_tag"`
	Deleted                bool           `json:"deleted" db:"deleted"`
	OriginatorCacheGUID    sql.NullString `json:"originator_cache_guid" db:"originator_cache_guid"`
	OriginatorClientItemID sql.NullString `json:"originator_client_item_id" db:"originator_client_item_id"`
	Specifics              []byte         `json:"specifics" db:"specifics"`
	DataTypeID             int            `json:"data_type_id" db:"data_type_id"`
	Folder                 bool           `json:"folder" db:"folder"`
	ClientDefinedUniqueTag sql.NullString `json:"client_defined_unique_tag" db:"client_defined_unique_tag"`
	UniquePosition         []byte         `json:"unique_position" db:"unique_position"`
	ClientID               string         `db:"client_id"`
}

// InsertSyncEntity inserts a new sync entity into postgres database.
func (pg *Postgres) InsertSyncEntity(entity *SyncEntity) error {
	// TODO: Ensure the uniqueness of client_defined_unique_tag here
	stmt := `INSERT INTO sync_entities(id, parent_id, old_parent_id, version, mtime, ctime, name, non_unique_name, server_defined_unique_tag, deleted, originator_cache_guid, originator_client_item_id, specifics, data_type_id, folder, client_defined_unique_tag, unique_position, client_id) VALUES(:id, :parent_id, :old_parent_id, :version, :mtime, :ctime, :name, :non_unique_name, :server_defined_unique_tag, :deleted, :originator_cache_guid, :originator_client_item_id, :specifics, :data_type_id, :folder, :client_defined_unique_tag, :unique_position, :client_id)`
	_, err := pg.NamedExec(stmt, *entity)
	if err != nil {
		fmt.Println("Insert error: ", err.Error())
	}
	return err
}

// UpdateSyncEntity updates a sync entity in postgres database.
func (pg *Postgres) UpdateSyncEntity(entity *SyncEntity) error {
	if entity.Deleted {
		return pg.deleteSyncEntity(entity.ID)
	}

	stmt := `UPDATE sync_entities SET parent_id = :parent_id, old_parent_id = :old_parent_id, version = :version, mtime = :mtime, name = :name, non_unique_name = :non_unique_name, specifics = :specifics, folder = :folder, unique_position = :unique_position WHERE id = :id`
	_, err := pg.NamedExec(stmt, *entity)
	if err != nil {
		fmt.Println("Update error: ", err.Error())
	}
	return err
}

func (pg *Postgres) deleteSyncEntity(id string) error {
	_, err := pg.Exec(`UPDATE sync_entities SET deleted = true WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Delete error: ", err.Error())
	}
	return err
}

// CheckVersion get the sync entry with id saved in the database and checks
// the saved server version against the client version.
func (pg *Postgres) CheckVersion(id string, clientVersion int64) (bool, error) {
	var serverVersion int64
	err := pg.Get(&serverVersion, "SELECT version FROM sync_entities WHERE id = $1", id)
	if err != nil {
		fmt.Println("Get version error: ", err.Error(), "id: ", id)
		return false, nil
	}

	return clientVersion == serverVersion, nil
}

// GetUpdatesForType retrieves sync entities of a data type where it's mtime
// is later than the client token.
func (pg *Postgres) GetUpdatesForType(dataType int32, clientToken int64, fetchFolders bool) (entities []SyncEntity, err error) {
	stmt := "SELECT * FROM sync_entities WHERE data_type_id = $1 AND mtime > $2"
	if !fetchFolders {
		stmt += "AND folder = false"
	}
	stmt += " ORDER BY mtime"
	err = pg.Select(&entities, stmt, dataType, clientToken)
	return
}

// CreateDBSyncEntity converts a protobuf sync entity into a DB sync entity.
func CreateDBSyncEntity(entity *sync_pb.SyncEntity, cacheGUID string) (*SyncEntity, error) {
	var err error
	var specifics []byte
	// if entity.Specifics != nil {  // TODO: make sure this is present in the
	// validator
	specifics, err = proto.Marshal(entity.Specifics)
	if err != nil {
		fmt.Println("Marshal Error", err.Error())
		return nil, err
	}
	// }

	// TODO: wrap getting type ID into an util function
	structField := reflect.ValueOf(entity.Specifics.SpecificsVariant).Elem().Type().Field(0)
	tag := structField.Tag.Get("protobuf")
	s := strings.Split(tag, ",")
	dataTypeID, _ := strconv.Atoi(s[1])

	var uniquePosition []byte
	if entity.UniquePosition != nil {
		uniquePosition, err = proto.Marshal(entity.UniquePosition)
		if err != nil {
			fmt.Println("Marshal Error", err.Error())
			return nil, err
		}
	}

	deleted := false
	if entity.Deleted != nil {
		deleted = *entity.Deleted
	}
	folder := false
	if entity.Folder != nil {
		folder = *entity.Folder
	}

	id := *entity.IdString
	originatorCacheGUID := sql.NullString{Valid: false}
	originatorClientItemID := sql.NullString{Valid: false}
	if len(cacheGUID) > 0 {
		if *entity.Version == 0 {
			id = uuid.NewV4().String()
		}
		originatorCacheGUID = sql.NullString{String: cacheGUID, Valid: true}
		originatorClientItemID = sql.NullString{String: *entity.IdString, Valid: true}
	}

	return &SyncEntity{
		ID:                     id,
		ParentID:               utils.StringOrNull(entity.ParentIdString),
		OldParentID:            utils.StringOrNull(entity.OldParentId),
		Version:                *entity.Version,
		Ctime:                  *entity.Mtime,
		Mtime:                  time.Now().Unix(),
		Name:                   utils.StringOrNull(entity.Name),
		NonUniqueName:          utils.StringOrNull(entity.NonUniqueName),
		ServerDefinedUniqueTag: utils.StringOrNull(entity.ServerDefinedUniqueTag),
		Deleted:                deleted,
		OriginatorCacheGUID:    originatorCacheGUID,
		OriginatorClientItemID: originatorClientItemID,
		ClientDefinedUniqueTag: utils.StringOrNull(entity.ClientDefinedUniqueTag),
		Specifics:              specifics,
		Folder:                 folder,
		UniquePosition:         uniquePosition,
		DataTypeID:             dataTypeID,
		ClientID:               "brave",
	}, nil
}

// CreatePBSyncEntity converts a DB sync entity to a protobuf sync entity.
func CreatePBSyncEntity(entity *SyncEntity) (*sync_pb.SyncEntity, error) {
	pbEntity := &sync_pb.SyncEntity{
		IdString:               &entity.ID,
		ParentIdString:         utils.StringPtr(&entity.ParentID),
		Version:                &entity.Version,
		Mtime:                  &entity.Mtime,
		Ctime:                  &entity.Ctime,
		Name:                   utils.StringPtr(&entity.Name),
		NonUniqueName:          utils.StringPtr(&entity.NonUniqueName),
		ServerDefinedUniqueTag: utils.StringPtr(&entity.ServerDefinedUniqueTag),
		ClientDefinedUniqueTag: utils.StringPtr(&entity.ClientDefinedUniqueTag),
		OriginatorCacheGuid:    utils.StringPtr(&entity.OriginatorCacheGUID),
		OriginatorClientItemId: utils.StringPtr(&entity.OriginatorClientItemID),
		Deleted:                &entity.Deleted,
		Folder:                 &entity.Folder}

	specifics := &sync_pb.EntitySpecifics{}
	err := proto.Unmarshal(entity.Specifics, specifics)
	if err != nil {
		fmt.Println("[CreatePBSyncEntity] Error when unmarshalling specifics:", err.Error())
		return nil, err
	}
	pbEntity.Specifics = specifics

	if entity.UniquePosition != nil {
		uniquePosition := &sync_pb.UniquePosition{}
		err := proto.Unmarshal(entity.UniquePosition, uniquePosition)
		if err != nil {
			fmt.Println("[CreatePBSyncEntity] Error when unmarshalling specifics:", err.Error())
			return nil, err
		}
		pbEntity.UniquePosition = uniquePosition
	}

	return pbEntity, nil
}
