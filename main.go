package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	_ "gocloud.dev/blob/azureblob"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

func main() {
	//blobfs()
	//connectionString()
	//connectAccountKey()
	sasUrl()
}

func blobfs() {

}

func connectAccountKey() {
	// general-uri:= https://<account>.blob.core.windows.net/
	emulatorUri := "https://account1.blob.localhost:10000/"
	cred, err := azblob.NewSharedKeyCredential("account1", "key1")
	handleError(err)

	azblobClient, err := azblob.NewClientWithSharedKeyCredential(emulatorUri, cred, nil)
	handleError(err)

	containerName := "stashqa"
	err = createContainer(ctx, azblobClient, containerName)
	if err != nil {
		err = fmt.Errorf("can't create container: %w", err)
		handleError(err)
	}

	data := []byte("\nHello, world! This is a blob.\n")
	blobName := "sample-blob.txt"
	// Upload to data to blob storage
	fmt.Printf("Uploading a blob named %s\n", blobName)
	_, err = azblobClient.UploadBuffer(ctx, containerName, blobName, data, &azblob.UploadBufferOptions{})
	handleError(err)

	listBlob(azblobClient, containerName)

}

func connectionString() {
	containerName := "stashqa"
	azblobClient, err := azblob.NewClientFromConnectionString("DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;BlobEndpoint=https://azurite.example.com:10000/devstoreaccount1;", nil)
	//azblobClient, err := azblob.NewClientFromConnectionString("DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;BlobEndpoint=https://127.0.0.1:10000/devstoreaccount1;", nil)
	if err != nil {
		handleError(err)
	}
	//err = createContainer(ctx, azblobClient, containerName)
	//if err != nil {
	//	err = fmt.Errorf("can't create container: %w", err)
	//	handleError(err)
	//}
	//
	//data := []byte("\nHello, world! This is a blob.\n")
	//blobName := "sample-blob.txt"
	//
	//// Upload to data to blob storage
	//fmt.Printf("Uploading a blob named %s\n", blobName)
	//_, err = azblobClient.UploadBuffer(ctx, containerName, blobName, data, &azblob.UploadBufferOptions{})
	//handleError(err)
	listBlob(azblobClient, containerName)
}

func createContainer(ctx context.Context, azblobClient *azblob.Client, containerName string) error {
	if _, err := azblobClient.CreateContainer(ctx, containerName, nil); err != nil {
		return fmt.Errorf("can't create containerName '%s': %w", containerName, err)
	}

	return nil
}

func listBlob(client *azblob.Client, containerName string) {
	fmt.Println("Listing the blobs in the container:")

	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}
func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func sasUrl() {
	//emulatorUri := "https://account1.blob.localhost:10000/"
	//emulatorUri := "https://azurite.example.com:10000/devstoreaccount1"
	//demo SAS token:
	//sv=2018-03-28&spr=https%2Chttp&st=2024-04-13T12%3A58%3A47Z&se=2024-04-14T12%3A58%3A47Z&sr=c&sp=rcwl&sig=Oe3PWgcL1cSvwb8PyrrC1X%2Bu%2F4YqH20qxm1Eb9uWtKY%3D
	accountName := os.Getenv("AZURE_ACCOUNT_NAME")
	accountKey := os.Getenv("AZURE_ACCOUNT_KEY")
	containerName := "stashqa"

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)
	sasQueryParams, err := sas.BlobSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		StartTime:     time.Now().UTC(),
		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour),
		Permissions:   to.Ptr(sas.BlobPermissions{Read: true, Create: true, Write: true, Tag: true}).String(),
		ContainerName: containerName,
	}.SignWithSharedKey(credential)
	handleError(err)

	//sasURL := fmt.Sprintf("https://%s.blob.core.localhost/?%s", accountName, sasQueryParams.Encode())
	fmt.Println(sasQueryParams.Encode())

}
