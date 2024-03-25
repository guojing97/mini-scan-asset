CREATE TABLE devices(
    id INT NOT NULL AUTO_INCREMENT,
    name_device VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
)ENGINE= InnoDB;