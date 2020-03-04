module github.com/lyraproj/hiera_azure

go 1.12

require (
	github.com/Azure/azure-sdk-for-go v33.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.1
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/lyraproj/dgo v0.4.1
	github.com/lyraproj/hierasdk v0.4.1
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.3.0+incompatible
