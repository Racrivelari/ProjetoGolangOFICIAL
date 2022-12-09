-- SCRIPT PRA MY SQL

drop database if exists deposito;
CREATE database IF NOT EXISTS deposito DEFAULT CHARACTER SET utf8MB4 ;
USE deposito;
#SET SQL_SAFE_UPDATES = 0;

CREATE TABLE IF NOT EXISTS Product (
  id_prod integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name_prod VARCHAR(45) NOT NULL,
  price_prod VARCHAR(45) NOT NULL,
  code_prod VARCHAR(45) NOT NULL,
  created_at datetime NULL DEFAULT NOW(),
  Unique(code_prod));
  
CREATE TABLE IF NOT EXISTS User (
  id_user integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  email_user VARCHAR(45) NOT NULL,
  senha_user VARCHAR(45) NOT NULL,
  Unique(email_user)); 
  
CREATE TABLE IF NOT EXISTS Logs  (
	id_log integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
	updated_at datetime DEFAULT NOW(),
    id_prod INT NOT NULL,
	constraint fk_product_log foreign key (id_prod) references Product(id_prod)
);

INSERT INTO Product (name_prod, price_prod, code_prod) values ("Iphone", "5000", "rcb");
INSERT INTO Product (name_prod, price_prod, code_prod) values ("Galaxy S22", "4000", "cel2");
INSERT INTO Product (name_prod, price_prod, code_prod) values ("Pocophone", "1000", "cel3");
INSERT INTO Product (name_prod, price_prod, code_prod) values ("Windows Phone", "1000", "cel4");

-- insert into Logs (id_prod, updated_at) values (1, current_timestamp());
  
select * from Product;
select * from Logs;


DELIMITER $$

CREATE TRIGGER product_update
AFTER UPDATE
ON Product FOR EACH ROW
BEGIN
        INSERT INTO Logs(id_prod,updated_at)
        VALUES(old.id_prod, current_timestamp());
END$$

DELIMITER ;

-- update Product set price_prod = "1500" where id_prod = 1;
-- update Product set name_prod = "Computador" where id_prod = 1;





  

  
