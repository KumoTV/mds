package models

import "github.com/google/uuid"

type MetadataStore interface {
	ImportMetadata(m VodMetadata) uuid.UUID
	GetMetadataById(metadata_id uuid.UUID) VodMetadata
}
