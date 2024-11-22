package entity

import "github.com/google/uuid"


type VoteKind string

var UpVote VoteKind = "upvote"
var DownVote VoteKind = "downvote"

type Vote struct {
	ID uuid.UUID
	PostID uuid.UUID
	VoterID uuid.UUID
	Kind VoteKind
}