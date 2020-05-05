# JWTCracker

JWTCracker is a very simple command line application used to look for JWT secret using brute force method.

This library is created purely for learning pouropses- if you are looking for real tool for cracking JSON Web Token
you should probably take a look on [Hashcat](https://github.com/hashcat/hashcat), [JohnTheRipper](https://github.com/magnumripper/JohnTheRipper),
or [c-jwt-cracker](https://github.com/brendan-rius/c-jwt-cracker).

## Usage

The application takes JWT as the only argument- it should automatically detect used alghoritm and start looking for the valid signature.

### Without building

```shell
go run . <jwt-you-want-to-crack>
```

### Building the library first

```shell
go build .
./<name of the created executable> <jwt-you-want-to-crack>
```

## Supported alghoritms

- HS256
- HS384
- HS512

## Contributors:
Thanks to [u/gnisten_](https://www.reddit.com/user/gnisten_/) for [fixing header recognition](https://www.reddit.com/r/golang/comments/gcrsmv/i_wrote_an_app_for_cracking_jwt_as_my_first/fpe45bo/).
