CREATE TABLE customers (
  customer_id INT PRIMARY KEY,
  name VARCHAR(255),
  date_of_birth DATE,
  city VARCHAR(255),
  zipcode VARCHAR(255),
  status BOOLEAN
);

INSERT INTO customers (customer_id,name,date_of_birth,city,zipcode,status)
VALUES
  (2000,"Steve","1978-12-15","Delhi","110075",1),
  (2001,"Arian","1988-05-21","Newburgh, NY","12550",1),
  (2002,"Hadley","1988-04-30","Englewood, NJ","07631",1),
  (2003,"Ben","1988-01-04","Manchester, NH","03102",0),
  (2004,"Nina","1988-05-14","Clarlston, MI", "48348",1),
  (2005,"Osman","1988-11-08","Hyattsvile, MD","20782",0);

CREATE TABLE accounts(
  account_id INT PRIMARY KEY,
  customer_id INT,
  opening_date DATETIME,
  account_type VARCHAR(255),
  amount DECIMAL,
  status BOOLEAN
);

INSERT INTO accounts (account_id,customer_id,opening_date,account_type,amount,status)
VALUES
  (1,2000,"2000-12-15","saving",20000,1),
  (2,2001,"2001-12-15","checking",5000,1),
  (3,2002,"2002-12-15","saving",8000,1),
  (4,2003,"2000-12-15","checking",7500,1),
  (5,2004,"2001-12-15","saving",40250,1),
  (6,2005,"2002-12-15","checking",15000,1);