USE dummy;
CREATE Table tblAuthors
(
   Id int identity primary key,
   Author_name nvarchar(50),
   country nvarchar(50)
)
CREATE Table tblBooks
(
   Id int identity primary key,
   Auhthor_id int foreign key references tblAuthors(Id),
   Price int,
   Edition int
)

Declare @Id int
Set @Id = 1

While @Id <= 10000
Begin 
   Insert Into tblAuthors values ('Author - ' + CAST(@Id as nvarchar(10)),
              'Country - ' + CAST(@Id as nvarchar(10)) + ' name')
   Print @Id
   Set @Id = @Id + 1
End


-- Create Database dummy;



USE dummy;
GO
EXEC sp_spaceused @updateusage = N'TRUE';
GO


USE dummy;  
GO
SELECT (total_log_size_in_bytes - used_log_space_in_bytes)*1.0/1024/1024 AS [free log space in MB]  
FROM sys.dm_db_log_space_usage;


BACKUP DATABASE dummy TO DISK = '/sql-backup/dummy.bak';
BACKUP DATABASE dummy TO DISK = 'N';
GO

use master
RESTORE DATABASE dummy FROM DISK = '/sql-backup/dummy.bak';


DROP database dummy;


use dummy;
select * FROM sys.dm_db_log_space_usage;

use dummy;
SELECT SUM(used_log_space_in_bytes) FROM sys.dm_db_log_space_usage;

use dummy;
SELECT SUM(allocated_extent_page_count)*8*1024 FROM sys.dm_db_file_space_usage;


SELECT name from sys.databases;



SELECT 1;