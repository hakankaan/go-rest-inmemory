# go-rest-inmemory
REST-API for in-memory database store

### Details
- Endpoint for setting a key value pair
- Endpoint for getting value of a key
- Endpoint for flushing whole store datas
- Writing to disk withing a certain time range
- When the application stops and stands up again, if there is a saved file, reload existing data into memory

### Running
In the main directory:
`chmod +x script/run-dev.sh`
`chmod +x script/run-prod.sh`

For development:
`script/run-dev.sh`

For production:
`script/run.sh`