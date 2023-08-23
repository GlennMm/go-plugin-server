package utils

import "xorm.io/xorm"

func MutateDb[T interface{}](db *xorm.Engine, data *T) error {
	session := db.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}

	if _, err := session.InsertOne(&data); err != nil {
		return err
	}

	return session.Commit()
}
