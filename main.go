package main

import (
	"flag"
	"github.com/andygeiss/hashing/api"
	"github.com/andygeiss/hashing/impl"
	"golang.org/x/net/context"
	"log"
	"fmt"
)

func main() {

	path := flag.String("path", "", "Target path")
	stats := flag.String("stats", "", "Type of statistics (totalClasses|totalExceptions|classExceptions)")
	flag.Parse()

	if *path == "" || *stats == "" {
		fmt.Printf("Usage of [hashing]:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n  hashing -path ./testdata -stats totalClasses\n\n")
		return
	}

	hashing := impl.NewDefaultHashing()
	service := impl.NewDefaultParsing(*path, hashing)

	var request *api.ParseExceptionsRequest
	switch *stats {
	case "totalClasses":
		request = &api.ParseExceptionsRequest{Type: api.StatisticsType_TOTAL_CLASSES}
	case "totalExceptions":
		request = &api.ParseExceptionsRequest{Type: api.StatisticsType_TOTAL_EXCEPTIONS}
	case "classExceptions":
		request = &api.ParseExceptionsRequest{Type: api.StatisticsType_CLASS_EXCEPTIONS}
	}

	_, err := service.ParseExceptions(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
}
