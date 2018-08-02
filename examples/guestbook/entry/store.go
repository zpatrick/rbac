package entry

type Store interface {
	ListEntries() ([]Entry, error)
	ReadEntry(entryID string) (Entry, error)
}
