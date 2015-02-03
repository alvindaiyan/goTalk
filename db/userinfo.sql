CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL CONSTRAINT unq_username UNIQUE,
    password character varying(500) NOT NULL,
    Created date,
    CONSTRAINT user_pkey PRIMARY KEY (uid)
)