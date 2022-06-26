package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pontuando/scraper-livelo/internal/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURL = os.Getenv("MONGO_URL")

func mock(filePath string) *httptest.Server {
	os.Setenv("LIVELO_URL", "http://localhost:5002")
	
	ts := httptest.NewUnstartedServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("---------requested")

		f, err := os.OpenFile(filePath, os.O_RDONLY, 400)
		if err != nil {
			panic(fmt.Sprintf("fails open partners.json: %s", err))
		}

		defer f.Close()

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic(fmt.Sprintf("fails reading partners.json content: %s", err))
		}

		fmt.Fprintf(w, string(content))
    }))

	l, _ := net.Listen("tcp", ":5002")
	ts.Listener = l

	return ts
}


func TestMain(t *testing.T){
	cleanCollection(t)
	
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	
	t.Run("shouldInsertAllPartnersOnFirstTime", shouldInsertAllPartnersOnFirstTime)
	t.Run("testInsertNewPartner", testInsertNewPartner)
}

func shouldInsertAllPartnersOnFirstTime(t *testing.T) {
	ts := mock("../test/partners.json")
	ts.Start()
	defer ts.Close()

	expectedPartnersAmount := 329

	RunScraper()

	ctx := context.Background()

	mgCli, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		t.Errorf("fails attempting connect to mongodb: %s", err)
	}

	collection := mgCli.Database(model.DatabaseName).Collection(model.PartnerCollectionName)
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		t.Errorf("fails attempting find partners: %s", err)
	}

	if count != int64(expectedPartnersAmount) {
		t.Errorf("Expected count [%d] but found [%d]", expectedPartnersAmount, count)
	}
}

func testInsertNewPartner(t *testing.T) {
	ts := mock("../test/one_partner.json")
	ts.Start()
	defer ts.Close()

	RunScraper()
}

func cleanCollection(t *testing.T) {
	ctx := context.Background()

	mgCli, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		t.Errorf("fails attempting connect to mongodb: %s", err)
	}

	collection := mgCli.Database(model.DatabaseName).Collection(model.PartnerCollectionName)
	_, err = collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		t.Errorf("fails attempting find partners: %s", err)
	}
}