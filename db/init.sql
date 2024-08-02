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