# This project is implemented based on CleanArchitecture , designed by BOB

    ! Document

[https://medium.com/@jamal.kaksouri/building-better-go-applications-with-clean-architecture-a-practical-guide-for-beginners-98ea061bf81a]

    ! Video

[https://www.youtube.com/watch?v=ffYCgcDgsfw]

# How to create migrate file

migrate create -ext sql -dir Migration/sql -seq create_users_table

# How to run migrate up

make <command_in_makefile>

# How to connect to redis from docker, and how to check if bit exist

docker exec -it authentication_redis redis-cli

-   CHECK BIT: getbit u:bit 4097051691
-   SET BIT : setbit u:bit 4097051691
