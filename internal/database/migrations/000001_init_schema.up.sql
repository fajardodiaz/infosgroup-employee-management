CREATE TABLE "state" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "position" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "project" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "team" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "employee" (
  "id" bigserial PRIMARY KEY,
  "employee_cod" varchar(25) UNIQUE NOT NULL,
  "full_name" varchar(125) NOT NULL,
  "birth" timestamp NOT NULL,
  "ingress_date" timestamp NOT NULL,
  "end_evaluation_date" timestamp NOT NULL,
  "phone" varchar(15),
  "gender" varchar(1),
  "created_at" timestamp DEFAULT (now()),
  "state_id" integer,
  "position_id" integer,
  "team_id" integer
);

CREATE INDEX ON "employee" ("id", "employee_cod", "full_name");

ALTER TABLE "employee" ADD FOREIGN KEY ("state_id") REFERENCES "state" ("id");

ALTER TABLE "employee" ADD FOREIGN KEY ("position_id") REFERENCES "position" ("id");

ALTER TABLE "employee" ADD FOREIGN KEY ("team_id") REFERENCES "team" ("id");

CREATE TABLE "employee_project" (
  "employee_id" bigserial,
  "project_id" bigserial,
  PRIMARY KEY ("employee_id", "project_id")
);

ALTER TABLE "employee_project" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

ALTER TABLE "employee_project" ADD FOREIGN KEY ("project_id") REFERENCES "project" ("id");
