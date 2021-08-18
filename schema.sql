CREATE TABLE open_markets (
    id int(11) NOT NULL,
    registry_id varchar(10) NOT NULL,
    name text NOT NULL,
    latitude varchar(15) NOT NULL,
    longitude varchar(15) NOT NULL,
    set_cens varchar(30) NOT NULL,
    area_p varchar(30) NOT NULL,
    address_street text NOT NULL,
    address_number text NOT NULL,
    address_neighborhood text NOT NULL,
    address_reference text NOT NULL,
    district_code varchar(10) NOT NULL,
    district_name varchar(150) NOT NULL,
    sub_city_hall_code varchar(10) NOT NULL,
    sub_city_hall_name varchar(255) NOT NULL,
    sub_city_hall_region5 varchar(30) NOT NULL,
    sub_city_hall_region8 varchar(30) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE open_markets
    ADD PRIMARY KEY (id),
    ADD UNIQUE KEY registry_id_unique (registry_id) USING BTREE;

ALTER TABLE open_markets MODIFY id int(11) NOT NULL AUTO_INCREMENT;