# sql-server-backup-restore

**Configure SQL Server to Trust Our Self-Signed Certificate**

```shell
cd ./certs
docker cp public.crt sql-server:/usr/local/share/ca-certificates
docker exec -it -u 0 sql-server update-ca-certificates
```
We now need to tell SQL Server to trust the self-signed certificate. We do that by copying the certificate into the known CA location for SQL Server, setting permissions on the files, and then restarting the container. 

This will allow SQL Server to trust the self-signed certificate.
```shell
docker exec -it -u 0 sql-server mkdir /usr/local/share/ca-certificates/mssql-ca-certificates/
docker cp public.crt sql-server:/usr/local/share/ca-certificates/mssql-ca-certificates/
docker exec -it -u 0 sql-server chown 10001:10001 -R /usr/local/share/ca-certificates/mssql-ca-certificates/
docker stop sql-server && docker start sql-server
```

Once you restart the container, you can look in the SQL Server’s error lot to confirm its loads the self-signed certificate. 
Once you see ‘Installing Client TLS certificates to the store,’ your certificate is loaded by SQL Server.

```bash
docker logs sql-server --follow | grep 'Installing Client TLS certificates to the store'
```