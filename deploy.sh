DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
echo Deploying from $DIR

# echo Installing all necessary packages
# go get -u -v -f all

echo Installing app entry package
# Update and install the entry module
cd $DIR
git pull
go install

# Update and install the sub modules
cd ../BotPlatform
git pull
go install github.com/julwrites/BotPlatform/pkg/def
go install github.com/julwrites/BotPlatform/pkg/secrets
go install github.com/julwrites/BotPlatform/pkg/platform
cd ../ScriptureBot

echo Deploying app
gcloud app deploy --quiet