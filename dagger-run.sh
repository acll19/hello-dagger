cp -r dagger-node/dagger .
cp dagger-node/dagger.json .

dagger call run --source=$1 --command="$2";

rm -rf ./dagger ./dagger.json