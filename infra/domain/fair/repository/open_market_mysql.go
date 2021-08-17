package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	domainEntity "github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
	domainRepository "github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/repository"
)

type OpenMarketRepositoryMysql struct {
	db *sql.DB
}

func NewOpenMarketRepositoryMysql(db *sql.DB) OpenMarketRepositoryMysql {
	return OpenMarketRepositoryMysql{
		db: db,
	}
}

func (r OpenMarketRepositoryMysql) Create(openMarket domainEntity.OpenMarket) (*domainEntity.OpenMarket, error) {
	query := `INSERT INTO open_markets (
		registry_id, 
		name, 
		latitude,
		longitude,
		set_cens,
		area_p,
		address_street,
		address_number,
		address_neighborhood,
		address_reference,
		district_code,
		district_name,
		sub_city_hall_code,
		sub_city_hall_name,
		sub_city_hall_region5,
		sub_city_hall_region8,
		created_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	createdAt := time.Now()
	openMarket.CreatedAt = &createdAt

	res, err := stmt.Exec(
		openMarket.RegistryID,
		openMarket.Name,
		openMarket.Latitude,
		openMarket.Longitude,
		openMarket.SetCens,
		openMarket.AreaP,
		openMarket.Address.Street,
		openMarket.Address.Number,
		openMarket.Address.Neighborhood,
		openMarket.AddressReference,
		openMarket.District.Code,
		openMarket.District.Name,
		openMarket.SubCityHall.Code,
		openMarket.SubCityHall.Name,
		openMarket.SubCityHall.Region5,
		openMarket.SubCityHall.Region8,
		openMarket.CreatedAt.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			return nil, err
		}
		if me.Number == 1062 {
			return nil, domainRepository.OpenMarketRepositoryAlreadyExistError
		}

		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	fmt.Println("Insert id", id)

	return &openMarket, nil
}

func (r OpenMarketRepositoryMysql) Update(openMarket domainEntity.OpenMarket) (*domainEntity.OpenMarket, error) {
	return nil, nil
}

func (r OpenMarketRepositoryMysql) Delete(RegistryID string) error {
	return nil
}

func (r OpenMarketRepositoryMysql) GetByRegistryID(RegistryID string) (*domainEntity.OpenMarket, error) {
	return nil, nil
}

func (r OpenMarketRepositoryMysql) GetListByCriteria(searchCriteria domainEntity.OpenMarketSearchCriteria) ([]*domainEntity.OpenMarket, error) {
	return nil, nil
}
