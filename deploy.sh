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
cd ../BotSecrets
git pull
go install
cd ../BotMultiplexer
git pull
go install
cd ../ScriptureBot

echo Deploying app
gcloud app deploy --quiet