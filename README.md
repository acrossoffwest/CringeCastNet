# CringeCastNet

inspired by [CringeCast](https://github.com/cr1tbit/CringeCast) and with help of ChatGPT

## Usage
1) Run somewhere a MQTT server. A MQTT docker container you could find in `./docker` directory.

2) Run server:

        cringecast-server
        
3) Run clients:
    
        cringecast-client


## Server & Client have a args:
### MQTT Arguments:

--mqtt-url= - URL to MQTT server, default: tcp://localhost:1883

--mqtt-topic= - queue topic, default: cringecast

--mqtt-username= - default: admin

--mqtt-password= - default: pass

## API

Text to speech

    GET /say?query={text}&lang={lang code: en, pl}
    POST /say
    Request JSON Body:
    {
        "query": "{text}",
        "lang": "{lang code: en, pl}"
    }

Play audio file from URL:

    GET /play?url={url}
    POST /play
    Request JSON Body:
    {
        "url": "{url}"
    }