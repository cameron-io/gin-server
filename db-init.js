db = db.getSiblingDB('gopher');
db.createUser({
    user: process.env.DATABASE_USER,
    pwd: process.env.DATABASE_PASS,
    roles: [{ role: 'dbOwner', db: 'gopher' }]
});