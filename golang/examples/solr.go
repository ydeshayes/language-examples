package examples

import (
  "fmt"
  psh "github.com/platformsh/config-reader-go/v2"
  gosolr "github.com/platformsh/config-reader-go/v2/gosolr"
  solr "github.com/rtt/Go-Solr"
)

func UsageExampleSolr() string {

  // Create a NewRuntimeConfig object to ease reading the Platform.sh environment variables.
  // You can alternatively use os.Getenv() yourself.
  config, err := psh.NewRuntimeConfig()
  if err != nil {
    panic(err)
  }

  // Get the credentials to connect to the Solr service.
  credentials, err := config.Credentials("solr")
  if err != nil {
    panic(err)
  }

  // Retrieve Solr formatted credentials.
  formatted, err := gosolr.FormattedCredentials(credentials)
  if err != nil {
    panic(err)
  }

  // Connect to Solr using the formatted credentials.
  connection := &solr.Connection{URL:formatted}

  // Add a document and commit the operation.
  docAdd := map[string]interface{}{
    "add": []interface{}{
      map[string]interface{}{"id": 123, "name": "Valentina Tereshkova"},
    },
  }

  respAdd, err := connection.Update(docAdd, true)
  if err != nil {
    panic(err)
  }

  // Select the document.
  q := &solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"id:123"},
		},
	}

  resSelect, err := connection.CustomSelect(q, "query")
  if err != nil {
    panic(err)
  }

  // Delete the document and commit the operation.
  docDelete := map[string]interface{}{
    "delete": map[string]interface{}{
      "id": 123,
    },
  }

  resDel, err := connection.Update(docDelete, true)
  if err != nil {
    panic(err)
  }

  message := fmt.Sprintf(`Adding one document: %s
Selecting document (1 expected): %d
Deleting document: %s
  `, respAdd, resSelect.Results.NumFound, resDel)

  return message
}
