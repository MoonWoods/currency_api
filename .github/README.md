# 💵 💷 Currency API 💴 💶
Golang API to check currency details and exchange rates.

At this time, 200 currencies are available and include cryptocoins.

## Endpoints

- `localhost:8080/health`

    Just returns a 200 indicating that the api is live.

- `localhost:8080/all`

    Returns a JSON dictionary of available currencies in the format: 

        <currency-code>:<currency-common-name>
    
    Ex:

    ```json
    {
        ...
        "btc": "Bitcoin",
        "usd": "US Dollar",
        "mxn": "Mexican Peso",
        ...
    }
    ```

- `localhost:8080/exchange/<currency-code>`

    Returns an object that includes the date and a list of exchange with all currencies based off the requested one:

    ```json
    {
        "date": "2025-02-18",
        "eur": {
            "1inch": 3.9004839,
            "aave": 0.0040031729,
            "ada": 1.35441158,
            "aed": 3.84382391,
            ...
        }
    }
    ```

## To run API

Ensure you have built your binary:

```sh
go build -o currency_api .
```

and then run it with:

```sh
./currency_api
```

Or just run it with:
```sh
go run .
```
