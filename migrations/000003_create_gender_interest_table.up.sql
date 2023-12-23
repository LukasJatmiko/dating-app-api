CREATE TABLE IF NOT EXISTS gender_interest(
   id serial PRIMARY KEY,
   user_id int NOT NULL,
   gender_id int NOT NULL,
   created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
   updated_at TIMESTAMP WITHOUT TIME ZONE,
   deleted_at TIMESTAMP WITHOUT TIME ZONE
);