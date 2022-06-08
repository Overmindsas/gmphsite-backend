--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.3 (Debian 13.3-1.pgdg100+1)

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
-- Name: userdata; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userdata (
    qotd_date character varying(255),
    id integer,
    dialoge boolean,
    private_field boolean,
    tags character varying(255),
    url character varying(255),
    favorite_count integer,
    upvotes_count integer,
    downvotes_count integer,
    author character varying(255),
    author_permalink character varying(255),
    body character varying(255)
);


ALTER TABLE public.userdata OWNER TO postgres;

--
-- Data for Name: userdata; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userdata (qotd_date, id, dialoge, private_field, tags, url, favorite_count, upvotes_count, downvotes_count, author, author_permalink, body) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

