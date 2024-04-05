/* Delete Credential */

-- DROP CREDENTIAL azure_secret;
-- GO


/* Azurite Emulator */

-- IF NOT EXISTS  
-- (SELECT * FROM sys.credentials   
-- WHERE name = 'azure_secret')  
-- CREATE CREDENTIAL [azure_secret] WITH IDENTITY = 'devstoreaccount1'  
-- ,SECRET = 'Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==';

/* Azure Storage */

-- IF NOT EXISTS  
-- (SELECT * FROM sys.credentials   
-- WHERE name = 'azure_secret')  
-- CREATE CREDENTIAL [azure_secret] WITH IDENTITY = 'stashqa'  
-- ,SECRET = 'VnbR75fFvUwGMFzLUfELju054pcdxqpHgvSdPavikhsI44RrzStNTWkoWmrCrHM/BukG/654mr0Z+ASt8RuwNA==';


/* SAS */

IF NOT EXISTS
(SELECT * FROM sys.credentials
WHERE name = 'https://stashqa.blob.core.windows.net/stashqa')
CREATE CREDENTIAL [https://stashqa.blob.core.windows.net/stashqa]
   WITH IDENTITY = 'SHARED ACCESS SIGNATURE',
   SECRET = '';



/* SAS */

IF NOT EXISTS
(SELECT * FROM sys.credentials
WHERE name = 'http://127.0.0.1:10000/devstoreaccount1/kubestash')
CREATE CREDENTIAL [http://127.0.0.1:10000/devstoreaccount1/kubestash]
   WITH IDENTITY = 'SHARED ACCESS SIGNATURE',
   SECRET = '';

SELECT name FROM sys.credentials
GO