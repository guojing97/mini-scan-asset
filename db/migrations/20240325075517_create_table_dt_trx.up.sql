CREATE TABLE dt_trx(
    id INT NOT NULL AUTO_INCREMENT,
    device_id INT NOT NULL,
    notes TEXT NOT NULL,
    is_checked BOOLEAN NOT NULL DEFAULT 0,
    date_checked DATETIME,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
)ENGINE= InnoDB;