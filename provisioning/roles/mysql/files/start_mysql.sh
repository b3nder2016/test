CONSULHOST="172.1.1.3"
CONSULPORT="8500"

MYSQLHOST="172.1.1.2"
MYSQLPORT="3306"
MYSQLUSER="root"
MYSQLPASS="test"

curl -X PUT -d @- $CONSULHOST:8500/v1/kv/mysqlhost <<< $MYSQLHOST
curl -X PUT -d @- $CONSULHOST:8500/v1/kv/mysqlport <<< $MYSQLPORT
curl -X PUT -d @- $CONSULHOST:8500/v1/kv/mysqluser <<< $MYSQLUSER
curl -X PUT -d @- $CONSULHOST:8500/v1/kv/mysqlpass <<< $MYSQLPASS

exec /usr/sbin/mysqld
