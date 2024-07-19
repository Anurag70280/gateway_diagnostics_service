CREATE TABLE IF NOT EXISTS users(
   "id" serial PRIMARY KEY,
   "name" VARCHAR(48) NOT NULL,
   "email" VARCHAR(48),
   "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);