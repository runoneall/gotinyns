package records

type Record struct {
	Name  []byte `db:"name"`
	Type  uint16 `db:"type"`
	Value []byte `db:"value"`
	TTL   uint32 `db:"ttl"`
}
