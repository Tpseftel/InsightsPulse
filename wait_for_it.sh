#!/usr/bin/env sh

set -e

# Extract the host from the first argument
host="$1" 
shift
cmd="$@"

# INFO: This script waits for MariaDB to be up and running before executing the provided command
until mariadb -h "$host" -P "$DB_PORT" -u "$DB_USER" -p"$DB_PASSWORD" -e "SELECT 1" >/dev/null 2>&1; do
  >&2 echo "MariaDB is unavailable - sleeping"
  sleep 1
done

>&2 echo "MariaDB is up - executing command"
exec $cmd