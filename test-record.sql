-- 1) A www.example.com -> 1.2.3.4
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x7777772e6578616d706c652e636f6d,
        1,
        0x01020304,
        300
    );

-- 2) PTR 1.2.3.4 -> host.example.com
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x342e332e322e312e696e2d616464722e61727061,
        12,
        0x686f73742e6578616d706c652e636f6d,
        300
    );

-- 3) AAAA www.example.com -> 2001:db8::1
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x7777772e6578616d706c652e636f6d,
        28,
        0x20010db8000000000000000000000001,
        300
    );

-- 4) CNAME alias.example.com -> www.example.com
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x616c6961732e6578616d706c652e636f6d,
        5,
        0x7777772e6578616d706c652e636f6d,
        300
    );

-- 5) MX example.com -> 10 mail.example.com
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x6578616d706c652e636f6d,
        15,
        0x3130206d61696c2e6578616d706c652e636f6d,
        300
    );

-- 6) TXT example.com -> "hello world"
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x6578616d706c652e636f6d,
        16,
        0x68656c6c6f20776f726c64,
        300
    );

-- 7) NS example.com -> ns1.example.com
INSERT INTO
    records (name, type, value, ttl)
VALUES
    (
        0x6578616d706c652e636f6d,
        2,
        0x6e73312e6578616d706c652e636f6d,
        300
    );