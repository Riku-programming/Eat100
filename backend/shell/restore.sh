#!/bin/zsh

source ../.env
DATETIME=$(date +%Y%m)

docker cp "$PROJECT_PATH"/backend//dbdata/"$DATETIME".sql "$DB_HOST":/var/lib/mysql/"$DATETIME".sql
docker container exec -it "$DB_HOST" bash -c "mysql -h $DB_HOST -P $PORT -u $DB_USER -p$DB_PASSWORD $DB_NAME < /var/lib/mysql/$DATETIME.sql"
docker container exec "$DB_HOST" rm /var/lib/mysql/"$DATETIME".sql
