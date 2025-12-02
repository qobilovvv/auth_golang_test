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


### How to run:
1. make run
2. go run main.go

## details
As you can see in .env.example we have variables ADMIN_PHONE and ADMIN_PHONE,
and this we need to create super admin(sys_user), when first we run the project,
script will check the sysusers and if there is no other sysusers, it will create one, with phone and password from .env, you can also delete it in main.go
```
   helpers.InitSuperAdmin(sysUsersRepo)
```