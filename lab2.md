### GO VAULT CLIENT LIBRARY FOR READING SECRETS
```
package main

import (
	"fmt"
	"log"

	vault "github.com/mch1307/vaultlib"
)

func main() {
    // Create a new config. Reads env variables, fallback to default value if needed
    vcConf := vault.NewConfig()

    // Create new client
	vaultCli, err := vault.NewClient(vcConf)
	if err != nil {
		log.Fatal(err)
	}

    // Get the Vault secret kv_v2/path/my-secret
	secret, err := vaultCli.GetSecret("secret/path/my-secret")
	if err != nil {
		fmt.Println(err)
	}

    // We know our secret is a key/value so we use secret.KV
    // if our secret was of JSON type, we would have used secret.JSONSecret
	for k, v := range secret.KV {
		fmt.Printf("Secret %v: %v\n", k, v)
	}
}
```
