#! /bin/sh

go build
# ./ex11 create -repo git-api-test -owner negibokken
# ./ex11 get -repo git-api-test -owner negibokken -issue 123
# ./ex11 edit -repo git-api-test -owner negibokken -issue 123
# ./ex11 close -repo git-api-test -owner negibokken -issue 123

./ex11 -repo git-api-test -owner negibokken
./ex11 -repo git-api-test -owner negibokken -issue 123
./ex11 -repo git-api-test -owner negibokken -issue 123
./ex11 -repo git-api-test -owner negibokken -issue 123

