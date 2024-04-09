### Backup and Restore SQL Server 2022 to S3 Object Storage using MinIO with Docker Compose

**First Start Minio Container**

```bash
docker compose -f minio-s3.yaml up -d minio
```
**Extract IP address of the MinIO container**

We need to get the IP address of the MinIO container so we can use it when we create the SQL Server container and the TLS certificate.

```bash
set MINIO_IP (docker container inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'  minio)
```

**Creating Self-Signed Certificate**

```bash
envsubst < ./certs/template.cnf | tee ./certs/openssl.cnf
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout ./certs/private.key -out ./certs/public.crt -config ./certs/openssl.cnf
```

**Configure MinIO to use the Self-Signed Certificate**

Copy the key and the certificate into known MinIO locations for the client and server and restart the container.
MinIO looks in `/root` dir for these files,

```bash
docker cp ./certs/public.crt  minio:/root/.minio/certs/
docker cp ./certs/private.key minio:/root/.minio/certs/
docker stop minio && docker start minio
```

**Start SQL Server Container**
```bash
docker compose -f minio-s3.yaml up sql-server
```
**Create Bucket**

Access the Minio UI console at `https://MINIO_IP:9000` using `anisur` and `anisur123` credentials. Create a bucket named `stashqa` after logging in.

**Backup & Restore Database**
```bash
sqlcmd -S localhost -U SA -P S0methingS@Str0ng! -No -Q "CREATE CREDENTIAL [s3://s3.example.com:9000/stashqa] WITH IDENTITY = 'S3 Access Key', SECRET = 'anisur:anisur123'"
sqlcmd -S localhost -U SA -P S0methingS@Str0ng! -No -Q "CREATE DATABASE TestDB1"
sqlcmd -S localhost -U SA -P S0methingS@Str0ng! -No -Q "BACKUP DATABASE TestDB1 TO URL = 's3://s3.example.com:9000/stashqa/TestDB1.bak' WITH COMPRESSION, STATS = 10, FORMAT, INIT;"
sqlcmd -S localhost -U SA -P S0methingS@Str0ng! -No -Q "RESTORE DATABASE TestDB1 FROM URL = 's3://s3.example.com:9000/stashqa/TestDB1.bak' WITH FILE = 1, RECOVERY, REPLACE;"
```

**Resources**
- https://learn.microsoft.com/en-us/sql/relational-databases/tutorial-sql-server-backup-and-restore-to-s3?view=sql-server-ver16&tabs=SSMS
- https://learn.microsoft.com/en-us/sql/relational-databases/backup-restore/sql-server-backup-and-restore-with-s3-compatible-object-storage?view=sql-server-ver16#prerequisites-for-the-s3-endpoint
- https://learn.microsoft.com/en-us/sql/relational-databases/backup-restore/sql-server-backup-to-url-s3-compatible-object-storage-best-practices-and-troubleshooting?view=sql-server-ver16&source=recommendations
- https://www.nocentino.com/posts/2022-06-10-setting-up-minio-for-sqlserver-object-storage/