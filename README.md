## Assignment-1: Book Information Service

I have in this assignment developed a REST web application in Go that provides the client with information about books available in a given language based on the Gutenberg library. The service further determines the number of potential readers presumed to be able to read books in that language.

The REST web services being used for this purposes are:

Gutendex API

- Endpoint: http://129.241.150.113:8000/books/
- Documentation: http://129.241.150.113:8000/

Language2Countries API

- Endpoint: http://129.241.150.113:3000/language2countries/
- Documentation: http://129.241.150.113:3000/

REST Countries API

- Endpoint: http://129.241.150.113:8080/v3.1
- Documentation: http://129.241.150.113:8080/

The first API focuses on the provision of Gutenberg library information, whereas the second one provides the mapping from language to country information. The last service provides detailed country information.

## Deployment

Deployment

- Render:

* The web service has been deployed on render: https://cloud-assignment-1-j2kj.onrender.com/

- Localhost

* Clone the repository by use of either SSH or HTTPS.
* Open the program in your preferred IDE and run the program using “go run main.go”
* The default server runs on https://localhost8080

## Endpoints

The web service has three resource root paths:

```
/librarystats/v1/bookcount/
/librarystats/v1/readership/
/librarystats/v1/status/
```

If the web service is run on render the url will look like this:

```
https://cloud-assignment-1-j2kj.onrender.com/librarystats/v1/bookcount/
https://cloud-assignment-1-j2kj.onrender.com/librarystats/v1/readership/
https://cloud-assignment-1-j2kj.onrender.com/librarystats/v1/status/
```

If the web service is run on localhost the url will look like this:

```
http://localhost:8080/librarystats/v1/bookcount/
http://localhost:8080/librarystats/v1/readership/
http://localhost:8080/librarystats/v1/status/
```

## Return book count for a given language(s)

The initial endpoint focuses returns the count of books for any given language, identified via country 2-letter language ISO codes (ISO 639 Set 1), as well as the number of unique authors. This can be a single as well as multiple languages (comma-separated language codes).

- Request

```
Method: GET
Path: bookcount/?language={:two_letter_language_code+}/
```

two_letter_language_code is the corresponding 2-letter language ISO codes (ISO 639 Set 1)

Example Request:
- bookcount/?language=no,fi

* Example Response:

```
[
    {
        "language": "no",
        "books": 21,
        "authors": 16,
        "fraction": 0.00028779739063699157
    },
    {
        "language": "fi",
        "books": 2834,
        "authors": 887,
        "fraction": 0.03883894309834448
    }
]
```

## Return the number of potential readers for a given language

The second endpoint returns the number of potential readers for books in a given language, i.e., the population per country in which that language is official (and hence assuming that the inhabitants can potentially read it). This is reported in addition to the number of books and authors associated with a given language.

- Request

```
Method: GET
Path: readership/{:two_letter_language_code}{?limit={:number}}
```

{:two_letter_language_code} refers to the ISO639 Set 1 identifier of the language for which you establish readership.
{?limit={:number}} is an optional parameter that limits the number of country entries that are reported (in addition to the total number).

Example Request:
- readership/no/?limit=5

* Example Response:

```
[
    {
        "country": "Iceland",
        "isocode": "IS",
        "books": 7,
        "authors": 8,
        "readership": 366425
    },
    {
        "country": "Norway",
        "isocode": "NO",
        "books": 21,
        "authors": 16,
        "readership": 5379475
    },
    {
        "country": "Svalbard and Jan Mayen Islands",
        "isocode": "SJ",
        "books": 0,
        "authors": 0,
        "readership": 2562
    }
]
```

# Getting a status overview of services

The diagnostics interface indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service. 

- Request

```
Method: GET
Path: status/
```

Example Request:
* status/

* Example Response:

{
    "gutendexapi": 200,
    "languageapi": 204,
    "countriesapi": 200,
    "version": "v1",
    "uptime": 1850.885586077
}