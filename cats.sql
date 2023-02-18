CREATE TABLE "cats"(
                       "id" BIGSERIAL PRIMARY KEY,
                       "name" VARCHAR NOT NULL,
                       "is_strip" BOOLEAN NOT NULL DEFAULT FALSE,
                       "color" VARCHAR
)
