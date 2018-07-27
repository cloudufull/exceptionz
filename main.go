package main

import (
	"flag"
	"github.com/andygeiss/hashing/api"
	"github.com/andygeiss/hashing/impl"
	"golang.org/x/net/context"
	"log"
)

func main() {

	path := flag.String("path", "testdata", "Target path")
	stats := flag.String("stats", "totalClasses", "Type of statistics (totalClasses|totalExceptions|classExceptions)")
	flag.Parse()

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
