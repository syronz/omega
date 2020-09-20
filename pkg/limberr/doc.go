//Package limberr means limb + error, errors which can be append to the limb of main error
/*
	Package limberr implements error handling based on chain and parent/child
	relation, id consist two part. generating he error and parse it.

	Generating:
		Any error in any level (repo, service, api and model) could be entered to
		the limberr process. limberr don't delete or affect the original error but
		each time it can extend the error by wrapping or add properties.
		For wrapping default Go's Errorf with %w implemented

	Examples:
		err = limberr.AddCode(err, "E1098312")
		err = limberr.AddMessage(err, "database error")
		err = limberr.AddPath(err, "/roles/")
		err = limberr.AddType(err, "http://54323452", "duplicate error")
		err = limberr.AddDomain(err, "base")
		err = fmt.Errorf("some extra information, %w", err)

	Parsing:
		For parsing you can use it in response to add path in easiear way


	Final format:
	It would be in JSON and it looks like below
		{
			"type": "http//link.com/to/order",
			"title": "duplication",
			"message": "user with this name already exist",
			"code": "E321343",
			"path": "users/32",
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

*/
package limberr
