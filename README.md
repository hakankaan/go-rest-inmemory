# go-rest-inmemory
REST-API for in-memory database store

### Running 🚀
In the main directory:
`chmod +x script/run.sh`
`script/run-dev.sh`


### Details 📙
- Endpoint for setting a key value pair
- Endpoint for getting value of a key
- Endpoint for flushing whole store datas
- Writing to disk withing a certain time range
- When the application stops and stands up again, if there is a saved file, reload existing data into memory

### Endpoints 📍
**Live:** go-re-LoadB-1TGDOITRSHVHB-88b23d1182a1b292.elb.us-east-1.amazonaws.com

**GET /api/datas/{key}**
Returns value of given key

**POST /api/datas**
Creates key-value pair with given request body below
- key: string
- value: string

**DELETE /api/datas**
Flushes in-memory store