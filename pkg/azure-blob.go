package pkg

import (
	"context"
	"errors"
	"net/url"
	"fmt"
	"io/ioutil"
	"github.com/Azure/azure-storage-blob-go/azblob"
)

type AzureBlob struct {
	sourceFilePath          *string
	storageAccountName      *string
	storageAccountAccessKey *string
	blobContainerPath       *string
	blobContainerURL        azblob.ContainerURL
}

func (b *AzureBlob) HttpPut() error {
	// Method Vars
	var err error

	f, err := ioutil.ReadFile(*b.sourceFilePath)
	if err != nil {
		return err
	}

	ctx := context.Background()

	blobURL := b.blobContainerURL.NewBlockBlobURL(*b.sourceFilePath)

	_, err = azblob.UploadBufferToBlockBlob(ctx, f, blobURL, azblob.UploadToBlockBlobOptions{Parallelism: 16})
	if err != nil {
		return err
	}

	return nil
}

func NewAzureBlob(config map[string]*string) (*AzureBlob, error) {
	var err error

	if *config["TargetFilePath"] == "" {
		return nil,errors.New("Unable to find property StorageAccountName")
	}

	if *config["StorageAccountName"] == "" {
		return nil,errors.New("Unable to find property StorageAccountName")
	}

	if *config["StorageAccountAccessKey"] == "" {
		return nil,errors.New("Unable to find property StorageAccountName")
	}

	if *config["BlobContainerPath"] == "" {
		return nil,errors.New("Unable to find property StorageAccountName")
	}

	credential, err := azblob.NewSharedKeyCredential(*config["StorageAccountName"], *config["StorageAccountAccessKey"])
	if err != nil {
		return nil,err
	}
	azurePipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	u, err := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", *config["StorageAccountName"], *config["BlobContainerPath"]))
	if err != nil {
		return nil,err
	}
	containerURL := azblob.NewContainerURL(*u, azurePipeline)

	return &AzureBlob{
		sourceFilePath:          config["TargetFilePath"],
		storageAccountName:      config["StorageAccountName"],
		storageAccountAccessKey: config["StorageAccountAccessKey"],
		blobContainerPath:       config["BlobContainerPath"],
		blobContainerURL:        containerURL,
	}, nil
}
