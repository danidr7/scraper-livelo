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

	// Buscar lista com todos os partners na API da LIVELO
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

	// Buscar partners na base de dados
	savedPartners, err := mongoConfig.GetAllPartners(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newPartners, err := checkNewPartners(savedPartners, partners)

	if len(newPartners) == 0 {
		log.Println("No partners to add!")
		return
	}

	log.Printf("%d partners to add!", len(newPartners))

	// Salvar novos partners
	err = mongoConfig.SavePartners(ctx, partners)
	if err != nil {
		log.Fatal(err)
	}


	// Iterar lista de todos os partners e consultar API da LIVELO pelo código do partner
	// Consultar em chunks


	// Salvar pontuação do dia
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