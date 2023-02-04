CREATE TABLE "cats"(
                       "id" BIGSERIAL PRIMARY KEY,
                       "name" VARCHAR NOT NULL,
                       "IS_STRIP" BOOLEAN NOT NULL DEFAULT FALSE,
                       "COLOR" VARCHAR
)
