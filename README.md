# Auth

## Roles

1. /roles -> returns list of roles
2. /roles/create -> creates role
3. /roles/{id} -> update role

## OTP

1. /otp/send -> sends message with verification code to email
2. /otp/confirm -> confirms code and returns jwt token that contains id, exp
