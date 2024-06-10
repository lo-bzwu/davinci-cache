# Davinci Cache

This simple go app provides a cache for the davinci timetable system. 

## Deployment 

`docker run -e DAVINCI_URL="https://xxxx" -p 8000:8000 ghcr.io/lo-bzwu/davinci-cache`

## URLs

- http://localhost:8000/classes
This endpoint lists all classes available on the instance as newline-seperated values

- http://localhost:8000/lessons?classes=xxxxx&teachers=xxxxxxx
This endpoint lists lessons for all of the classes and teachers provided. It includes lookups to check the name of the class, for example.
