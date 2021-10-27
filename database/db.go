package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "urls"
)

type Provider interface {
	IncreaseKey(key Keys)
	GetDomainByID(id int) (string, error)
	GetAllDomains() ([]Domains, error)
	RemoveDomain(domain Domains)
}

type DataBase struct {
	db *gorm.DB
}

var Instance Provider = newDB()

func newDB() *DataBase {
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Keys{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Domains{})

	if err != nil {
		panic(err)
	}
	initDomainTable(db)
	return &DataBase{db: db}
}

func (d *DataBase) IncreaseKey(key Keys) {
	objFromDB := Keys{}
	err := d.db.Table("keys").Where("user_key = ?", key.UserKey).First(&objFromDB).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Print("did not find record in keys")
			key.VisitCounter++
			err := d.db.Table("keys").Create(&key).Error

			if err != nil {
				log.Print("Error when creating new key on db")

			}
			return
		}
		log.Print("failed to get key from db " + err.Error())
		return
	}

	objFromDB.VisitCounter++

	err = d.db.Table("keys").Where("user_key=?", key.UserKey).Update("visit_counter", objFromDB.VisitCounter).Error
	if err != nil {
		log.Print(fmt.Sprintf("Error when updating counter for key %s on db :%s", objFromDB.UserKey, err.Error()))

	}
}

func (d *DataBase) GetDomainByID(id int) (string, error) {
	objFromDB := Domains{}

	err := d.db.Table("domains").Select("*").Where("id=?", id).First(&objFromDB).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Print("did not find record in domains table for specific id")
			return "", err
		}
		log.Print("failed to get domain from db " + err.Error())
		return "", err

	}

	return objFromDB.Domain, nil
}

func (d *DataBase) GetAllDomains() ([]Domains, error) {
	var objFromDB []Domains

	result := d.db.Find(&objFromDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return objFromDB, nil
}

func (d *DataBase) RemoveDomain(domain Domains) {
	log.Print(fmt.Sprintf("Will remove url: %s from db", domain.Domain))
	d.db.Delete(&domain)
}

//just init the first time - will add data
func initDomainTable(db *gorm.DB) {

	var domains = map[int]string{1: "www.ikea.co.il", 2: "www.yad2.co.il", 3: "www.walla.co.il", 4: "www.clalit.co.il", 5: "www.clalit.co.il/notexistingpage"}

	for i, domain := range domains {

		domainOBj := Domains{uint(i), domain}

		err := db.Table("domains").FirstOrCreate(&domainOBj).Error

		if err != nil {
			log.Print("Error when creating new domain on db")
		}
	}

}
