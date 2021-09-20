--
-- PostgreSQL database dump
--

-- Dumped from database version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)

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
-- Name: leads; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.leads (
    id integer NOT NULL,
    email character varying(150) NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    salutation character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.leads OWNER TO golang;

--
-- Name: leads_id_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.leads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.leads_id_seq OWNER TO golang;

--
-- Name: leads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: golang
--

ALTER SEQUENCE public.leads_id_seq OWNED BY public.leads.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO golang;

--
-- Name: leads id; Type: DEFAULT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.leads ALTER COLUMN id SET DEFAULT nextval('public.leads_id_seq'::regclass);


--
-- Name: leads leads_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.leads
    ADD CONSTRAINT leads_pkey PRIMARY KEY (id);


--
-- Name: leads_email_idx; Type: INDEX; Schema: public; Owner: golang
--

CREATE UNIQUE INDEX leads_email_idx ON public.leads USING btree (email);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: golang
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

