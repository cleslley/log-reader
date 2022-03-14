db.getSiblingDB('logsdb').createUser(
    {
        user: "userdb",
        pwd: "password",
        roles: [
            {
                role: "readWrite",
                db: "logsdb"
            }
        ]
    }
);