CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   nickname VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   password VARCHAR (300) NOT NULL,
   profile_picture_url varchar(500) NULL,
   birthday DATE NOT NULL,
   gender_id int not null,
   created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
   updated_at TIMESTAMP WITHOUT TIME ZONE,
   deleted_at TIMESTAMP WITHOUT TIME ZONE
);