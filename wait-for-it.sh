#!/usr/bin/env sh
# wait-for-it.sh

set -e

host="$1"
shift
cmd="$@"

# INFO: This script is used to wait for the MariaDB to be up and running before executing the command
until mariadb -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" -e "SELECT 1" >/dev/null 2>&1; do
  >&2 echo "MariaDB is unavailable - sleeping"
  echo "host: $DB_HOST", "user: $DB_USER,  password: $DB_PASSWORD"
  sleep 1
done

>&2 echo "MariaDB is up - executing command"
exec $cmd