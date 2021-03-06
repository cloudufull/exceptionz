package impl

import (
	"bufio"
	"fmt"
	"github.com/andygeiss/exceptionz/api"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// DefaultParsing uses a HashingServer to parse and store unique exceptions.
type DefaultParsing struct {
	hashing api.HashingServer
	path    string
}

// NewDefaultParsing creates a new ParsingServer and returns its address.
func NewDefaultParsing(path string, hashing api.HashingServer) api.ParsingServer {
	return &DefaultParsing{
		hashing: hashing,
		path:    path,
	}
}

// ParseExceptions creates statistics about exceptions for directories and regular files.
func (p *DefaultParsing) ParseExceptions(ctx context.Context, req *api.ParseExceptionsRequest) (*api.ParseExceptionsResponse, error) {
	classMap := make(map[uint64]string)
	counterMap := make(map[uint64]uint64)
	exceptionMap := make(map[uint64]string)
	// 1) Get path stat.
	info, err := os.Stat(p.path)
	if err != nil {
		return nil, err
	}
	// 2) Handle directories and regular files.
	switch mode := info.Mode(); {
	case mode.IsDir():
		files, err := ioutil.ReadDir(p.path)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			addExceptionsFrom(p.path+"/"+file.Name(), p.hashing, classMap, counterMap, exceptionMap)
		}
	case mode.IsRegular():
		addExceptionsFrom(p.path, p.hashing, classMap, counterMap, exceptionMap)
	}
	// 3) Print type-specific statistics.
	switch req.Type {
	case api.StatisticsType_TOTAL_CLASSES:
		classes := make(map[string]int, 0)
		for _, classname := range classMap {
			classes[classname]++
		}
		fmt.Printf("Unique Exceptions;Classname\n")
		for classname, _ := range classes {
			fmt.Printf("%d;%s\n", classes[classname], classname)
		}
	case api.StatisticsType_TOTAL_EXCEPTIONS:
		exceptions := make(map[string]int, 0)
		for _, message := range exceptionMap {
			exceptions[message]++
		}
		fmt.Printf("Exception Occured;Exception\n")
		for name, _ := range exceptions {
			fmt.Printf("%d;%s\n", exceptions[name], name)
		}
	case api.StatisticsType_CLASS_EXCEPTIONS:
		fmt.Printf("Exception Occured;Exception;Classname\n")
		for key, _ := range exceptionMap {
			classname := classMap[key]
			exception := exceptionMap[key]
			counter := counterMap[key]
			fmt.Printf("%d;%s;%s\n", counter, exception, classname)
		}
	}
	return &api.ParseExceptionsResponse{}, nil
}

func addExceptionsFrom(path string, hashing api.HashingServer, classMap map[uint64]string, counterMap map[uint64]uint64, exceptionMap map[uint64]string) {
	// 1) Prepare RegExpr
	exceptionExpr := regexp.MustCompile(`[\.|\w]+Exception\:`)
	timestampExpr := regexp.MustCompile(`^\d+\-\d+\-\d+\s+\d+\:\d+\:\d+`)
	// 2) Open regular file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 3) Scan files for exceptions
	scanner := bufio.NewScanner(file)
	classname := ""
	for scanner.Scan() {
		// DATE TIME SERVLET UUID DATE TIME LOGLEVEL CLASSNAME CLIENTID USERID IP UUID2 MESSAGE ...
		line := scanner.Text()
		// if a timestamp is found then update the classname
		if timestampExpr.MatchString(line) {
			parts := strings.Split(line, " ")
			// skip non-error messages
			// skip messages with less than 12 parts
			if len(parts) < 12 && parts[6] != "ERROR" {
				continue
			}
			classname = parts[7]
		} else if exceptionExpr.MatchString(line) {
			classname = strings.Replace(classname, "[", "", -1)
			classname = strings.Replace(classname, "]", "", -1)
			id := exceptionExpr.FindStringSubmatch(line)[0]
			id = strings.Replace(id, ":", "", -1)
			// create a hash value by using classname and id.
			res, _ := hashing.Hash(context.Background(), &api.HashRequest{Plain: []byte(classname + id)})
			classMap[res.Hash] = classname
			counterMap[res.Hash]++
			exceptionMap[res.Hash] = id
		}
	}
}
