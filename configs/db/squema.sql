CREATE TABLE coffee
(
    id                UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    specie            VARCHAR(50) NOT NULL,
    owner             VARCHAR(100) NOT NULL,
    country_of_origin VARCHAR(100) NOT NULL,
    company           VARCHAR(100),
    aroma             NUMERIC(2, 2) NOT NULL,
    flavor            NUMERIC(2, 2) NOT NULL,
    aftertaste        NUMERIC(2, 2) NOT NULL,
    acidity           NUMERIC(2, 2) NOT NULL,
    body              NUMERIC(2, 2) NOT NULL,
    sweetness         NUMERIC(2, 2) NOT NULL
);