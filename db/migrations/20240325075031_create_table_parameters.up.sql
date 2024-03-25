CREATE TABLE parameters(
    id INT NOT NULL AUTO_INCREMENT,
    device_id INT NOT NULL,
    parameter_name VARCHAR(100) NOT NULL,
    type_param VARCHAR(50) NOT NULL,
    good VARCHAR(100) NOT NULL,
    bad VARCHAR(100) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (device_id) REFERENCES devices(id)
)ENGINE= InnoDB;