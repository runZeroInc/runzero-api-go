package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	runzeroapi "github.com/RumbleDiscovery/runzero-api-go"
	rc "github.com/RumbleDiscovery/runzero-api-go/client"
)

const defaultTimeout = time.Second * 10

func main() {
	log.SetFlags(0)
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	client := rc.NewClient(&rc.Config{})
	log.Printf("runZero Components")
	log.Printf("==================")
	aver, _, err := client.PublicApi.GetLatestAgentVersion(ctx)
	if err != nil {
		log.Fatalf("failed to get agent version: %s", err)
	}
	log.Printf("     Agent: %s", aver.Version)

	sver, _, err := client.PublicApi.GetLatestScannerVersion(ctx)
	if err != nil {
		log.Fatalf("failed to get scanner version: %s", err)
	}
	log.Printf("   Scanner: %s", sver.Version)

	pver, _, err := client.PublicApi.GetLatestPlatformVersion(ctx)
	if err != nil {
		log.Fatalf("failed to get platform version: %s", err)
	}
	log.Printf("  Platform: %s", pver.Version)

	apiKey := os.Getenv(rc.APIKey)
	if apiKey == "" {
		log.Printf("Set the %s environment variable to test authenticated APIs", rc.APIKey)
		return
	}

	if strings.HasPrefix(apiKey, "O") {
		org, _, err := client.OrganizationApi.GetOrganization(ctx)
		if err != nil {
			log.Fatalf("failed to get organization %s", err)
		}

		log.Printf("")
		log.Printf("Organization %s has %d assets", org.Name, org.AssetCount)
	} else if strings.HasPrefix(apiKey, "C") {
		ex, _, err := client.AccountApi.GetAccountAgents(ctx, &runzeroapi.AccountApiGetAccountAgentsOpts{})
		if err != nil {
			log.Fatalf("failed to get explorers (%s)", err)
		}

		log.Printf("")
		log.Printf("Account has %d registered explorers", len(ex))
	}
}
