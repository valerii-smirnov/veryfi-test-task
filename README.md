# Veryfi test assignment (Valerii Smirnov)

# Part One

## Overview

In the first part of the test assignment, I implemented an event source application that monitors the state of the 
directory specified in the launch configuration and processes the invoices uploaded to that directory using the 
Veryfi OCR SDK.

1. The `document` service monitors files in a directory. When new files are added to the directory, it interacts with the 
    Veryfi OCR SDK and saves the results in the database as raw data. Additionally, after successful processing, it 
    publishes an event into NATS Jetstream to notify the second service about the event. The service also implements a 
    gRPC server to allow other services to request data about processed files.
2. The `stats` service listens to events about created and processed documents from NATS Jetstream. When it receives an 
    event about new processed document, it makes a gRPC request to the Document service to obtain raw data 
    about the document. Based on this raw data, it stores information in its own database. Additionally, while 
    processing events, the service enriches data with coordinates using the Google Geocoding API based on the address 
    obtained from the document. Furthermore, the `stats` service implements a REST server with three endpoints to obtain 
    statistics over a specified time interval based on the data that has been stored in the database thanks to 
    Veryfi OCR SDK, Google Geocoding API, and the Document service.

Both applications, under the hood, use the Worker Pool pattern and process events concurrently in multiple threads, 
which significantly improves the application's performance when a large number of invoices are uploaded to the directory 
simultaneously.

## Stack

1. Golang
2. PostgreSQL
2. NATS Jetstream
3. gRPC
4. GORM
5. Uber FX 
6. gomock
7. sqlmock

## How to run

The application is fully containerized and configured to be launched with a single command.
### Requirements:
1. docker
2. golang 1.21

### Config and run:
1. Add necessary values into [.docker.env](part_one/services/document/.docker.env) for `document` service. Replace `<your_value>` hints with your secrets.
2. Add necessary values into [.docker.env](part_one/services/stats/.docker.env) for `stats` service. Replace `<your_value>` hints with your secrets.
3. Run the command
```shell
cd part_one && \
docker-compose up -d --build
```

## How to use

[docker-compose.yaml](part_one/docker-compose.yaml) file is configured to listen to changes in the [data](part_one/data)
directory. Of course, you can reconfigure it by changing the mount folder option for document service, but to keep the 
simplicity, I recommend using this [data](part_one/data) folder to upload receipts.

When all services will be up and running, you'll be able to upload the receipts to the [data](part_one/data).
After the files are uploaded, you can interact with the REST server to obtain statistics over specified time intervals.

REST server is available on http://localhost:8080

Available endpoints:

1. Get total paid taxes for time period
```shell
curl --location 'localhost:8080/stats/total/tax?from=2023-10-10%2000%3A00%3A00&to=2023-10-14%2000%3A00%3A00'
```
2. Get total discounts for time period
```shell
curl --location 'localhost:8080/stats/total/discount?from=2023-10-10%2000%3A00%3A00&to=2023-10-14%2000%3A00%3A00'
```
3. Get purchases geography for time period
```shell
curl --location 'localhost:8080/stats/geography?from=2023-10-10%2000%3A00%3A00&to=2023-10-14%2000%3A00%3A00'
```

## Additional Information

Since this is a test assignment, I didn't focus extensively on full functionality and covering all possible cases. 
My main goal was to demonstrate how I design applications and the development approaches I use. Additionally, 
due to time constraints, I wasn't able to provide full unit test coverage for the entire codebase. Therefore, I covered 
a few layers of one of the services with unit tests, solely for demonstration purposes. Please refer to packages 
[usecases](part_one/services/stats/internal/usecases) and 
[repositories](part_one/services/stats/internal/repositories) for this purpose.

Run tests:
```shell
cd part_one/services/stats && \
go test --race ./...
```

Time spent: ~16 hours.
# Part Two

You can find the implementation of the anagram grouping function and tests [here](part_two)

The function result is not sorted alphabetically for groups and values within groups.
This is because the algorithm implementation is based on a hash table,
which does not guarantee the order of elements in Go.
I've omitted the result sorting to achieve the optimal algorithm complexity of O(N*M).
I could add sorting to the result of the function, but this would increase the algorithm's complexity to O(N * M log M).
In this solution, I wanted to implement the most efficient algorithm possible.

Run tests:
```shell
cd part_two && \
go test --race ./...
```