CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR NOT NULL,
    "id_card_file" VARCHAR NOT NULL,
    "id_card_number" VARCHAR NOT NULL,
    "id_card_laser_number" VARCHAR NOT NULL,
    "date_of_birth" TIMESTAMP(3) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

