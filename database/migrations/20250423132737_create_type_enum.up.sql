DROP TYPE IF EXISTS type;
DROP TYPE IF EXISTS recurring_frequency;

CREATE TYPE type AS ENUM ('income','expense','tranfer');

CREATE TYPE recurring_frequency AS ENUM ('daily','weekly','monthly','yearly');