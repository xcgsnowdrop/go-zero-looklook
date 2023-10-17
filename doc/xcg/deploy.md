# 允许Redis服务远程访问
- 腾讯云服务器防火墙需要设置放开对6379端口的访问限制
- vim /etc/redis.conf && 修改bind 127.0.0.1为bind 0.0.0.0
- 重启redis服务 systemctl restart redis

# 允许mysql远程服务访问

# 