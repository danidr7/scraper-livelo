package internal

import (
	"github.com/pontuando/scraper-livelo/internal/pkg/config"
	"github.com/pontuando/scraper-livelo/internal/pkg/model"
	"github.com/pontuando/scraper-livelo/internal/app/livelo/service"
	"github.com/pontuando/scraper-livelo/internal/app/livelo/data"
	"log"
	"context"
	"time"
)

func RunScraper() {
	env := config.InitEnvs()

	serviceConfig := service.NewServiceConfig(env.LiveloPartnersURL)
	
	partners, err := serviceConfig.GetPartners()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("len(partners)")
	log.Println(len(partners))

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	mongoConfig := data.NewMongoConfig(env.MongoURL)

	savedPartners, err := mongoConfig.GetAllPartners(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newPartners, err := checkNewPartners(savedPartners, partners)
	if err != nil {
		log.Fatalf("fails attempting check new partners: %s", err)
	}

	if len(newPartners) == 0 {
		log.Println("No partners to add!")
		return
	}

	log.Printf("%d partners to add!", len(newPartners))

	err = mongoConfig.SavePartners(ctx, partners)
	if err != nil {
		log.Fatal(err)
	}


	// TODO
	// - Iterates over all partners and get offer information through the partner ID
	// 		- The fetch can be done in chunks
	// - save the points/score of all partners
}

func checkNewPartners(saved []model.PartnerTreated, toVerify []model.PartnerTreated) ([]model.PartnerTreated, error) {
	
	news := make([]model.PartnerTreated, 0)
	for _, v := range toVerify {
		exists := false
		for _, vv := range saved {
			if v.Code == vv.Code {
				exists = true
				break
			}
		}
		if !exists {
			news = append(news, v)
		}
	}

	return news, nil
}