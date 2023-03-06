CREATE TABLE coffee
(
    id                UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    specie            VARCHAR(50) NOT NULL,
    owner             VARCHAR(100) NOT NULL,
    country_of_origin VARCHAR(100) NOT NULL,
    company           VARCHAR(100),
    aroma             float4 NOT NULL,
    flavor            float4 NOT NULL,
    aftertaste        float4 NOT NULL,
    acidity           float4 NOT NULL,
    body              float4 NOT NULL,
    sweetness         float4 NOT NULL
);