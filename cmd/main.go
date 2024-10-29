package main

import (
	"fmt"

	_ "cloud.google.com/go"
	_ "cloud.google.com/go/bigquery"
	_ "cloud.google.com/go/cloudsqlconn"
	_ "cloud.google.com/go/compute/metadata"
	_ "cloud.google.com/go/container"
	_ "cloud.google.com/go/datastore"
	_ "cloud.google.com/go/logging"
	_ "cloud.google.com/go/profiler"
	_ "cloud.google.com/go/pubsub"
	_ "cloud.google.com/go/storage"
	_ "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	_ "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	_ "github.com/BurntSushi/toml"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/IBM/sarama"
	_ "github.com/adrg/xdg"
	_ "github.com/cenkalti/backoff"
	_ "github.com/dustin/go-humanize"
	_ "github.com/emicklei/proto"
	_ "github.com/go-stack/stack"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/gomicro/postal"
	_ "github.com/google/gofuzz"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/mux"
	_ "github.com/hablullah/go-hijri"
	_ "github.com/hashicorp/go-multierror"
	_ "github.com/hmarr/codeowners"
	_ "github.com/iancoleman/strcase"
	_ "github.com/spf13/cobra"
)

func main() {
	fmt.Println("foo")
}
