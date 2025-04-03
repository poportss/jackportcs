SELECT * FROM "schema_versions" WHERE service = 'auth' AND "schema_versions"."deleted_at" IS NULL ORDER BY version DESC,"schema_versions"."id" LIMIT 1

SHOW search_path;

ALTER ROLE jackportcs_dev SET search_path TO jackportcs_schema;
