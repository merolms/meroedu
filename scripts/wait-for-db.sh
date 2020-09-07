#!/bin/sh
until docker container exec -it db mariadb-admin ping -P 3306 -uroot -proot | grep "mysqld is alive" ; do
  >&2 echo "Database is unavailable - waiting for it... 😴"
  sleep 1
done