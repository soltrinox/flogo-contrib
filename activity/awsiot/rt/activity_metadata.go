package awsiot

var jsonMetadata = `{
  "name": "tibco-awsiot",
  "version": "0.0.1",
  "title": "AWS IoT Activity",
  "description": "Simple AWS IoT Activity",
  "inputs":[
    {
      "name": "thingName",
      "type": "string",
      "required": true
    },
    {
      "name": "awsEndpoint",
      "type": "string",
      "required": true
    },
    {
      "name": "desired",
      "type": "params"
    },
    {
      "name": "reported",
      "type": "params"
    }
  ]
}
`