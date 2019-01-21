#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL

CREATE DATABASE shopify;

EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname shopify <<-EOSQL

--
-- Name: cart_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE cart_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE cart_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE carts (
    id bigint DEFAULT nextval('cart_id_seq'::regclass) NOT NULL,
    item_id bigint NOT NULL
);


ALTER TABLE carts OWNER TO postgres;

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE item_id_seq
    START WITH 100
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE item_id_seq OWNER TO postgres;

--
-- Name: inventory; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE inventory (
    product_id bigint NOT NULL,
    item_id bigint DEFAULT nextval('item_id_seq'::regclass) NOT NULL,
    available boolean
);


ALTER TABLE inventory OWNER TO postgres;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE product_id_seq
    START WITH 100
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE product_id_seq OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE products (
    product_id bigint DEFAULT nextval('product_id_seq'::regclass) NOT NULL,
    price double precision,
    title character varying
);


ALTER TABLE products OWNER TO postgres;

--
-- Name: products_view; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW products_view AS
SELECT
    NULL::bigint AS product_id,
    NULL::character varying AS title,
    NULL::double precision AS price,
    NULL::bigint AS inventory_count;


ALTER TABLE products_view OWNER TO postgres;

--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY carts (id, item_id) FROM stdin;
\.


--
-- Data for Name: inventory; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY inventory (product_id, item_id, available) FROM stdin;
1	2	t
2	3	t
1	1	t
11	8	t
11	7	t
11	6	t
11	5	t
11	4	t
3	17	t
3	16	t
3	15	t
3	14	t
3	13	t
3	12	t
3	11	t
3	10	t
3	9	t
10	23	t
10	22	t
10	21	t
9	27	t
9	26	t
9	25	t
9	24	t
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY products (product_id, price, title) FROM stdin;
1	19.9899999999999984	Regular T-Shirt
2	29.9899999999999984	Regular Long Sleeve Shirt
3	39.990000000000002	Regular Hoodie
9	59.990000000000002	Premium Hoodie
10	29.9899999999999984	Compression Pants
11	199.990000000000009	Winter Jacket
\.


--
-- Name: cart_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('cart_id_seq', 6, true);


--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('item_id_seq', 27, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('product_id_seq', 11, true);


--
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (id, item_id);


--
-- Name: inventory inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY inventory
    ADD CONSTRAINT inventory_pkey PRIMARY KEY (product_id, item_id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);


--
-- Name: products_view _RETURN; Type: RULE; Schema: public; Owner: postgres
--

CREATE OR REPLACE VIEW products_view AS
 SELECT p.product_id,
    p.title,
    p.price,
    sum(
        CASE
            WHEN i.available THEN 1
            ELSE 0
        END) AS inventory_count
   FROM (products p
     LEFT JOIN inventory i ON ((p.product_id = i.product_id)))
  GROUP BY p.product_id
  ORDER BY p.product_id;


--
-- PostgreSQL database dump complete
--

EOSQL
