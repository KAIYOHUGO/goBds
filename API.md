ID:name base63url encode

```bash
/+
 |-api/+
 |     |-session #post=create session, delete=signout
 |     |
 |     |-user/+ #post=create account, delete=delete account
 |     |      |
 |     |      |
 |     |      |-{UserID}+
 |     |                |-servers
 |     |                |-config
 |     |
 |     |
 |     |-server/+ #post=create server, delete=delete server
 |              |
 |              |
 |              |{ServerID}/+
 |                          |-file
 |                          |-input
 |
 |
 |-wss/+
       |-server/{ServerID}/+
                            |-terminal/{SessionID}

```