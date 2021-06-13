#!/bin/zsh
echo build domainfinder...
go build -o domainfinder

echo build synonyms...
cd ./synonyms
go build -o ../lib/synonyms
cd ../

echo build available...
cd ./available
go build -o ../lib/available
cd ../

echo build sprinkle...
cd ./sprinkle
go build -o ../lib/sprinkle
cd ../

echo build coolify...
cd ./coolify
go build -o ../lib/coolify
cd ../

echo build domainify...
cd ./domainify
go build -o ../lib/domainify
cd ../

echo build success
