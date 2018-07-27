# hashing - Exceptions

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/hashing)](https://goreportcard.com/report/github.com/andygeiss/hashing)
[![Build Status](https://travis-ci.org/andygeiss/hashing.svg?branch=master)](https://travis-ci.org/andygeiss/hashing)

A practical implementation of using hash/maps to create statistics about "Exceptions" in [Apache Tomcat](http://tomcat.apache.org/) logfiles.

## Installation

Use your Go toolkit to get the sources and compile <code>hashing</code> from scratch:

    go get -u github.com/andygeiss/hashing

## Usage

The current version supports three different statistics (see [API](https://github.com/andygeiss/hashing/blob/master/api/parsing.proto)):
1. Get the total amount of unique Exceptions occured:               <b>totalExceptions</b>
2. Get the amount of unique Exceptions occured per Class:           <b>totalClasses</b>
3. Get the amount of Exceptions mapped to each Class (detailed):    <b>classExceptions</b>

        Usage of [hashing]:
        
          -path string
                Target path
          -stats string
                Type of statistics (totalClasses|totalExceptions|classExceptions)
        
        Examples:
          hashing -path ./testdata -stats totalClasses

## Output

The output will be written as CSV with semicolon as separator:

    Exception Occured;Exception;Classname
    1;java.io.IOException;Classname
