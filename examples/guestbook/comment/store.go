package comment

type Store interface {
	GetComment(commentID string) (Comment, error)
}
