# OMEGA

[![BuildStatus](https://api.travis-ci.org/syronz/omega.svg?branch=master)](http://travis-ci.org/syronz/omega) 
[![ReportCard](https://goreportcard.com/badge/github.com/syronz/omega)](https://goreportcard.com/report/github.com/syronz/omega) 
[![codecov](https://codecov.io/gh/syronz/omega/branch/master/graph/badge.svg)](https://codecov.io/gh/syronz/omega)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6938819425f94f6f9d8046b4fdfdcbc1)](https://www.codacy.com/manual/syronz/omega?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=syronz/omega&amp;utm_campaign=Badge_Grade)
[![Coverage Status](https://coveralls.io/repos/github/syronz/omega/badge.svg?branch=master)](https://coveralls.io/github/syronz/omega?branch=master)
[![codebeat badge](https://codebeat.co/badges/f7ed90cf-4793-4b82-acd3-00fecf4e3817)](https://codebeat.co/projects/github-com-syronz-omega-master)
[![Maintainability](https://api.codeclimate.com/v1/badges/129904e9ab5aca417faa/maintainability)](https://codeclimate.com/github/syronz/omega/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/129904e9ab5aca417faa/test_coverage)](https://codeclimate.com/github/syronz/omega/test_coverage)
[![GolangCI](https://golangci.com/badges/github.com/gojek/darkroom.svg)](https://golangci.com/r/github.com/syronz/omega)
[![GoDoc](https://godoc.org/github.com/syronz/omega?status.png)](https://godoc.org/github.com/syronz/omega)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


## Run
in the main directory

```bash
source config/envs.sample
reflex -r '\.go' -s -- sh -c 'go run cmd/omega/main.go'
```

## Logrus levels

```go
plog.ServerLog.Trace(err.Error())
plog.ServerLog.Debug(err.Error())
plog.ServerLog.Info(err.Error())
plog.ServerLog.Warn(err.Error())
plog.ServerLog.Error(err.Error())
plog.ServerLog.Fatal(err.Error())
plog.ServerLog.Panic(err.Error())
```

#TODO
- [ ] if types.Resource not used in core it should moved to the base domain, in the future I decide about that
- [x] apilogger should be moved to other place
- [ ] implement refresh token
- [ ] common method for repo based on inheritance

# Requesed RMS part
1. inventory import should lock the price for agent
2. transfer should be like bellow:

  location a => location b

  item | QTY | Price | Total

  -----|-----|-------|-------

  item1| 32  | 30000 | 960000

3. expiration date on direct-recharge invoice
4. bulk direct recharge
5. finance report: separate direct recharge
6.
7. notification or approve management for return items
8. unique serial for serial base items
9. special process for updating the phone
10. enable static ip


