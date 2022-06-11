# Digital Ocean Functions Challenge

This repository is a just a showcase of different languages calling and being a Digital Ocean function.

<img width="958" alt="Screen Shot 2022-06-11 at 10 37 04 AM" src="https://user-images.githubusercontent.com/4959922/173192429-13f2e0ca-113b-4fa5-add8-f4ed62dd873e.png">

## Table Of Contents

| Language | Subject             | Link                                                                            |
| -------- | ------------------- | ------------------------------------------------------------------------------- |
| Go       | CLI Using Functions | https://github.com/j4ng5y/digitalocean-functions-challenge/blob/main/go/main.go |

## Deploy with DigitalOcean CLI `doctl serverless`
```bash
doctl sls deploy .
```
Output from deploy will look like this:
```
Deploying '/path/to/main/digitalocean-functions-challenge'
  to namespace 'fn-...'
  on host 'https://example.doserverless.co'
Submitted action 'go' for remote building and deployment in runtime go:default (id: some-id)
Processing of 'go' is still running remotely ...
Deployment status recorded

Deployed functions ('doctl sls fn get <funcName> --url' for URL):
  - challenge/go
```

And to run the Function we can use `doctl serverlesss function invoke` or the shorter form shown below.
```
doctl sls fn invoke challenge/go -p sammyname:yay -p sammytype:robot
```
