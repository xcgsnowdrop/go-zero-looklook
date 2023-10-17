# 允许Redis服务远程访问
- 腾讯云服务器防火墙需要设置放开对6379端口的访问限制
- vim /etc/redis.conf && 修改bind 127.0.0.1为bind 0.0.0.0
- 重启redis服务 systemctl restart redis

# 允许mysql远程服务访问
- 腾讯云服务器防火墙需要设置放开对3306端口的访问限制
- vim /etc/my.cnf.d/mysql-server.cnf && 修改bind-address=0.0.0.0
- 重启mysql服务 systemctl restart mysqld

# Mysql导出
## 导出数据库:
- mysqldump -u 用户名 -p --databases 数据库名 > 导出文件名.sql
- mysqldump -u root -p --databases looklook_travel > looklook_travel.sql
## 导出数据表：
- mysqldump -u 用户名 -p 数据库名 表名 > 导出文件名.sql
- mysqldump -u root -p looklook_travel homestay > homestay.sql
