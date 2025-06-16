[psql]
dbname = "${DB_NAME}"
host   = "${DB_HOST}"
port   = ${DB_PORT}
user   = "${DB_USER}"
pass   = "${DB_PASSWORD}"
sslmode = "disable"
blacklist = ["goose_db_version"]
[psql.schema]
migration_dir = "migrations"
[psql.output]
dir = "internal/repository/models"
pkgname = "models"
