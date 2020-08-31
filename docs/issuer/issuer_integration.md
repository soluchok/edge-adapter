# Issuer Adapter

## 1 Issuer Adapter APIs
This section contains the details of Issuer Adapter APIs to be consumed by the Issuer.

### 1.1 Create Profile API - HTTP POST /profile

Creates a new Issuer profile with Adapter.

#### Request 
id : profile id <br/>
name : profile name <br/>
supportedVCContexts : VC contexts supported by the issuer <br/>
url : issuer service callback urls <br/>
supportsAssuranceCredential: falg for issuer assurance support - [refer](#24-user-data-api---http-post-urlfromissuerprofileassurance)

```
{
   "id":"2cb3d5e0-de69-4418-b673-4e37ad1d0775",
   "name":"Issuer Profile 1",
   "supportedVCContexts":[
      "https://w3id.org/citizenship/v3"
   ],
   "url":"http://issuer.example.com",
   "supportsAssuranceCredential":false
}
```

#### Response
```
{
   "id":"2cb3d5e0-de69-4418-b673-4e37ad1d0775",
   "name":"Issuer Profile 1",
   "url":"http://issuer.example.com",
   "supportedVCContexts":[
      "https://w3id.org/citizenship/v3"
   ],
   "supportsAssuranceCredential":false,
   "credentialSigningKey":"did:example:def567#key1",
   "presentationSigningKey":"did:example:def567#key1",
   "createdAt":"2020-07-27T14:31:04.250212Z"
}
```

### 1.2 Retrieve Profile API - HTTP GET /profile/{profileID}
Retrieves the issuer profile

#### Response
```
{
   "id":"2cb3d5e0-de69-4418-b673-4e37ad1d0775",
   "name":"Issuer Profile 1",
   "url":"http://issuer.example.com",
   "supportedVCContexts":[
      "https://w3id.org/citizenship/v3"
   ],
   "supportsAssuranceCredential":false,
   "credentialSigningKey":"did:example:def567#key1",
   "presentationSigningKey":"did:example:def567#key1",
   "createdAt":"2020-07-27T14:31:04.250212Z"
}
```


### 1.3 Wallet Connect API - HTTP Redirect /{profileID}/connect/wallet?state=<state>
Redirect API to connect Issuer to Wallet through Issuer Adapter. Once connected, the issuer adapter will redirect to Issuer callback url.

state : unique ID generated by the issuer for tracking the session. The unmodified value will be sent back to Issuer during callback.

## 2 Issuer APIs
This section contains the details of the APIs that need to be exposed by the Issuer to be consumed by Issuer Adapter.

### 2.1 Get Issuer Token API - HTTP POST <urlFromIssuerProfile>/token

#### Request 
```
{
   "state":"2a445655-2f36-4e26-b7fc-6811f2d6051d"
}
```
state: unmodified value from the issuer received by the adapter during wallet connect. <br/>

#### Response 
```
{
   "token":"4e725916-1992-4765-aadd-85a96e76aeaa"
}
```
token: This will be used by adapter to get the user data from the issuer <br/>

### 2.2 Wallet Connect Callback API - HTTP REDIRECT <urlFromIssuerProfile>/cb?state={state}
state: unmodified value from the issuer received by the adapter during wallet connect. <br/>

[TODO - Send wallet connect error to issuer](https://github.com/trustbloc/edge-adapter/issues/213)

### 2.3 User Data API - HTTP POST <urlFromIssuerProfile>/data
When RP asks for the user credential, the Issuer adapter calls this API to get the user data from the issuer and creates a VC.

#### Request 
```
{
   "token":"2a445655-2f36-4e26-b7fc-6811f2d6051d"
}
```

#### Response
```
{
   "data":{
      "user data"
   },
   "metadata":{
      "contexts":[
         "<json-ld contexts>"
      ],
      "scopes":[
         "<json-ld scopes>"
      ],
      "name":"<name>",
      "description":"<description>"
   }
}
```

##### Sample Response
```
{
   "data":{
      "id":"did:example:b34ca6cd37bbf23",
      "givenName":"JOHN",
      "familyName":"SMITH",
      "gender":"Male",
      "image":"data:image/png;base64,iVBORw0KGgo...kJggg==",
      "residentSince":"2015-01-01",
      "lprCategory":"C09",
      "lprNumber":"999-999-999",
      "commuterClassification":"C1",
      "birthCountry":"Bahamas",
      "birthDate":"1958-07-17"
   },
   "metadata":{
      "contexts":[
         "https://w3id.org/citizenship/v1"
      ],
      "scopes":[
         "PermanentResidentCard"
      ],
      "name":"Permanent Residence Card",
      "description":"Permanent Residence Card for John Smith (Issued by Government of Wonderland)"
   }
}
```

### 2.4 User Data API - HTTP POST <urlFromIssuerProfile>/assurance
When RP asks for the user assurance data, the Issuer adapter calls this API to get the assurance data from the issuer and creates a VC.
Note: This API is applicable when Issuer supports assurance credential - [refer](#11-create-profile-api---http-post-profile)

#### Request 
```
{
   "token":"2a445655-2f36-4e26-b7fc-6811f2d6051d"
}
```

#### Response
```
{
   "data":{
      "user data"
   },
   "metadata":{
      "contexts":[
         "<json-ld contexts>"
      ],
      "scopes":[
         "<json-ld scopes>"
      ],
      "name":"<name>",
      "description":"<description>"
   }
}
```

##### Sample Response
```
{
   "data":{
      "document_number":"123-456-789",
      "evidence_id":"d4d18a776cc6",
      "comments":"DL verified digitally"
   },
   "metadata":{
      "contexts":[
         "https://trustbloc.github.io/context/vc/examples/driver-license-evidence-v1.jsonld"
      ],
      "scopes":[
         "DrivingLicenseEvidence"
      ],
      "name":"Drivers License Evidence",
      "description":"Drivers License Evidence for John Smith"
   }
}
```