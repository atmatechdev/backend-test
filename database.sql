--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted timestamp with time zone,
    title text NOT NULL,
    description text NOT NULL,
    content text NOT NULL,
    created_by_id bigint
);


ALTER TABLE public.books OWNER TO postgres;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO postgres;

--
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    username text,
    password text,
    active boolean DEFAULT true
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: books id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (id, created_at, updated_at, deleted, title, description, content, created_by_id) FROM stdin;
1	2022-03-23 09:48:18.717821+07	2022-03-23 09:48:18.717821+07	\N	1984	About superpower	Book content of 1984	1
2	2022-03-23 09:48:41.560318+07	2022-03-23 09:48:41.560318+07	2022-03-23 09:49:58.039422+07	Oliver Twist	Poor boy	Book content of Oliver Twist	1
3	2022-03-23 09:51:05.718123+07	2022-03-23 09:56:14.437564+07	\N	Oliver Twist	Poor boy	UPDATED book content	1
9	2022-03-23 11:06:35.839649+07	2022-03-23 11:06:35.839649+07	\N	The Stranger	Description of The Stranger	Content of The Stranger	1
10	2022-03-23 11:07:47.09895+07	2022-03-23 11:07:47.09895+07	\N	Peter Pan	Description of Peter Pan	Content of Peter Pan	1
12	2022-03-23 11:08:28.118824+07	2022-03-23 11:08:28.118824+07	\N	Uncle Tom's Cabin	Description of Uncle Tom's Cabin	Content of Uncle Tom's Cabin	1
13	2022-03-23 11:08:56.718911+07	2022-03-23 11:08:56.718911+07	\N	Frankenstein	Description of Frankenstein	Content of Frankenstein	1
14	2022-03-23 11:09:13.004996+07	2022-03-23 11:09:13.004996+07	\N	Utopia	Description of Utopia	Content of Utopia	1
15	2022-03-23 11:09:24.98726+07	2022-03-23 11:09:24.98726+07	\N	Fahrenheit 451	Description of Fahrenheit 451	Content of Fahrenheit 451	1
16	2022-03-23 11:10:01.128086+07	2022-03-23 11:10:01.128086+07	\N	Lord of The Flies	Description of Lord of The Flies	Content of Lord of The Flies	1
17	2022-03-23 11:10:12.757044+07	2022-03-23 11:10:12.757044+07	\N	The Count of Monte Cristo	Description of The Count of Monte Cristo	Content of The Count of Monte Cristo	1
18	2022-03-23 11:14:05.359042+07	2022-03-23 11:14:05.359042+07	\N	Catch-22	Description of Catch-22	Content of Catch-22	1
19	2022-03-23 11:14:19.447731+07	2022-03-23 11:14:19.447731+07	\N	Life of Pi	Description of Life of Pi	Content of Life of Pi	1
20	2022-03-23 11:16:48.441959+07	2022-03-23 11:16:48.441959+07	\N	Ulysses	Description of Ulysses	Content of Ulysses	5
11	2022-03-23 11:08:13.333419+07	2022-03-23 11:08:13.333419+07	2022-03-23 11:17:23.187123+07	Peter Pan	Description of Peter Pan	Content of Peter Pan	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, created_at, username, password, active) FROM stdin;
1	2022-03-23 09:47:21.890698+07	test	$2a$10$1H2PcEYczUMube5Tsnw5NOaI7yC5VRfGp.3/yvztbENMXJ6Y8sfsu	t
4	2022-03-23 10:09:10.739265+07	test2	$2a$10$mycUYHq4W9XGg1U9lIbXeOV5bkDgPnl6x6OeFsrxEGvCsiMBnauIK	t
5	2022-03-23 11:15:46.780576+07	test3	$2a$10$cv4HivvJRYBps90w/CvSCet27g6.7N3dowMfiTrqH9GWOjfzsp78C	t
\.


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.books_id_seq', 20, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- Name: users idx_users_username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT idx_users_username UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_books_deleted; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_books_deleted ON public.books USING btree (deleted);


--
-- Name: books fk_books_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT fk_books_user FOREIGN KEY (created_by_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

