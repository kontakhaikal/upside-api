package dto

import "github.com/google/uuid"

type CreatePostRequest struct {
	Body     string
	SideID   uuid.UUID
	AuthorID uuid.UUID
}

type CreatePostResponse struct {
	PostID uuid.UUID `json:"post_id"`
}
