# Test coverage in Go, the whole package

This repository contains the code examples for the talk 

"Test coverage in Go, the whole package"

[GopherCon UK, 2021 schedule](https://www.gophercon.co.uk/schedule/)

## Blog posts

* [Go test coverage, working with packages and sub-packages](https://eleni.blog/2021/01/10/test-coverage-in-go-working-with-packages-and-sub-packages/)
* [Deep diving in the Go coverage profile](https://eleni.blog/2021/01/24/deep-diving-in-the-go-coverage-profile/)

## The cmd script

Navigate in the directory of the module you want to generate the coverage stats for and run the following 

```shell
$ go test --count=1 -coverprofile=coverage.out ./... ; \
cat coverage.out | \
awk 'BEGIN {cov=0; stat=0;} \
	$3!="" { cov+=($3==1?$2:0); stat+=$2; } \
    END {printf("Total coverage: %.2f%% of statements\n", (cov/stat)*100);}'

```