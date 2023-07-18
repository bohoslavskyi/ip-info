# IP Info App

## Getting Started
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
        "currentTime": "18.07.2023 14:54",
        "currencies": [
            {
                "currency": "AFN",
                "rateToUAH": 2.341429
            },
            {
                "currency": "FKP",
                "rateToUAH": 0.020729
            },
            {
                "currency": "LSL",
                "rateToUAH": 0.488623
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
        "currentTime": "18.07.2023 14:54",
        "currencies": [
            {
                "currency": "DZD",
                "rateToUAH": 3.657598
            },
            {
                "currency": "GHS",
                "rateToUAH": 0.307516
            },
            {
                "currency": "SEK",
                "rateToUAH": 0.277022
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
        "currentTime": "18.07.2023 07:54",
        "currencies": [
            {
                "currency": "BBD",
                "rateToCAD": 1.51553
            },
            {
                "currency": "DKK",
                "rateToCAD": 5.033273
            },
            {
                "currency": "KES",
                "rateToCAD": 107.509479
            }
            ...
        ]
    }
]
```
