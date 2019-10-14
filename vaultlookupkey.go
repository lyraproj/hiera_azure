package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/go-autorest/autorest"
	"github.com/lyraproj/dgo/dgo"
	"github.com/lyraproj/hierasdk/hiera"
	"github.com/lyraproj/hierasdk/plugin"
	"github.com/lyraproj/hierasdk/register"
)

func main() {
	register.LookupKey(`azure_key_vault`, AzureKeyVaultLookupKey)
	plugin.ServeAndExit()
}

// AzureKeyVaultLookupKey looks up a single value from an Azure Key Vault
func AzureKeyVaultLookupKey(hc hiera.ProviderContext, key string) dgo.Value {
	if key == `lookup_options` {
		return nil
	}
	vaultName, ok := hc.StringOption(`vault_name`)
	if !ok {
		panic(fmt.Errorf(`missing required provider option 'vault_name'`))
	}
	var authorizer autorest.Authorizer
	var err error
	if os.Getenv("AZURE_TENANT_ID") != "" && os.Getenv("AZURE_CLIENT_ID") != "" && os.Getenv("AZURE_CLIENT_SECRET") != "" {
		authorizer, err = auth.NewAuthorizerFromEnvironment()
	} else {
		authorizer, err = auth.NewAuthorizerFromCLI()
	}
	if err != nil {
		panic(err)
	}
	client := keyvault.New()
	client.Authorizer = authorizer
	resp, err := client.GetSecret(context.Background(), "https://"+vaultName+".vault.azure.net", key, "")
	if err != nil {
		if ResponseWasStatusCode(resp.Response, http.StatusNotFound) {
			return nil
		}
		panic(err)
	}
	return hc.ToData(*resp.Value)
}

func ResponseWasStatusCode(resp autorest.Response, statusCode int) bool {
	if r := resp.Response; r != nil {
		if r.StatusCode == statusCode {
			return true
		}
	}
	return false
}
