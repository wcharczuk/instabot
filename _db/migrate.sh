#! /bin/sh

if [ ! -z "$INSTABOT__APP" ]; then 
    : ${INSTABOT__HOST:?"INSTABOT__HOST is requried"}
    
    for file in $(find ./_db/migrations -name "*.sql" -type f);
    do
        ssh dokku@${INSTABOT__HOST} postgres:connect ${INSTABOT__APP} < ${file}  
    done
else
    : ${DB_SCHEMA:?"DB_SCHEMA is required"}
    
    for file in $(find ./_db/migrations -name "*.sql" -type f);
    do
        psql $DB_SCHEMA -f ${file};
    done
fi