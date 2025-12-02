# Auth | Test Task


## OTP

1. /otp/send    -> sends message with verification code to email +
2. /otp/confirm -> confirms code and returns jwt token that contains id, exp +

## Auth

1. /auth/signup -> sign up for users +
2. /auth/login  -> login for users and sysusers +

## SysUsers
1. /sysusers/create -> create sysusers +

## Roles

1. /roles        -> returns list of roles +
2. /roles/create -> creates role +
3. /roles/{id}   -> update role +