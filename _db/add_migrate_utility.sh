#! /bin/sh

: ${INSTABOT_APP:?"INSTABOT_APP is required"}
: ${INSTABOT_HOST:?"INSTABOT_HOST is requried"}

# FOR EACH FILE IN ./migrate_utility (these are necessary for migrations)
for file in $(find ./_db/migrate_utility -name "*.sql" -type f);
do
    ssh dokku@${INSTABOT_HOST} postgres:connect ${INSTABOT_APP} < ${file}
done