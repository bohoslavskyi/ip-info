# IP Info App

## Getting Sterted
This app allows to get information about country and currency based on IP address.

### Prerequisites
Required software:
* **Docker**: >=__23.x.x__
* **Go**: >=__1.20.x__ (optional, it's needed to run the app locally)

## How to up and run
1. Clone the repository:
```shell
git clone https://github.com/bohoslavskyi/ip-info && cd ip-info
```
2. Prepare the configuration file:
```shell
cp .env.example .env
```
3. Build the app:
```shell
make build
```
4. Run the app:
```shell
make run port=8000:8000
```

## Usage
According to get information about a country and a currency based on IP addresses, it needs to do the following:
1. Make a request:
```cURL
curl --location 'localhost:8000/ip-info' \
--header 'Content-Type: application/json' \
--data '{
    "ips": [
        "93.79.41.110",
        "93.79.0.180",
        "24.48.0.1"
    ]
}'
```
2. Obtain a result:
```json
[
    {
        "ip": "93.79.0.180",
        "country": "Ukraine",
        "city": "Sumy",
        "latitude": 50.9207,
        "longitude": 34.7959,
        "currentTime": "07.718.18000 11:00",
        "currencies": [
            {
                "currency": "GIP",
                "rateToUAH": 0.020729
            },
            {
                "currency": "SGD",
                "rateToUAH": 0.035799
            },
            {
                "currency": "UGX",
                "rateToUAH": 100.361446
            }
            ...
        ]
    },
    {
        "ip": "93.79.41.110",
        "country": "Ukraine",
        "city": "Sumy",
        "latitude": 50.9207,
        "longitude": 34.7959,
        "currentTime": "07.718.18000 11:00",
        "currencies": [
            {
                "currency": "IRR",
                "rateToUAH": 1151.707146
            },
            {
                "currency": "LAK",
                "rateToUAH": 519.668212
            },
            {
                "currency": "TJS",
                "rateToUAH": 0.295177
            }
            ...
        ]
    },
    {
        "ip": "24.48.0.1",
        "country": "Canada",
        "city": "Montreal",
        "latitude": 45.504,
        "longitude": -73.552,
        "currentTime": "07.718.18000 11:00",
        "currencies": [
            {
                "currency": "MVR",
                "rateToCAD": 11.672773
            },
            {
                "currency": "SGD",
                "rateToCAD": 1.001431
            },
            {
                "currency": "ILS",
                "rateToCAD": 2.752658
            }
            ...
        ]
    }
]
```
