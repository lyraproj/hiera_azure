module github.com/lyraproj/hiera_azure

go 1.12

require (
	github.com/Azure/azure-sdk-for-go v33.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.1
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/lyraproj/dgo v0.0.0-20190918142629-ee8b686ca69b
	github.com/lyraproj/hiera v0.0.0-20190820132249-b08bb003a3b6
	github.com/lyraproj/issue v0.0.0-20190606092846-e082d6813d15
	github.com/lyraproj/pcore v0.0.0-20190918201925-7e14d50f3d7d
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.3.0+incompatible
	github.com/lyraproj/hiera => github.com/thallgren/hiera v0.0.0-20190920152706-3eecfa8da3f1
)
