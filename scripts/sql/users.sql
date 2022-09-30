SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres;
--
CREATE TABLE public.users (
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(255) NOT NULL,
    email CHARACTER VARYING(255) NOT NULL,
    password_hash CHARACTER VARYING(255) NOT NULL,
);

ALTER TABLE public.users OWNER TO postgres;

---
--- password is 'supersecret'
---
INSERT INTO "public"."users"("name", "email", "password_hash") 
VALUES(E'admin', E'admin@fakemail.com', E'$2a$14$kgcE0T7wwmgtmfqLYbCHGeFg1dLVtqX5GwhjP7E/.wtruZSTsWuR2');