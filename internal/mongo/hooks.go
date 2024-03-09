package mongo

type BeforeInsertHook interface {
	BeforeInsert() error
}

type BeforeUpdateHook interface {
	BeforeUpdate() error
}
