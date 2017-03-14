package models

import (
	"log"
	"github.com/google/uuid"
)

type VodMetadata struct {
	Id       int      `db:"id"`
	MetadataId uuid.UUID `db:"metadata_id"`
	Metadata Metadata `db:"metadata"`
}

func (db *Db) ImportMetadata(m VodMetadata) uuid.UUID {

	metadata_id, _ := uuid.NewUUID()

	err := db.QueryRow(
		"insert into vod (metadata_id, metadata) values ($1, $2) returning metadata_id",
		metadata_id,
		m.Metadata,
	).Scan(&metadata_id)

	if err != nil {
		log.Fatal(err)
	}
	return metadata_id
}

func (db *Db) GetMetadataById(metadata_id uuid.UUID) VodMetadata {
	m := VodMetadata{}
	err := db.QueryRow(
		"select metadata_id,metadata from vod where metadata_id = ($1)",
		metadata_id,
	).Scan(&m.MetadataId, &m.Metadata)

	if err != nil {
		log.Fatal(err)
		return m
	}
	// return empty Metadata if not found
	return m
}
