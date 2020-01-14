# Wire-Sandbox


[![BuildStatus](https://api.travis-ci.org/syronz/omega.svg?branch=master)](http://travis-ci.org/syronz/omega) 
[![ReportCard](https://goreportcard.com/badge/github.com/syronz/omega)](https://goreportcard.com/report/github.com/syronz/omega) 
[![codecov](https://codecov.io/gh/syronz/omega/branch/master/graph/badge.svg)](https://codecov.io/gh/syronz/omega)
[![Coverage Status](https://coveralls.io/repos/github/syronz/omega/badge.svg?branch=master)](https://coveralls.io/github/syronz/omega?branch=master)
[![codebeat badge](https://codebeat.co/badges/f7ed90cf-4793-4b82-acd3-00fecf4e3817)](https://codebeat.co/projects/github-com-syronz-omega-master)
[![GolangCI](https://golangci.com/badges/github.com/gojek/darkroom.svg)](https://golangci.com/r/github.com/syronz/omega)
[![GoDoc](https://godoc.org/github.com/syronz/omega?status.png)](https://godoc.org/github.com/syronz/omega)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


Simple app for analyzing golnag clean design.
Inspired by

https://hellokoding.com/crud-restful-apis-with-go-modules-wire-gin-gorm-and-mysql/

and 

https://github.com/qiangxue/go-restful-api

### Run
in the main directory

```
source config/envs.sample
reflex -r '\.go' -s -- sh -c 'go run cmd/omega/main.go'
```


