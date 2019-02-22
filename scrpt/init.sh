mysql -uroot -p123456 -Dyoawo < merchant.sql
mysql -uroot -p123456 -Dyoawo < staff.sql
mysql -uroot -p123456 -Dyoawo < transaction_foot.sql
mysql -uroot -p123456 -Dyoawo < withdrawal_foot.sql
mysql -uroot -p123456 -Dyoawo < withdrawal_account.sql

~/install/redis-5.0.3/src/redis-cli -a 123456 -n 1 flushdb
