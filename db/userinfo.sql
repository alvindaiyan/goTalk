CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL CONSTRAINT unq_username UNIQUE,
    password character varying(500) NOT NULL,
    Created date,
    CONSTRAINT user_pkey PRIMARY KEY (uid)
)

CREATE TABLE useruserlink
(
    id serial NOT NULL,
    user1id serial NOT NULL,
    user2id serial NOT NULL,
    created date,
    CONSTRAINT useruserlinkkey PRIMARY KEY (id, user1id, user2id),  
    FOREIGN KEY (user1id) REFERENCES userinfo(uid),
    FOREIGN KEY (user2id) REFERENCES userinfo(uid)
)

ALTER TABLE useruserlink add CONSTRAINT ck_userids check(user1id <> user2id)