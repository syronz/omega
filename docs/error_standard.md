# Errors

In this project we follow https://tools.ietf.org/html/rfc7807 for defining the custom errors.
We also support multi-language for returning the errors

```JSON
{
  "type": "https://example.com/probs/out-of-credit",
  "title": "You do not have enough credit.",
  "detail": "Your current balance is 30, but that costs 50.",
  "instance": "/account/12345/msgs/abc",
  "balance": 30,
  "accounts": ["/account/12345",
  "/account/67890"]
}
```

Instance for field-errors
```JSON
{
  "type": "https://example.net/validation-error",
  "title": "Your request parameters didn't validate.",
  "invalid-params": [ 
    {
      "name": "age",
      "reason": "must be a positive integer"
    },
    {
      "name": "color",
      "reason": "must be 'green', 'red' or 'blue'"
    }
  ]
}
```

error-code and domain will added to the errors and the final view be like below
```JSON
{
  "type": "https://example.com/link/to/error",
  "title": "translated message to point the error, not accept params",
  "detail": "explain the error in translated form, it accept params also",
  "instance": "/redirect/link",
  "code": "ER00343123",
  "invalid-params": [ 
    {
      "name": "age",
      "reason": "must be a positive integer"
    },
    {
      "name": "color",
      "reason": "must be 'green', 'red' or 'blue'"
    }
  ]
}

