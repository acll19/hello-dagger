cp -r dagger-node/dagger .
cp dagger-node/dagger.json .

dagger call publish --source=$1 --testCmd="$2" --buildCmd="$3"

rm -rf ./dagger ./dagger.json