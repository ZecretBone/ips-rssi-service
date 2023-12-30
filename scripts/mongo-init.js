rs.status();
use admin; // eslint-disable-line
db.createUser({user: 'admin', pwd: 'admin', roles: [ 
    { role: 'userAdminAnyDatabase', db: 'admin' },
    { role: "dbAdminAnyDatabase", db: "admin" }, 
    { role: "readWriteAnyDatabase", db: "admin" },
    { role: 'readWrite', db: 'ips-rssi-beta' }
]});
use ips-rssi-beta; // eslint-disable-line
db.createCollection("signal-stat-collection");
