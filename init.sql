
-- CREATE USER postgres;

CREATE DATABASE auth_svc;
GRANT ALL PRIVILEGES ON DATABASE auth_svc TO postgres;

CREATE DATABASE product_svc;
GRANT ALL PRIVILEGES ON DATABASE product_svc TO postgres;

CREATE DATABASE order_svc;
GRANT ALL PRIVILEGES ON DATABASE order_svc TO postgres;