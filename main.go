package main

import (
	"flag"
	"fmt"
	"github.com/andygeiss/exceptionz/api"
	"github.com/andygeiss/exceptionz/impl"
	"golang.org/x/net/context"
	"log"
)

func main() {
	path := flag.String("path", "", "Target path")
	stats := flag.String("stats", "", "Type of statistics (totalClasses|totalExceptions|classExceptions)")
	flag.Parse()
	if *path == "" || *stats == "" {
		fmt.Printf("Usage of [exceptionz]:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExamples:\n  exceptionz -path ./testdata -stats totalClasses\n\n")
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
