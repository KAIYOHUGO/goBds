```bash
/+
 |-api/+
 |     |-session #post=create session, delete=signout
 |     |
 |     |-user/+ #post=create account, delete=delete account
 |     |      |
 |     |      |
 |     |      |-{UserID}+
 |     |                |-server
 |     |                |-config
 |     |
 |     |
 |     |-server/{ServerID}+
 |                         |-file
 |
 |
 |-wss/+
       |-server/{ServerID}/+
                            |-terminal

```