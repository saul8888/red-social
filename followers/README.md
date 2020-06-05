
**POST**
/api/employee
```json
{
	"username":"saul8",
	"firstName": "test11",
	"lastName": "quispe",
	"email":"sa@example.com",
	"password":"1234",
	"avatar":"medium",
	"banner":"large",
	"biografia":"good",
	"rolesId": "chef",
	"location": "Peru",
	"website": "saul"
}
```

**GET**
/api/user?Id=5ed7e29e9adc36e4785bcfb0

**GET**
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


