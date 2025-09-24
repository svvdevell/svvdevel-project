--
-- PostgreSQL database dump
--

\restrict cD1MtvCanALga3fVH6x1H92ckedwJvRIc84Y3EeGuTs6rI6Wqu3cOqbiDwVgDIr

-- Dumped from database version 16.10 (Debian 16.10-1.pgdg13+1)
-- Dumped by pg_dump version 16.10 (Debian 16.10-1.pgdg13+1)

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
-- Name: car_images; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.car_images (
    id integer NOT NULL,
    car_request_id integer,
    file_name character varying(255) NOT NULL,
    file_path character varying(500) NOT NULL,
    file_size integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.car_images OWNER TO developer;

--
-- Name: car_images_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.car_images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.car_images_id_seq OWNER TO developer;

--
-- Name: car_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.car_images_id_seq OWNED BY public.car_images.id;


--
-- Name: car_requests; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.car_requests (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    car_brand character varying(255) NOT NULL,
    phone character varying(20) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.car_requests OWNER TO developer;

--
-- Name: car_requests_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.car_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.car_requests_id_seq OWNER TO developer;

--
-- Name: car_requests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.car_requests_id_seq OWNED BY public.car_requests.id;


--
-- Name: car_images id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.car_images ALTER COLUMN id SET DEFAULT nextval('public.car_images_id_seq'::regclass);


--
-- Name: car_requests id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.car_requests ALTER COLUMN id SET DEFAULT nextval('public.car_requests_id_seq'::regclass);


--
-- Data for Name: car_images; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.car_images (id, car_request_id, file_name, file_path, file_size, created_at) FROM stdin;
\.


--
-- Data for Name: car_requests; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.car_requests (id, name, car_brand, phone, description, created_at) FROM stdin;
\.


--
-- Name: car_images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.car_images_id_seq', 1, false);


--
-- Name: car_requests_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.car_requests_id_seq', 1, false);


--
-- Name: car_images car_images_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.car_images
    ADD CONSTRAINT car_images_pkey PRIMARY KEY (id);


--
-- Name: car_requests car_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.car_requests
    ADD CONSTRAINT car_requests_pkey PRIMARY KEY (id);


--
-- Name: car_images car_images_car_request_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.car_images
    ADD CONSTRAINT car_images_car_request_id_fkey FOREIGN KEY (car_request_id) REFERENCES public.car_requests(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict cD1MtvCanALga3fVH6x1H92ckedwJvRIc84Y3EeGuTs6rI6Wqu3cOqbiDwVgDIr

