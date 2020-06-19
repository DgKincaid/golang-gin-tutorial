// NOT A REAL PWD used for development purposes
db.createUser(
  {
    user: 'admin',
    pwd: 'pwd123',
    roles: [
      {
        role: 'readWrite',
        db: 'gomongo'
      }
    ]
  }
)