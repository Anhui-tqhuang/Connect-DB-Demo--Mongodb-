package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	hostname0, hostname1 string
	port0, port1         string
	cafile               string
	authsource           string
	database             string
	username             string
	password             string
)

func init() {
	hostname0 = os.Getenv("HOSTNAME0")
	hostname1 = os.Getenv("HOSTNAME1")
	port0 = os.Getenv("PORT0")
	port1 = os.Getenv("PORT1")
	cafile = os.Getenv("CAFILE")
	authsource = os.Getenv("AUTHSOURCE")
	database = os.Getenv("DATABASE")
	username = authsource
	password = os.Getenv("PASSWORD")
}

const mongoDBConnectionString = "mongodb://%s:%s@%s:%s,%s:%s/%s?authSource=%s&replicaSet=replset"

func main() {
	// Load CA file.
	caCert, err := ioutil.ReadFile(cafile)
	if err != nil {
		glog.Error(err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	url := fmt.Sprintf(mongoDBConnectionString, username, password, hostname0, port0, hostname1, port1, database, authsource)
	opts := options.Client().ApplyURI(url)
	opts.TLSConfig = &tls.Config{
		RootCAs: caCertPool,
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		glog.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		glog.Error(err)
		return
	}

	collection := client.Database(database).Collection("numbers")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel2()

	res, err := collection.InsertOne(ctx2, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		glog.Error(err)
		return
	}
	id := res.InsertedID

	fmt.Println(id)
}
