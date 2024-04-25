CREATE DATABASE demo;
GO

BACKUP DATABASE demo TO DISK = '/sql-backup/demo.bak' WITH FORMAT, COMPRESSION;
BACKUP DATABASE demo TO DISK = 'N' WITH FORMAT, COMPRESSION;
GO

USE master;
DROP DATABASE demo;
GO

USE master
RESTORE DATABASE demo FROM DISK = '/sql-backup/demo.bak' WITH REPLACE, RECOVERY
GO


SELECT name FROM sys.databases;
GO

USE demo
CREATE TABLE Data (ID INT, NAME NVARCHAR(255), AGE INT);
INSERT INTO Data(ID, Name, Age) VALUES (4, 'Anisur Rahman', 30);

INSERT INTO Data(ID, Name, Age) VALUES (1, 'John Doe', 25), (2, 'Jane Smith', 30), (3, 'Bob Johnson', 22);




USE demo
select name FROM sys.tables;

USE demo
Select * from data;

USE demo
DROP table data;

USE master
DROP DATABASE demo;




SELECT database_name, backup_size, compressed_backup_size,
backup_size/compressed_backup_size AS CompressedRatio
FROM msdb..backupset; 



SELECT * FROM msdb.dbo.backupset;








USE dummy
GO
EXEC sp_spaceused @updateusage = N'TRUE'
GO





SELECT bs.database_name,
    backuptype = CASE 
        WHEN bs.type = 'D' AND bs.is_copy_only = 0 THEN 'Full Database'
        WHEN bs.type = 'D' AND bs.is_copy_only = 1 THEN 'Full Copy-Only Database'
        WHEN bs.type = 'I' THEN 'Differential database backup'
        WHEN bs.type = 'L' THEN 'Transaction Log'
        WHEN bs.type = 'F' THEN 'File or filegroup'
        WHEN bs.type = 'G' THEN 'Differential file'
        WHEN bs.type = 'P' THEN 'Partial'
        WHEN bs.type = 'Q' THEN 'Differential partial'
        END + ' Backup',
    CASE bf.device_type
        WHEN 2 THEN 'Disk'
        WHEN 5 THEN 'Tape'
        WHEN 7 THEN 'Virtual device'
        WHEN 9 THEN 'Azure Storage'
        WHEN 105 THEN 'A permanent backup device'
        ELSE 'Other Device'
        END AS DeviceType,
    bms.software_name AS backup_software,
    bs.recovery_model,
    bs.compatibility_level,
    BackupStartDate = bs.Backup_Start_Date,
    BackupFinishDate = bs.Backup_Finish_Date,
    LatestBackupLocation = bf.physical_device_name,
    backup_size_mb = CONVERT(DECIMAL(10, 2), bs.backup_size / 1024. / 1024.),
    compressed_backup_size_mb = CONVERT(DECIMAL(10, 2), bs.compressed_backup_size / 1024. / 1024.),
    database_backup_lsn, -- For tlog and differential backups, this is the checkpoint_lsn of the FULL backup it is based on.
    checkpoint_lsn,
    begins_log_chain,
    bms.is_password_protected
FROM msdb.dbo.backupset bs
LEFT JOIN msdb.dbo.backupmediafamily bf
    ON bs.[media_set_id] = bf.[media_set_id]
INNER JOIN msdb.dbo.backupmediaset bms
    ON bs.[media_set_id] = bms.[media_set_id]
WHERE bs.backup_start_date > DATEADD(MONTH, - 2, sysdatetime()) --only look at last two months
ORDER BY bs.database_name ASC,
    bs.Backup_Start_Date DESC;




SELECT
    backup_size / (1024 * 1024) AS backup_size_MB
FROM
    msdb..backupset;
