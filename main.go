package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	"io"
	"log"
	"os"
)

const (
	ConnectionStringFormat = "DefaultEndpointsProtocol=http;AccountName=%s;AccountKey=%s;"
	AccountName            = "devstoreaccount1"
	AccountKey             = "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;BlobEndpoint=http://127.0.0.1:10000/devstoreaccount1"
	ContainerName          = "anisur"
)

var ctx = context.Background()

func main() {
	//connectionString()
	//connectAccountKey()
	blobfs()
}
func blobfs() {
	localbucket, err := blob.OpenBucket(ctx, "azblob://anisur?protocol=http&domain=127.0.0.1:10000&localemu=true")
	if err != nil {
		handleError(err)
	}
	defer localbucket.Close()

	r, err := localbucket.NewReader(ctx, "main.go", nil)
	if err != nil {
		handleError(err)
	}
	defer r.Close()
	// Readers also have a limited view of the blob's metadata.
	fmt.Println("Content-Type:", r.ContentType())
	fmt.Println()
	// Copy from the reader to stdout.
	if _, err := io.Copy(os.Stdout, r); err != nil {
		handleError(err)
	}
}
func connectAccountKey() {
	// general-uri:= https://<account>.blob.core.windows.net/
	emulatorUri := "http://127.0.0.1:10000/devstoreaccount1"
	cred, err := azblob.NewSharedKeyCredential("devstoreaccount1", "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==")
	handleError(err)

	azblobClient, err := azblob.NewClientWithSharedKeyCredential(emulatorUri, cred, nil)
	handleError(err)

	readFile(azblobClient, "main.go")
}

func connectionString() {
	connString := fmt.Sprintf(ConnectionStringFormat, AccountName, AccountKey)
	azblobClient, err := azblob.NewClientFromConnectionString(connString, nil)
	if err != nil {
		handleError(err)
	}
	err = createContainer(ctx, azblobClient, ContainerName)
	if err != nil {
		err = fmt.Errorf("can't create container: %w", err)
		handleError(err)
	}
	err = uploadFile(ctx, azblobClient, "main.go", "./main.go")
	if err != nil {
		err = fmt.Errorf("can't Upload : %w", err)
		handleError(err)
	}

}
func createContainer(ctx context.Context, azblobClient *azblob.Client, containerName string) error {
	if _, err := azblobClient.CreateContainer(ctx, containerName, nil); err != nil {
		return fmt.Errorf("can't create containerName '%s': %w", containerName, err)
	}

	return nil
}

func uploadFile(ctx context.Context, azblobClient *azblob.Client, blobName string, file string) (err error) {
	inputFile, err := os.Open(file) //nolint:gosec
	if err != nil {
		return fmt.Errorf("can't open file '%s': %w", file, err)
	}

	defer func() {
		closeErr := inputFile.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	_, err = azblobClient.UploadFile(ctx, ContainerName, blobName, inputFile, nil)
	if err != nil {
		return fmt.Errorf("can't upload file '%s' to containerName '%s': %w", file, ContainerName, err)
	}

	return nil
}

func readFile(azblobClient *azblob.Client, blobName string) {
	// Download the blob's contents and ensure that the download worked properly
	blobDownloadResponse, err := azblobClient.DownloadStream(context.TODO(), ContainerName, blobName, nil)
	handleError(err)

	// Use the bytes.Buffer object to read the downloaded data.
	// RetryReaderOptions has a lot of in-depth tuning abilities, but for the sake of simplicity, we'll omit those here.
	reader := blobDownloadResponse.Body
	downloadData, err := io.ReadAll(reader)
	handleError(err)
	fmt.Println(string(downloadData))
	err = reader.Close()
	if err != nil {
		return
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
