package couch

import (
  cloudant "github.com/IBM-Bluemix/go-cloudant"
)

func Db(dbName string) (db *cloudant.DB, err error) {

  // Get the credentials from config
  client, err := cloudant.NewClient(YOUR_USERNAME, YOUR_PASSWORD) // I hardcode these for now
  if err != nil {
    return db, err
  }

  return client.DB(dbName), nil
}