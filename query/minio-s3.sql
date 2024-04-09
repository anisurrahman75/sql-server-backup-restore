CREATE CREDENTIAL [s3://s3.example.com:9000/stashqa] WITH IDENTITY = 'S3 Access Key', SECRET = 'anisur:anisur123';
GO

CREATE DATABASE TestDB1;
GO

SELECT name FROM sys.credentials
GO

SELECT name FROM sys.databases
GO

BACKUP DATABASE TestDB1 TO URL = 's3://s3.example.com:9000/stashqa/TestDB1.bak' WITH COMPRESSION, STATS = 10, FORMAT, INIT;
GO

RESTORE DATABASE TestDB1 FROM URL = 's3://s3.example.com:9000/stashqa/TestDB1.bak' WITH FILE = 1, RECOVERY, REPLACE;
GO