package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	gcloud_blob "gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var ctx = context.Background()

func main() {
	//xxorm()
	blobfs()
	//sasBlob()
	//downloadBlobAndStream()
}

func blobfs() {
	dir, fileName := path.Split("demo-mssql-db-backup-frequent-backup-1714475100-5wpzd/dummy.bak")
	bucket, err := openBucket(ctx, dir)
	localFilePath := "./dummy.bak"
	f, err := os.Create(localFilePath)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n\n", err)
	}
	defer f.Close()
	//if err := bucket.Download(ctx, fileName, f, nil); err != nil {
	//	fmt.Println(err)
	//}

	x, err := bucket.Open(fileName)

	_, err = io.Copy()
	if err != nil {
		panic(err)
	}

	println("Data copied successfully!")
}

func openBucket(ctx context.Context, dir string) (*gcloud_blob.Bucket, error) {
	var bucket *gcloud_blob.Bucket
	var err error

	bucket, err = gcloud_blob.OpenBucket(ctx, "azblob://kubestashqa")
	if err != nil {
		return nil, err
	}

	suffix := strings.Trim(path.Join("sunny", dir), "/") + "/"
	if suffix == string(os.PathSeparator) {
		return bucket, nil
	}
	return gcloud_blob.PrefixedBucket(bucket, suffix), nil
}

func downloadBlobAndStream() {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
	accountKey := os.Getenv("AZURE_STORAGE_KEY")
	//fmt.Println(accountName)
	//fmt.Println(accountKey)

	// Create a blobClient object to a blob in the container (we assume the container & blob already exist).
	blobURL := fmt.Sprintf("https://%s.blob.core.windows.net/stashqa/demo.bak", accountName)
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)
	blobClient, err := blob.NewClientWithSharedKeyCredential(blobURL, credential, nil)
	handleError(err)

	var size int64
	var buf []byte
	size, err = blobClient.DownloadBuffer(ctx, buf, nil)
	if err != nil {
		fmt.Println("azure Blobstore: %w", err)
		handleError(err)
	}
	fmt.Println(size)
}

func sasBlob() {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
	accountKey := os.Getenv("AZURE_STORAGE_KEY")
	fmt.Println(accountName)
	fmt.Println(accountKey)
	containerName := "stashqa"
	_ = containerName
	// Create a credential object using your account name and key
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}

	// Create BlobSASSignatureValues for specifying SAS parameters
	sasQueryParams := sas.AccountSignatureValues{
		Protocol: sas.ProtocolHTTPSandHTTP, // Use HTTPS protocol
		ResourceTypes: to.Ptr(sas.AccountResourceTypes{
			Object: true,
		}).String(),
		ExpiryTime: time.Now().UTC().Add(2 * time.Hour), // SAS expiry time
		Permissions: to.Ptr(sas.AccountPermissions{ // Permissions for the SAS token
			Read:   true,
			Write:  true,
			Create: true,
		}).String(),
	}

	// Generate SAS token with the specified parameters
	queryParams, err := sasQueryParams.SignWithSharedKey(credential)
	if err != nil {
		log.Fatal("Failed to generate SAS token with error: " + err.Error())
	}

	// Encode the SAS token into a string
	sasToken := queryParams.Encode()

	// Print the generated SAS token
	fmt.Println("Generated SAS Token:")
	fmt.Println(sasToken)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
