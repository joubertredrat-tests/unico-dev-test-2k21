package repository

import (
	"database/sql"
	"errors"
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
		fmt.Println("Insert id", id)
		return nil, err
	}

	return &openMarket, nil
}

func (r OpenMarketRepositoryMysql) Update(openMarket domainEntity.OpenMarket) (*domainEntity.OpenMarket, error) {
	query := `UPDATE open_markets SET
		name = ?,
		latitude = ?,
		longitude = ?,
		set_cens = ?,
		area_p = ?,
		address_street = ?,
		address_number = ?,
		address_neighborhood = ?,
		address_reference = ?,
		district_code = ?,
		district_name = ?,
		sub_city_hall_code = ?,
		sub_city_hall_name = ?,
		sub_city_hall_region5 = ?,
		sub_city_hall_region8 = ?,
		updated_at = ?
	WHERE registry_id = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	updatedAt := time.Now()
	openMarket.UpdatedAt = &updatedAt

	res, err := stmt.Exec(
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
		openMarket.UpdatedAt.Format("2006-01-02 15:04:05"),
		openMarket.RegistryID,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domainRepository.OpenMarketRepositoryNotFoundError
	}

	if rowsAffected > 1 {
		return nil, domainRepository.OpenMarketRepositoryHoustonError
	}

	return &openMarket, nil
}

func (r OpenMarketRepositoryMysql) Delete(RegistryID string) error {
	stmt, err := r.db.Prepare("DELETE FROM open_markets WHERE registry_id = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(RegistryID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domainRepository.OpenMarketRepositoryNotFoundError
	}

	if rowsAffected > 1 {
		return domainRepository.OpenMarketRepositoryHoustonError
	}

	return nil
}

func (r OpenMarketRepositoryMysql) GetByRegistryID(RegistryID string) (*domainEntity.OpenMarket, error) {
	query := `SELECT
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
		created_at,
		updated_at
	FROM open_markets WHERE registry_id = ?`
	row := r.db.QueryRow(query, RegistryID)

	var openMarket domainEntity.OpenMarket
	var address domainEntity.Address
	var district domainEntity.District
	var subCityHall domainEntity.SubCityHall

	err := row.Scan(
		&openMarket.Name,
		&openMarket.Latitude,
		&openMarket.Longitude,
		&openMarket.SetCens,
		&openMarket.AreaP,
		&address.Street,
		&address.Number,
		&address.Neighborhood,
		&openMarket.AddressReference,
		&district.Code,
		&district.Name,
		&subCityHall.Code,
		&subCityHall.Name,
		&subCityHall.Region5,
		&subCityHall.Region8,
		&openMarket.CreatedAt,
		&openMarket.UpdatedAt,
	)

	openMarket.RegistryID = RegistryID
	openMarket.Address = address
	openMarket.District = district
	openMarket.SubCityHall = subCityHall

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainRepository.OpenMarketRepositoryNotFoundError
		}

		return nil, err
	}

	return &openMarket, nil
}

func (r OpenMarketRepositoryMysql) GetListByCriteria(searchCriteria domainEntity.OpenMarketSearchCriteria) ([]*domainEntity.OpenMarket, error) {
	getWhereCriteria := func(searchCriteria domainEntity.OpenMarketSearchCriteria) string {
		// where := []string{}

		// if searchCriteria.DistrictName != "" {
		// 	where = append(where, fmt.Sprintf(`district_name LIKE "%%s%"`, searchCriteria.DistrictName))
		// }

		// if searchCriteria.SubCityHallRegion5 != "" {
		// 	where = append(where, fmt.Sprintf(`sub_city_hall_region5 LIKE "%%"`, searchCriteria.SubCityHallRegion5))
		// }

		// if searchCriteria.OpenMarketName != "" {
		// 	where = append(where, fmt.Sprintf(`name LIKE "%%"`, searchCriteria.OpenMarketName))
		// }

		// if searchCriteria.AddressNeighborhood != "" {
		// 	where = append(where, fmt.Sprintf(`address_neighborhood LIKE "%%"`, searchCriteria.AddressNeighborhood))
		// }

		// if len(where) > 0 {
		// 	return fmt.Sprintf(" WHERE %s", strings.Join(where, " AND "))
		// }

		return ""
	}

	where := getWhereCriteria(searchCriteria)
	query := fmt.Sprintf(`SELECT
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
		created_at,
		updated_at
	FROM open_markets%s`, where)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var openMarketList []*domainEntity.OpenMarket

	for rows.Next() {
		var openMarket domainEntity.OpenMarket
		var address domainEntity.Address
		var district domainEntity.District
		var subCityHall domainEntity.SubCityHall

		rows.Scan(
			&openMarket.RegistryID,
			&openMarket.Name,
			&openMarket.Latitude,
			&openMarket.Longitude,
			&openMarket.SetCens,
			&openMarket.AreaP,
			&address.Street,
			&address.Number,
			&address.Neighborhood,
			&openMarket.AddressReference,
			&district.Code,
			&district.Name,
			&subCityHall.Code,
			&subCityHall.Name,
			&subCityHall.Region5,
			&subCityHall.Region8,
			&openMarket.CreatedAt,
			&openMarket.UpdatedAt,
		)

		openMarket.Address = address
		openMarket.District = district
		openMarket.SubCityHall = subCityHall

		openMarketList = append(openMarketList, &openMarket)
	}

	return openMarketList, nil
}
