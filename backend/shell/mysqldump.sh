#!/bin/zsh

source ../.env.development
DATETIME=$(date +%Y%m)

docker container exec -it "$DB_HOST" bash -c "mysqldump -h $DB_HOST -P $PORT -u $DB_USER -p$DB_PASSWORD --no-tablespaces $DB_NAME > ./var/lib/mysql/$DATETIME.sql"
docker cp "$DB_HOST":/var/lib/mysql/"$DATETIME".sql "$PROJECT_PATH"/backend/dbdata/
docker container exec "$DB_HOST" rm /var/lib/mysql/"$DATETIME".sql


