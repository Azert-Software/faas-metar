# faas-metar

An OpenFaas function for getting METAR information for an aviation weather station

# Local Setup

Assuming you have [OpenFaas](https://github.com/openfaas/faas) setup on your local machine, clone this repo and execute the make file as shown:

```
 make deploy-test-local
 
```
This will build, deploy and also execute the function using the fass-cli to ensure it is working.

You can then execute it through the OpenFaas UI, Cli or Web Calls.

# Using the Function

It takes a single string of Phonetic Characters which represent the ICAO code of the weather station in question.

My nearest station is Belfast Harbour Airport which has an ICAO of EGAC. In order to get the weather data for this station submit the value

```
  "ECHO GOLF ALPHA CHARLIE"
  
```

# Coming Soon

Alexa Support :)
