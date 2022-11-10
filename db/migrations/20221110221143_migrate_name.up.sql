CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE articles (
  "id" uuid DEFAULT uuid_generate_v4 (),
  "name" text NOT NULL,
  "description" text NOT NULL
);