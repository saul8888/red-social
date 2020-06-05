
**POST**
/api/user/create
```json
{
	"username":"alonso",
	"firstName": "alonso",
	"lastName": "quispe",
	"email":"alonso@example.com",
	"password":"alonso",
	"location": "Peru"
}
```

/api/user/login
```json
{
	"email":"saul@example.com",
	"password":"saul"
}
```

**GET**
/api/user?Id=5ed988657518c4dd38b567fc


/api/user/total
```json
{
	"limit": 1,
	"offset": 0
}
```

**PUT**
/api/user?Id=5ecbf37c29bb3ef231594f75
```json
{
	"firstName": "saul",
	"lastName": "quispe",
	"email":"saul@example.com",
	"phoneNumber": "9445-98989",
	"addresses": "chorrillos"
}
```

**DELETE**
/api/employee?Id=5ecbf37c29bb3ef231594f75


