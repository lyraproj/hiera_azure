## Azure Key Vault lookup key

This function allows you to look up single values stored as secrets from an Azure Key Vault.

## Installation
Build the plugin from the root directory of this module:
```
go build -buildmode=plugin -o azure_key_vault.so
```
Then make the plugin available to Hiera. See
[Extending Hiera](https://github.com/lyraproj/hiera#Extending-Hiera) for info on how to do that.

#### A Note about debugging
When debugging from an IDE like JetBrains goland, the hiera binary will be built using `-gcflags 'all=N -l'` which
results in a slightly different binary signature. As a consequence, the loading of plugins will fail with a message
similar to "plugin was built with a different version of package X" unless the plugin is rebuilt using the same flags,
e.g.
```
go build -buildmode=plugin -o azure_key_vault.so -gcflags 'all=-N -l'
```

## Examples
A single option `vault_name` should be set in `hiera.yaml`:

    ---
    version: 5
    defaults:
      datadir: hiera
      data_hash: yaml_data

    hierarchy:
    - name: common
      path: common.yaml
    - name: secrets
      lookup_key: azure_key_vault
      options:
        vault_name: my-key-vault

There are two options for authentication, using a service principal or the Azure CLI

* To use a service principal set the environment variables `AZURE_TENANT_ID`, `AZURE_CLIENT_ID` and `AZURE_CLIENT_SECRET`

* If the above variables are not present the Azure CLI will be used (it must already be logged in)
