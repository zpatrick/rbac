package entry

type Store interface {
	GetEntry(entryID string) (Entry, error)
}
