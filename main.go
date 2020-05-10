package main

import (
	"flag"
	"azure-file-fetch/pkg"
)

func main() {
	var err error
	var httpFile *pkg.File
	var azureBlob *pkg.AzureBlob
	var configMap = make(map[string]*string)

	configMap["Url"] = flag.String("url", "", "Url to download a file from.")
	configMap["TargetFilePath"] = flag.String("target-file-path", "", "Target file path to save the file.")
	configMap["StorageAccountName"] = flag.String("storage-account-name", "", "Target Azure storage account name.")
	configMap["StorageAccountAccessKey"] = flag.String("storage-account-access-key", "", "Target Azure storage account access key.")
	configMap["BlobContainerPath"] = flag.String("blob-container-path","","Target Azure blob container path.")

	flag.Parse()

	httpFile, err = pkg.NewFile(configMap)
	if err != nil {
		panic(err)
	}

	err = httpFile.HttpGet()
	if err != nil {
		panic(err)
	}

	azureBlob, err = pkg.NewAzureBlob(configMap)
	if err != nil {
		panic(err)
	}

	err = azureBlob.HttpPut()
	if err != nil {
		panic(err)
	}
}
