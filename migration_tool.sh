if [ "$1" = "create" ]; then
    echo "Creating migration"
    ./migrate create -ext sql -dir=database/migrations -seq init
elif [ "$1" = "run" ]; then 
    echo "Applying migration"
    ./migrate -path=database/migrations -database postgres://postgres:1337@localhost:5432/doc-review -verbose up
fi
