BACKUP DATABASE dummy TO DISK = '/sql-backup/dummy_cmp.bak' WITH FORMAT, COMPRESSION;
BACKUP DATABASE dummy TO DISK = 'N';
GO



SELECT
    CONVERT(DECIMAL(10, 2), backup_size / 1024. / 1024.) AS backup_size_MB,
    CONVERT(DECIMAL(10, 2), compressed_backup_size / 1024. / 1024.) AS compressed_size_MB,
    type AS backup_type,
    backup_finish_date AS completion_time
FROM
    msdb..backupset
WHERE
    database_name = 'dummy'
ORDER BY
    backup_finish_date DESC;



USE dummy
GO
EXEC sp_spaceused @updateusage = N'TRUE'
GO


USE master;
GO
DROP DATABASE dummy;
GO

select name from sys.databases;
go

