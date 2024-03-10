package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Database) ExecuteTransaction(ctx context.Context, fn func(session mongo.SessionContext) error) error {
	session, err := d.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(DefaultContext())

	err = session.StartTransaction()
	if err != nil {
		return err
	}

	err = mongo.WithSession(ctx, session, func(session mongo.SessionContext) error {
		err := fn(session)
		if err != nil {
			session.AbortTransaction(session)
			return err
		}

		return session.CommitTransaction(session)
	})

	return err
}
