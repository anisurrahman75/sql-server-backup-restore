CREATE CREDENTIAL [https://backup.local/folder]
WITH IDENTITY='SHARED ACCESS SIGNATURE', SECRET = 'does_not_matter'

CREATE CREDENTIAL [https://backup.local/basebackups_005]
WITH IDENTITY='SHARED ACCESS SIGNATURE', SECRET = 'does_not_matter';

CREATE CREDENTIAL [https://backup.local/wal_005]
WITH IDENTITY='SHARED ACCESS SIGNATURE', SECRET = 'does_not_matter';